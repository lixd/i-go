package sync_pool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"sync"
	"testing"
	"time"
)

/*
sync.pool 使用 https://zhuanlan.zhihu.com/p/133638023
我们发现 Get 方法取出来的对象和上次 Put 进去的对象实际上是同一个，Pool 没有做任何“清空”的处理。但我们不应当对此有任何假设，
因为在实际的并发使用场景中，无法保证这种顺序，最好的做法是在 Put 前，将对象清空。
局部变量使用 sync.pool 可以降低分配压力，如果新建一个对象比较耗时则可以使用
非局部变量使用 sync.pool 可以降低分配压力+gc压力，比如 gin 框架，会给每个请求都创建一个 context，就是用的 sync.pool 进行复用。
*/
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func (s *Student) Reset() {
	s.Age = 0
	s.Name = ""
	s.Remark = [1024]byte{}
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var defaultStu, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
var emptyStu = Student{}

func BenchmarkUnmarshal(b *testing.B) {
	expire, _ := time.Parse("2006-01-02 15:04:05", "2021-10-19 12:50:43")
	fmt.Println(expire.Unix())
	return
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(defaultStu, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(defaultStu, stu)
		stu.Reset() // 将stu重置后再放进去
		studentPool.Put(stu)
	}
}

func TestA(t *testing.T) {
	// 103.101.225.229
	// ip, err := net.LookupIP("lee-seung-yeoul-0-korea-selatan-2010-dd-s8.utsmakassar.web.id")
	ip, err := net.LookupIP("perkuliahan-siang-pmb-kutaitimur.peta-lokasi.web.id")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ip)

	rawURL := "https://www.vaptcha.com/document/install.html"
	parse, err := url.Parse(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Home:", parse.Scheme+"://"+parse.Host)
	fmt.Println(parse.Path)
	day := time.Unix(1634790080, 1).Day()
	day2 := time.Now().Day()
	fmt.Printf("day1:%v day2:%v\n", day, day2)
}
