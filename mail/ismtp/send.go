package ismtp

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/smtp"
	"strings"
)

type Sender struct {
	Hostname string
}

// Send 直接将邮件发送给远程 MTA
/*
1.根据收件人获取到domain
2.解析该domain的mx记录
3.往该mx记录的域名发送邮件即可
	主要为如下几个步骤: Dial Hello Mail Rcpt Data Quit
4.其他可参考 SMTP rfc 文档 https://www.ietf.org/rfc/rfc2554.txt
*/
func (s *Sender) Send(from string, to []string, r io.Reader) error {
	// TODO: buffer r if sending to multiple recipients
	// TODO: group recipients with same domain

	for _, addr := range to {
		_, domain, err := splitAddress(addr)
		if err != nil {
			return err
		}

		mxs, err := net.LookupMX(domain)
		if err != nil {
			return err
		}
		if len(mxs) == 0 {
			mxs = []*net.MX{{Host: domain}}
		}

		for _, mx := range mxs {
			c, err := smtp.Dial(mx.Host + ":25")
			if err != nil {
				return err
			}

			if err := c.Hello(s.Hostname); err != nil {
				return err
			}

			if ok, _ := c.Extension("STARTTLS"); ok {
				tlsConfig := &tls.Config{ServerName: mx.Host}
				if err := c.StartTLS(tlsConfig); err != nil {
					return err
				}
			}

			if err := c.Mail(from); err != nil {
				return err
			}
			if err := c.Rcpt(addr); err != nil {
				return err
			}

			wc, err := c.Data()
			if err != nil {
				return err
			}
			if _, err := io.Copy(wc, r); err != nil {
				return err
			}
			if err := wc.Close(); err != nil {
				return err
			}

			if err := c.Quit(); err != nil {
				return err
			}
		}
	}

	return nil
}

// splitAddress 将user@domain 拆分为 user 和 domain
func splitAddress(addr string) (local, domain string, err error) {
	parts := strings.SplitN(addr, "@", 2)
	if len(parts) != 2 {
		return "", "", errors.New("mta: invalid mail address")
	}
	return parts[0], parts[1], nil
}
