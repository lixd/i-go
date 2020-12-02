package main

import (
	"encoding/xml"
	"github.com/sittipongwork/GoSitemapxml"
	"os"
)

func main() {
	Generate()
}

func Generate() {
	// Create Sitemap Container for keep data xml
	sitemap := gositemap.CreateSitemapContainer("https://www.sitemaps.org/schemas/sitemap/0.9")

	// Add Site you want to show in sitemap.xml
	sitemap.AddSite("https://website.com/", "2016-11-23")
	// You can add more one site : Example
	sitemap.AddSite("https://website.com/blog/1", "2016-11-23")
	sitemap.AddSite("https://website.com/blog/2", "2016-11-23")
	sitemap.AddSite("https://website.com/blog/3", "2016-11-23")
	sitemap.AddSite("https://website.com/contact", "2016-11-23")

	// Mashal XML Data
	output, _ := xml.MarshalIndent(sitemap, "  ", "    ")
	// Print Output
	os.Stdout.Write(output)

	// NOTE : in iris framework
	// you can run command , Dont MashalIndent
	// iris.XML(iris.StatusOK, sitemap)
}
