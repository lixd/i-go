package main

import (
	"bufio"
	"fmt"
	"io"
)

// 去掉无意义的判断

/*// AuthenticateRequest
func AuthenticateRequest(r *Request)error{
	err:=authenticate(r.User)
	if err != nil{
		return err
	}
	return nil
}
// AuthenticateRequest2 完全可以直接向下面这样
func AuthenticateRequest2(r *Request)error{
	return authenticate(r.User)
}
*/

// 一个写 HTTP 响应的例子

type Header struct {
	Key, Value string
}
type Status struct {
	Code   int
	Reason string
}

// WriteResponse 第一版本实现 好几个地方都需要判断 error 比较麻烦
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}
	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}
	_, err = io.Copy(w, body)
	return err
}

// 优化后
type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	// 第二步 在下一次调用的时候判定 err 字段是否不为空
	// 不为空说明上一次调用肯定报错了 这里直接返回
	if e.err != nil {
		return 0, e.err
	}
	var n int
	// 第一步 用 err字段把Write返回的错误接收一下
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

// WriteResponse2 中途不需要任何判断 直接在最后返回错误即可
func WriteResponse2(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	// 这里不判断 error
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	for _, h := range headers {
		// 这里也不判断
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}
	// 这里还是不判断
	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)
	// 最终返回的时候直接把 err 字段返回即可
	return ew.err
}

//  统计 io.Reader 读取的行数

// CountLines 第一版本 需要多个地方判断err 看着比较乱
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// CountLines2  优化后 中途不需要判断 error 最后直接返回即可
func CountLines2(r io.Reader) (int, error) {
	// 1.Scanner 将 error 存放在自己内部 不需要在外部判断
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	// 2. 返回的时候直接将 Scanner 中的 err 返回即可
	return lines, sc.Err()
}
