package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// linePoints()
	histogram()
}

func linePoints() {
	// new一个实例
	p, _ := plot.New()
	// 填充标题和XY轴图例
	p.Title.Text = "Hello Price"
	p.X.Label.Text = "Quantity Demand"
	p.Y.Label.Text = "Price"
	// 随便加一些mock数据
	points := plotter.XYs{
		{2.0, 60000.0},
		{4.0, 40000.0},
		{6.0, 30000.0},
		{8.0, 25000.0},
		{10.0, 23000.0},
	}
	// 添加到plot
	// 要画多个线分别添加即可
	_ = plotutil.AddLinePoints(p, points)
	// 然后设定图片大小和保存的位置 支持通过后缀调整文件格式
	_ = p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
func histogram() {
	// new一个实例
	p, _ := plot.New()
	// 填充标题和XY轴图例
	p.Title.Text = "Hello Price"
	p.X.Label.Text = "Quantity Demand"
	p.Y.Label.Text = "Price"
	// 随便加一些mock数据
	points := plotter.XYs{
		{2.0, 60000.0},
		{4.0, 40000.0},
		{6.0, 30000.0},
		{8.0, 25000.0},
		{10.0, 23000.0},
	}
	// 竖状图
	h, err := plotter.NewHistogram(points, points.Len())
	if err != nil {
		panic(err)
	}
	p.Add(h)
	// 然后设定图片大小和保存的位置 支持通过后缀调整文件格式
	_ = p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
