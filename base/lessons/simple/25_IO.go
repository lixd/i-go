package simple

import (
	"bufio"
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {
	//searchRepeat()
	GetUrl()
}
func searchRepeat() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
		if line == "A" {
			break
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s \n", line, n)
		}
	}
}
func GetUrl() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("resp %v", resp)
	ContentType := resp.Header.Get("Content-Type")
	fmt.Printf("Content-Type %v", ContentType)
	html.Parse(resp.Body)

}

type dollars float32
type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Printf("item:%v price:%v", item, price)
	}
	s := req.URL.Path
	fmt.Printf("url:%v", s)

}
