package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	// file -->file文件对象 file指针 file文件句柄几种叫法
	file, e := os.Open("d:/test.txt")
	if e != nil {
		fmt.Println("open file error=", e)
	}
	// &{0xc000084780} 可以看出 文件就是 一个指针
	// fmt.Println(file)

	// 关闭文件
	defer file.Close()

	/*
		const (
			defaultBufSize = 4096
		)
	*/
	// 创建一个reader 带缓冲的 默认缓冲大小为 4096
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		// 输出内容
		fmt.Print(str)
		if err == io.EOF { // io.EOF表示读到文件的末尾
			break
		}

	}
	fmt.Println()
	fmt.Println("文件读取结束")

	fmt.Println("----------------")
	// 使用 ioutil.ReadFile 一次性读取文件
	bytes, e := ioutil.ReadFile("d:/test.txt")
	if e != nil {
		fmt.Println("open file error=", e)
	}
	fmt.Println(string(bytes))

	fmt.Println("----------------")

	openFile, e := os.OpenFile("d:/test2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("OpenFile error=", e)
	}
	defer openFile.Close()
	writer := bufio.NewWriter(openFile)
	for i := 0; i < 5; i++ {
		// 调用WriteString 时先写入到缓存
		// writer.Flush()后才真正将缓冲的数据写到磁盘
		writer.WriteString("NewString \n")
	}
	e = writer.Flush()
	if e != nil {
		fmt.Println("Flush error=", e)
	}

	fmt.Println("-------------")
	src := "d:/abc.png"
	dst := "d:/def.png"
	written, e := CopyFile(src, dst)
	if e != nil {
		fmt.Printf("Copy error: %v \n", e)
	} else {
		fmt.Println("Copy Success")
	}
	fmt.Println(written)
	fmt.Println("-------------")
	fmt.Println("命令行的所有参数", len(os.Args))
	for i, value := range os.Args {
		fmt.Printf("index:%d value:%v \n", i, value)
	}

}

func CopyFile(src string, dst string) (written int64, err error) {
	srcFile, e := os.Open(src)
	if e != nil {
		fmt.Printf("Open error: %v \n", e)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, e := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Printf("Open error: %v \n", e)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

// lineByLine 逐行读取
func lineByLine(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Printf(line)
	}
	return nil
}

// write 写入文件的几种方法
func write() {
	writeOne()

}
func writeOne() {
	s := []byte("Data to write\n")

	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))
}

func writeTwo() {
	s := []byte("Data to write\n")
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
}
func writeThree() {
	s := []byte("Data to write\n")
	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 和 writeTwo 一样的 只是这个你带缓冲的
	w := bufio.NewWriter(f3)
	n, err := w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}
func writeFour() {
	s := []byte("Data to write\n")
	f4 := "f4.txt"
	// 这个只是把创建文件和写入数据封装了一下
	err := ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func writeFive() {
	s := []byte("Data to write\n")
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
