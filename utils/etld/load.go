package etld

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

// 利用 embed 在编译时将数据嵌入到变量中
// 然后用该变量初始化map
// 数据来源 https://publicsuffix.org/list/effective_tld_names.dat

//go:embed tldomains.dat
var tldomains string

func init() {
	reader := strings.NewReader(tldomains)
	br := bufio.NewReader(reader)
	for {
		l, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(l)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		TLDs[line] = struct{}{}
	}
	fmt.Println("TLD个数:", len(TLDs))
}
