package main

import (
	"crypto/rand"
	"fmt"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"io"
	"io/ioutil"
	mathrand "math/rand"
	"os"
	"time"
)

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o toolbar
func main() {
	//singleBar()
	//bytesDemo()
	bytesDemo2()
}

func bytesDemo() {
	// 传输文件时显示进度条以及速度等信息，提前知道文件大小的情况
	var total int64 = 1024 * 1024 * 1024 * 10
	reader := io.LimitReader(rand.Reader, total)

	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	bar := p.New(0,
		mpb.BarStyle().Rbound("|"),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.Name(" ] "),
			//decor.EwmaETA(decor.ET_STYLE_GO, 90),
			//decor.EwmaSpeed(decor.UnitKiB, "% .2f", decor.DSyncSpace),
			decor.AverageSpeed(decor.UnitKiB, "% .2f", decor.WCSyncWidth),
		),
	)
	bar.SetTotal(total, false)
	// create proxy reader
	proxyReader := bar.ProxyReader(reader)
	defer proxyReader.Close()

	// copy from proxyReader, ignoring errors
	_, _ = io.Copy(ioutil.Discard, proxyReader)

	p.Wait()
}

func bytesDemo2() {
	// 先创建一个文件用于删除
	file, err := os.Create("./tmp.txt")
	if err != nil {
		return
	}
	data := make([]byte, 1024*1024*1024*2)
	file.Write(data)
	file.Close()
	defer func() {
		os.Remove("./tmp.txt")
		os.Remove("./tmp2.txt")
	}()

	// 传输文件时显示进度条以及速度等信息，不提前知道文件大小的情况
	// 测试真正的传输文件时 EwmaSpeed 速度会计算错误，需要使用 AverageSpeed
	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	bar := p.New(1024*1024*1024*2,
		mpb.BarStyle(),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f"),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name(" ] "),
			decor.EwmaETA(decor.ET_STYLE_GO, 0),
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 0, decor.WCSyncWidth),
			//decor.AverageSpeed(decor.UnitKiB, "% .2f", decor.WCSyncWidth),
		),
	)

	// file to reader
	open, err := os.Open("./tmp.txt")
	if err != nil {
		panic(err)
	}
	stat, err := open.Stat()
	if err != nil {
		panic(err)
	}
	// set total as file size
	fmt.Println("size: ", stat.Size())
	bar.SetTotal(stat.Size(), false)

	// create proxy reader
	proxyReader := bar.ProxyReader(open)
	//defer proxyReader.Close()
	fmt.Println("before copy")
	// copy from proxyReader, ignoring errors
	create, err := os.Create("./tmp2.txt")
	if err != nil {
		panic(err)
	}
	defer create.Close()

	n, _ := io.Copy(create, proxyReader)
	fmt.Println("read number:", n)

	// triggering complete event now
	bar.SetTotal(-1, true)
	fmt.Println("copy finish!")
	fmt.Println("copy finish!")

}

func singleBar() {
	// initialize progress container, with custom width
	p := mpb.New(mpb.WithWidth(64))

	total := 100
	name := "Single Bar:"
	// create a single bar, which will inherit container's width
	bar := p.New(int64(total),
		// BarFillerBuilder with custom style
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
		mpb.PrependDecorators(
			// display our name with one space on the right
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			// replace ETA decorator with "done" message, OnComplete event
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), "done",
			),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	// simulating some work
	max := 100 * time.Millisecond
	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(mathrand.Intn(10)+1) * max / 10)
		bar.Increment()
	}
	// wait for our bar to complete and flush
	p.Wait()
}
