Google https://developers.google.com/search/docs/advanced/sitemaps/overview
Baidu https://ziyuan.baidu.com/college/courseinfo?id=267&page=2#h2_article_title2
一个Sitemap文件包含的网址不得超过 5 万个，且文件大小不得超过 10 MB。如果您的Sitemap超过了这些限值，请将其拆分为几个小的Sitemap。
```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- XML文件需以utf-8编码-->
<urlset>
<!--必填标签-->
    <url>
        <!--必填标签,这是具体某一个链接的定义入口，每一条数据都要用<url>和</url>包含在里面，这是必须的 -->
        <loc>http://www.yoursite.com/yoursite.html</loc>
        <!--必填,URL链接地址,长度不得超过256字节-->
        <lastmod>2009-12-14</lastmod>
        <!--可以不提交该标签,用来指定该链接的最后更新时间-->
        <changefreq>daily</changefreq>
        <!--可以不提交该标签,用这个标签告诉此链接可能会出现的更新频率 -->
        <priority>0.8</priority>
        <!--可以不提交该标签,用来指定此链接相对于其他链接的优先权比值，此值定于0.0-1.0之间-->
    </url>
    <url>
        <loc>http://www.yoursite.com/yoursite2.html</loc>
        <lastmod>2010-05-01</lastmod>
        <changefreq>daily</changefreq>
        <priority>0.8</priority>
    </url>
</urlset>
```
自适应网页
```xml
<mobile:mobile type="pc,mobile"/>
```

