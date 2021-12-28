package sitemap

import (
	"time"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
)

func GenerateSitemap(path string) {
	sm := stm.NewSitemap(1)
	// Your website's host name
	sm.SetDefaultHost("https://www.lixueduan.com")
	sm.SetVerbose(true)
	sm.SetCompress(false)
	sm.SetPretty(true)
	sm.SetFilename("sitemap")
	sm.SetPublicPath(path) // sitemap 生成文件输出路径，会再该路径下生成 /sitemaps 目录
	sm.Create()
	date := time.Now().Format("2006-01-02")
	url := stm.URL{{"loc", "https://www.lixueduan.com"}, {"lastmod", date}}
	sm.Add(url)
	sm.Finalize()
}
