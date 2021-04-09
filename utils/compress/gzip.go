package compress

// 数据压缩 使用标准库gzip包实现
import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Compress 压缩
func Compress(src []byte) []byte {
	// 自定义一个io.Writer传入并返回一个io.Writer
	// 将待压缩数据写入返回的io.Writer最后从传入的io.Writer中读取压缩后的数据
	buf := new(bytes.Buffer)
	zw := gzip.NewWriter(buf)
	zw.Write(src)
	zw.Flush()
	zw.Close() // Close不能用defer 必须在buf.Bytes()取值之前执行
	return buf.Bytes()
}

// DeCompress 解压
func DeCompress(compressed []byte) ([]byte, error) {
	// 将压缩后的数据当做io.Reader传入并从返回的io.Reader中去读解压后的数据
	buf := bytes.NewBuffer(compressed)
	zr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	return ioutil.ReadAll(zr)
}
