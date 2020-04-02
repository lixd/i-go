package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// bufio 包实现了带缓存的 I/O 操作

/**
 * 首先看reader和writer基本的结构
 * // Reader implements buffering for an io.Reader object.
 * type Reader struct {
 *	buf          []byte
 *	rd           io.Reader // reader provided by the client
 *	r, w         int       // buf read and write positions
 *	err          error
 *	lastByte     int
 *	lastRuneSize int
 * }
 *
 *
 * // Writer implements buffering for an io.Writer object.
 * // If an error occurs writing to a Writer, no more data will be
 * // accepted and all subsequent writes will return the error.
 * // After all data has been written, the client should call the
 * // Flush method to guarantee all data has been forwarded to
 * // the underlying io.Writer.
 * type Writer struct {
 *	err error
 *	buf []byte
 *	n   int
 *	wr  io.Writer
 * }
 *
 *
 *
 * // ReadWriter 集成了 bufio.Reader 和 bufio.Writer, 实现了 io.ReadWriter 接口
 * type ReadWriter struct {
 *	*Reader
 *	*Writer
 * }
 */

func main() {
	// 1: 使用bufio.NewReader构造一个reader
	inputReadBuf := strings.NewReader("1234567890")
	reader := bufio.NewReader(inputReadBuf)

	// 2: 使用bufio.NewWriter构造一个writer
	buf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf)

	// 3: 函数Peek函数: 返回缓存的一个Slice(引用,不是拷贝)，引用缓存中前n字节数据
	// > 如果引用的数据长度小于 n，则返回一个错误信息
	// > 如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过Peek的返回值，可以修改缓存中的数据, 但不能修改底层io.Reader中的数据
	fmt.Println("-------------------3----------------")
	b, err := reader.Peek(5)
	if err != nil {
		fmt.Printf("Read data error")
		return
	}
	// 修改第一个字符
	b[0] = 'A'
	// 重新读取
	b, _ = reader.Peek(5)
	writer.Write(b)
	writer.Flush()
	fmt.Println("buf(Changed): ", buf, "\ninputReadBuf(Not Changed): ", *inputReadBuf)

	// 4: Read函数, 每次读取一定量的数据, 这个由buf大小觉得, 所以我们可以循环读取数据, 直到Read返回0说明读取数据结束
	fmt.Println("-------------------4----------------")
	for {
		b1 := make([]byte, 3)
		n1, _ := reader.Read(b1)
		if n1 <= 0 {
			break
		}
		fmt.Println(n1, string(b1))
	}

	// 5: ReadByte和UnreadByte函数
	// ReadByte 从 b 中读出一个字节并返回, 如果 b 中无可读数据，则返回一个错误
	// UnreadByte 撤消最后一次读出的字节, 只有最后读出的字节可以被撤消, 无论任何操作，只要有内容被读出，就可以用 UnreadByte 撤消一个字节
	fmt.Println("-------------------5----------------")
	inputReadBuf2 := strings.NewReader("1234567890")
	reader2 := bufio.NewReader(inputReadBuf2)
	// 读一个字节
	b2, _ := reader2.ReadByte()
	fmt.Println(string(b2))
	// Unread一个字节
	reader2.UnreadByte()
	b2, _ = reader2.ReadByte()
	fmt.Println(string(b2))

	// 6: ReadRune和UnreadRune函数, 类似上面两个函数
	// ReadRune读出一个 UTF8 编码的字符并返回编码长度, 如果UTF8序列无法解码出一个正确的Unicode字符, 则只读出b中的一个字节，size 返回 1
	fmt.Println("-------------------6----------------")
	inputReadBuf3 := strings.NewReader("中文1234567890")
	reader3 := bufio.NewReader(inputReadBuf3)
	b3, size, _ := reader3.ReadRune()
	fmt.Println(string(b3), size)
	reader3.UnreadRune()
	b3, size, _ = reader3.ReadRune()
	fmt.Println(string(b3), size)
	// 执行UnreadRune时候, 如果之前一步不是ReadRune, 那么会报错, 看下面
	b33, _ := reader3.ReadByte()
	fmt.Println(string(b33))
	err3 := reader3.UnreadRune()
	if err3 != nil {
		fmt.Println("ERR")
	}

	// 7: 读取缓冲区中数据字节数(只有执行读才会使用到缓冲区, 否则是没有的)
	fmt.Println("-------------------7----------------")
	inputReadBuf4 := strings.NewReader("中文1234567890")
	reader4 := bufio.NewReader(inputReadBuf4)
	// 下面返回0, 因为还没有开始读取, 缓冲区没有数据
	fmt.Println(reader4.Buffered())
	// 下面返回strings的整体长度16(一个人中文是3长度)
	reader4.Peek(1)
	fmt.Println(reader4.Buffered())
	// 下面返回15, 因为readByte已经读取一个字节数据, 所以缓冲区还有15字节
	reader4.ReadByte()
	fmt.Println(reader4.Buffered())
	// 下面的特别有意思: 上面已经读取了一个字节, 想当于是将"中"读取了1/3, 那么如果现在使用readRune读取, 那么
	// 由于无法解析, 那么仅仅读取一个byte, 所以下面的结果很显然
	// 第一次: 无法解析, 那么返回一个byte, 所以输出的是14
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	// 第二次读取, 还剩下"中"最后一个字节, 所以也会err, 所以输出13
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	// 现在"中"读完了, 那么开始完整读取"文", 这个OK的, 可以解析的, 所以可以读取三字节, 那么剩下10字节
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())

	// 8: ReadSlice查找 delim 并返回 delim 及其之前的所有数据的切片, 该操作会读出数据，返回的切片是已读出数据的"引用"
	// 如果 ReadSlice 在找到 delim 之前遇到错误, 则读出缓存中的所有数据并返回，同时返回遇到error（通常是 io.EOF）
	// 如果 在整个缓存中都找不到 delim，则返回 ErrBufferFull
	// 如果 ReadSlice 能找到 delim，则返回 nil
	// 注意: 因为返回的Slice数据有可能被下一次读写操作修改, 因此大多数操作应该使用 ReadBytes 或 ReadString，它们返回数据copy
	// 不推荐!
	fmt.Println("-------------------8----------------")
	inputReadBuf5 := strings.NewReader("中文123 4567 890")
	reader5 := bufio.NewReader(inputReadBuf5)
	for {
		b5, err := reader5.ReadSlice(' ')
		fmt.Println(string(b5))
		// 读到最后
		if err == io.EOF {
			break
		}
	}

	// 9: ReadLine 是一个低级的原始的行读取操作, 一般应该使用 ReadBytes('\n') 或 ReadString('\n')
	// ReadLine 通过调用 ReadSlice 方法实现，返回的也是"引用", 回一行数据，不包括行尾标记（\n 或 \r\n）
	// 如果 在缓存中找不到行尾标记，设置 isPrefix 为 true，表示查找未完成
	// 如果 在当前缓存中找到行尾标记，将 isPrefix 设置为 false，表示查找完成
	// 如果 ReadLine 无法获取任何数据，则返回一个错误信息（通常是 io.EOF）
	// 不推荐!
	fmt.Println("-------------------9----------------")
	inputReadBuf6 := strings.NewReader("中文123\n4567\n890")
	reader6 := bufio.NewReader(inputReadBuf6)
	for {
		l, p, err := reader6.ReadLine()
		fmt.Println(string(l), p, err)
		if err == io.EOF {
			break
		}
	}

	// 10: ReadBytes查找 delim 并读出 delim 及其之前的所有数据
	// 如果 ReadBytes 在找到 delim 之前遇到错误, 则返回遇到错误之前的所有数据，同时返回遇到的错误（通常是 io.EOF）
	// 如果 ReadBytes 找不到 delim 时，err != nil
	// 返回的是数据的copy, 不是引用
	fmt.Println("-------------------10----------------")
	inputReadBuf7 := strings.NewReader("中文123;4567;890")
	reader7 := bufio.NewReader(inputReadBuf7)
	for {
		line, err := reader7.ReadBytes(';')
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}

	// 11: ReadString返回的是字符串, 不是bytes
	fmt.Println("-------------------11----------------")
	inputReadBuf8 := strings.NewReader("中文123;4567;890")
	reader8 := bufio.NewReader(inputReadBuf8)
	for {
		line, err := reader8.ReadString(';')
		fmt.Println(line)
		if err != nil {
			break
		}
	}

	//12: Flush函数用于提交数据, 立即更新
	// Available函数返回缓存中的可以空间
	// b10是保存数据的数组, 不是writer的缓冲区, 别搞错了
	fmt.Println("-------------------12----------------")
	b10 := bytes.NewBuffer(make([]byte, 30))
	// 下面会分配4096字节空间缓冲区
	writer10 := bufio.NewWriter(b10)
	writer10.WriteString("1234567890")
	// 此时没有flush, 那么输出的是"", 但是缓冲区使用了10个字节, 那么剩下4086, Buffered()返回的是缓冲区还没有提交的数据, 此处显然是10
	fmt.Println(writer10.Available(), writer10.Buffered(), b10)
	// 下面flush后, 将缓冲区的数据全部写入b10中, 缓冲区被清空, 所以缓冲区变成4096, Buffered()返回的是0, 说明数据被写入
	writer10.Flush()
	fmt.Println(writer10.Available(), writer10.Buffered(), b10)

	// 13: WriteString(...), Write(...), WriteByte(...), WriteRune(...)函数
	// 都是写数据函数
	fmt.Println("-------------------13----------------")
	b11 := bytes.NewBuffer(make([]byte, 1024))
	writer11 := bufio.NewWriter(b11)
	writer11.WriteString("ABC")
	writer11.WriteByte(byte('M'))
	// Rune的意思是: 代表一个字符, 那么需要一次一个字符写入
	writer11.WriteRune(rune('好'))
	writer11.WriteRune(rune('么'))
	writer11.Write([]byte("1234567890"))
	writer11.Flush()
	fmt.Println(b11)

	// 14: WriteTo函数
	fmt.Println("-------------------14----------------")
	inputReadBuf9 := strings.NewReader("中文1234567890")
	reader9 := bufio.NewReader(inputReadBuf9)
	b9 := bytes.NewBuffer(make([]byte, 0))
	reader9.WriteTo(b9)
	fmt.Println(b9)

	// 15: ReadFrom函数
	fmt.Println("-------------------15----------------")
	inputReadBuf15 := strings.NewReader("率哪来的顺丰内部了第三方吧")
	b15 := bytes.NewBuffer(make([]byte, 0))
	writer15 := bufio.NewWriter(b15)
	writer15.ReadFrom(inputReadBuf15)
	fmt.Println(b15)

	// ReadAt 根据字符串做偏移 即每次都是从第一个字节开始算
	fmt.Println("-------------------16----------------")
	r := strings.NewReader("abcdefghijklmn")

	var bufAt, buf1 []byte
	buf1 = make([]byte, 5)
	r.Read(buf1)
	fmt.Println("剩余未读的长度", r.Len())     //剩余未读的长度 9
	fmt.Println("已读取的内容", string(buf1)) //已读取的内容 abcde
	bufAt = make([]byte, 256)
	r.ReadAt(bufAt, 5)
	fmt.Println(string(bufAt)) //fghijklmn

	//测试下是否影响Len和Read方法
	fmt.Println("剩余未读的长度", r.Len())     //剩余未读的长度 9
	fmt.Println("已读取的内容", string(buf1)) //已读取的内容 abcde

	// 	Seek 根据指定做偏移 即每次都是从指定位置开始计算偏移量

	var buf2 []byte
	buf2 = make([]byte, 5)
	r.Read(buf2)
	fmt.Println(string(buf2)) //adcde

	buf2 = make([]byte, 5)
	r.Seek(-2, io.SeekCurrent) //从当前位置向前偏移两位 （5-2)
	r.Read(buf2)
	fmt.Println(string(buf2)) //defgh

	buf2 = make([]byte, 5)
	r.Seek(-3, io.SeekEnd) //设置当前位置是末尾前移三位
	r.Read(buf2)
	fmt.Println(string(buf2)) //lmn

	buf2 = make([]byte, 5)
	r.Seek(3, io.SeekStart) //设置当前位置是起始位后移三位
	r.Read(buf2)
	fmt.Println(string(buf2)) //defgh
}
