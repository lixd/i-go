package ismtp

import (
	"crypto/tls"
	"net"
	"time"

	"golang.org/x/net/idna"
)

// SMTP delivers remote mail via outgoing SMTP.
type SMTP struct {
	HelloDomain string
	Dinfo       interface{}
	STSCache    interface{}
}

func lookupMXs(domain string) ([]string, error, bool) {
	domain, err := idna.ToASCII(domain)
	if err != nil {
		return nil, err, true
	}

	mxs := []string{}

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		// There was an error. It could be that the domain has no MX, in which
		// case we have to fall back to A, or a bigger problem.
		// Unfortunately, go's API doesn't let us easily distinguish between
		// them. For now, if the error is permanent, we assume it's because
		// there was no MX and fall back, otherwise we return.
		// TODO: Use dnsErr.IsNotFound once we can use Go >= 1.13.
		dnsErr, ok := err.(*net.DNSError)
		if !ok {
			return nil, err, true
		} else if dnsErr.Temporary() {
			return nil, err, false
		}

		// Permanent error, we assume MX does not exist and fall back to A.
		mxs = []string{domain}
	} else {
		// Convert the DNS records to a plain string slice. They're already
		// sorted by priority.
		for _, r := range mxRecords {
			mxs = append(mxs, r.Host)
		}
	}

	// Note that mxs could be empty; in that case we do NOT fall back to A.
	// This case is explicitly covered by the SMTP RFC.
	// https://tools.ietf.org/html/rfc5321#section-5.1

	// Cap the list of MXs to 5 hosts, to keep delivery attempt times
	// sane and prevent abuse.
	if len(mxs) > 5 {
		mxs = mxs[:5]
	}

	return mxs, nil, true
}

// Deliver an email. On failures, returns an error, and whether or not it is
// permanent.
func (s *SMTP) Deliver(from string, to string, data []byte) (error, bool) {
	_, toDomain := Split(to)
	mxs, err, perm := lookupMXs(toDomain)
	if err != nil || len(mxs) == 0 {
		// Note this is considered a permanent error.
		// This is in line with what other servers (Exim) do. However, the
		// downside is that temporary DNS issues can affect delivery, so we
		// have to make sure we try hard enough on the lookup above.
		return err, perm
	}

	// a.stsPolicy = s.fetchSTSPolicy(toDomain)

	for _, mx := range mxs {
		// if a.stsPolicy != nil && !a.stsPolicy.MXIsAllowed(mx) {
		// 	a.tr.Printf("%q skipped as per MTA-STA policy", mx)
		// 	continue
		// }

		var permanent bool
		err, permanent = s.deliver(mx, from, to, data)
		if err == nil {
			return nil, false
		}
		if permanent {
			return err, true
		}
	}

	// We exhausted all MXs failed to deliver, try again later.
	return err, false
}

func (a *SMTP) deliver(mx, from, to string, data []byte) (error, bool) {
	// Do we use insecure TLS?
	// Set as fallback when retrying.
	insecure := false // 是否开启 TLS
	// secLevel := domaininfo.SecLevel_PLAIN

retry:
	conn, err := net.DialTimeout("tcp", mx+":"+"25", time.Minute)
	if err != nil {
		return err, false
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Minute))

	c, err := NewClient(conn, mx)
	if err != nil {
		return err, false
	}

	if err = c.Hello(a.HelloDomain); err != nil {
		return err, false
	}

	// 检测是否支持 TLS 扩展
	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{
			ServerName:         mx,
			InsecureSkipVerify: insecure,
		}
		err = c.StartTLS(config)
		if err != nil {
			// Unfortunately, many servers use self-signed certs, so if we
			// fail verification we just try again without validating.
			if insecure {
				// tlsCount.Add("tls:failed", 1)
				return err, false
			}

			insecure = true
			goto retry
		}

		if config.InsecureSkipVerify {
			// tlsCount.Add("tls:insecure", 1)
			// secLevel = domaininfo.SecLevel_TLS_INSECURE
		} else {
			// tlsCount.Add("tls:secure", 1)
			// secLevel = domaininfo.SecLevel_TLS_SECURE
		}
	} else {
		// tlsCount.Add("plain", 1)
	}

	// if !a.Dinfo.OutgoingSecLevel(a.toDomain, secLevel) {
	// 	// We consider the failure transient, so transient misconfigurations
	// 	// do not affect deliveries.
	// 	slcResults.Add("fail", 1)
	// 	return a.tr.Errorf("Security level check failed (level:%s)", secLevel), false
	// }
	// slcResults.Add("pass", 1)

	// if a.stsPolicy != nil && a.stsPolicy.Mode == sts.Enforce {
	// 	// The connection MUST be validated by TLS.
	// 	// https://tools.ietf.org/html/rfc8461#section-4.2
	// 	if secLevel != domaininfo.SecLevel_TLS_SECURE {
	// 		stsSecurityResults.Add("fail", 1)
	// 		return a.tr.Errorf("invalid security level (%v) for STS policy",
	// 			secLevel), false
	// 	}
	// 	stsSecurityResults.Add("pass", 1)
	// 	a.tr.Debugf("STS policy: connection is using valid TLS")
	// }

	if err = c.MailAndRcpt(from, to); err != nil {
		return err, IsPermanent(err)
	}

	w, err := c.Data()
	if err != nil {
		return err, IsPermanent(err)
	}
	_, err = w.Write(data)
	if err != nil {
		return err, IsPermanent(err)
	}

	err = w.Close()
	if err != nil {
		return err, IsPermanent(err)
	}

	_ = c.Quit()

	return nil, false
}
