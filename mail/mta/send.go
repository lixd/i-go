package mta

import (
	"bufio"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/smtp"
	"time"
)

/*
1.根据收件人获取到domain
2.解析该domain的mx记录
3.往该mx记录的域名发送邮件即可
	主要为如下几个步骤: Dial Hello Mail Rcpt Data Quit
4.其他可参考 SMTP rfc 文档 https://www.ietf.org/rfc/rfc2554.txt
*/

type Sender struct {
	Hostname string // 当前邮件服务器的域名 可以填 localhost
	Insecure bool   // 是否跳过TLS认证
}

// Send 越过 MUA 直接将邮件发送给远程 MTA
/*
from 发件人 xxx@qq.com
to 收件人 只支持单个收件人 u1@qq.cm
*/
func (s *Sender) Send(from, to string, email []byte) error {
	return s.doSend(from, to, email)
}

func (s *Sender) doSend(from, to string, data []byte) error {
	mxs, err := lookMX(to)
	if err != nil {
		return err
	}

	for _, mx := range mxs {
		err = s.sendBySMTP(from, to, data, mx)
		if err != nil {
			continue
		}
		// 往其中一个MX记录发送成功后就退出
		return nil
	}
	return err
}

// sendBySMTP 使用 SMTP 协议将邮件发送出去
/*
from 发件人 xxx@qq.com
to 收件人 xxx@qq.com
insecure 是否启用TLS false
data 邮件内容
mx 收件服务器的 MX 解析记录
*/
func (s *Sender) sendBySMTP(from, to string, data []byte, mx *net.MX) error {
	// c, err := smtp.Dial(mx.Host + ":25")
	c, err := clientWithTimeOut(mx.Host)
	if err != nil {
		return err
	}
	c, err = wrapperClient(c)
	if err != nil {
		return err
	}
	if err = c.Hello(s.Hostname); err != nil {
		return err
	}

	if ok, _ := c.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{ServerName: mx.Host, InsecureSkipVerify: s.Insecure}
		if err = c.StartTLS(tlsConfig); err != nil {
			return err
		}
	}
	// 域名国际化检测
	from, err = prepareForSMTPUTF8(c, from)
	if err != nil {
		return err
	}
	to, err = prepareForSMTPUTF8(c, to)
	if err != nil {
		return err
	}

	err = c.Mail(from)
	if err != nil {
		return err
	}
	err = c.Rcpt(to)
	if err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = wc.Write(data)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}
	err = c.Quit()
	if err != nil {
		return err
	}
	return nil
}

// lookMX 根据发送人地址(xxx@gmail.com) 解析出对应的 MX 记录
func lookMX(addr string) ([]*net.MX, error) {
	_, domain, err := SplitAddress(addr)
	if err != nil {
		return nil, err
	}

	mxs, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	if len(mxs) == 0 { // 如果对应域名没有解析MX记录则直接发送给主域名
		mxs = []*net.MX{{Host: domain}}
	}
	return mxs, nil
}

// wrapperClient 对 client 的 reader 进行包装
func wrapperClient(c *smtp.Client) (*smtp.Client, error) {
	// Wrap the textproto.Conn reader so we are not exposed to a memory
	// exhaustion DoS on very long replies from the server.
	// Limit to 2 MiB total (all replies through the lifetime of the client),
	// which should be plenty for our uses of SMTP.
	lr := &io.LimitedReader{R: c.Text.Reader.R, N: 2 * 1024 * 1024}
	c.Text.Reader.R = bufio.NewReader(lr)
	return c, nil
}

func clientWithTimeOut(mx string) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", mx+":"+"25", time.Minute)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(time.Minute))
	return smtp.NewClient(conn, mx)
}

// prepareForSMTPUTF8 域名国际化转换 IDNA
/*
如果是标准国际化域名则不需要转换
如果远程 MTA 支持 SMTPUTF8 也不需要转换
否则需要转换 不然远程 MTA 无法识别该域名
*/
func prepareForSMTPUTF8(c *smtp.Client, addr string) (string, error) {
	if IsAllASCII(addr) {
		return addr, nil
	}

	if ok, _ := c.Extension("SMTPUTF8"); ok {
		return addr, nil
	}

	user, domain, _ := SplitAddress(addr)

	if !IsAllASCII(user) { // 用户名部分无法转换直接报错返回
		return addr, errors.New("local part is not ASCII but server does not support SMTPUTF8")
	}

	domain, err := ToASCII(domain)
	if err != nil {
		// non-ASCII domain is not IDNA safe
		return addr, err
	}
	return user + "@" + domain, nil
}
