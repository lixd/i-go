package main

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	path := "D:/wlinno/document/test2.xlsx"
	excelizeExcel(path)
	//xlxs(path)
}

const (
	ProductID     = "商品id"
	ProductName   = "商品名称"
	Img           = "商品主图"
	OriginPrice   = "商品价格(单位：元)"
	DiscountPrice = "优惠券面额"
	Validity      = "优惠券结束时间"
	Link          = "优惠券链接"
	// 这个需要前端传
	DefaultAdId = "DefaultAdId"
)

type AdTbInfo struct {
	//Id            primitive.ObjectID `bson:"_id"`
	AdId          string  `bson:"AdId"`          // 关联 当前记录具体属于哪个广告
	Img           string  `bson:"Img"`           // 图片链接
	Link          string  `bson:"Link"`          // 广告链接(优惠券链接)
	ProductId     string  `bson:"ProductId"`     // 商品ID
	ProductName   string  `bson:"ProductName"`   // 商品名称
	OriginalPrice float64 `bson:"OriginalPrice"` // 原价
	DiscountPrice float64 `bson:"DiscountPrice"` // 折扣价
	Validity      int64   `bson:"Validity"`      // 有效期(活动截止日期)

	CreateTime int64 `bson:"CreateTime"`
	UpdateTime int64 `bson:"UpdateTime"`
}

func excelizeExcel(path string) {
	var (
		columnMap = make(map[string]int)
		list      = make([]AdTbInfo, 0)
	)

	// 读取excel 只支持xlsx文件
	//xlsx, err := excelize.OpenFile(path)
	file, err := os.Open(path)
	all, err := ioutil.ReadAll(file)
	reader := bytes.NewReader(all)
	xlsx, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	sheetMap := xlsx.GetSheetMap()
	fmt.Println("xlsx.Sheet", sheetMap)
	for _, v := range sheetMap {
		rows := xlsx.GetRows(v)
		//fmt.Println("rows", rows)
		// 定位标题 需要的数据在多少列
		titleRow := rows[0:1]
		for _, row := range titleRow {
			for index, colCell := range row {
				locationTitle(&columnMap, colCell, index)
				fmt.Print(colCell, "\t")
			}
		}
		fmt.Print(columnMap, "\t")
		// 拼接数据
		dataRow := rows[1:]
		for _, row := range dataRow {
			originalPrice, err := strconv.ParseFloat(row[columnMap[OriginPrice]], 64)
			if err != nil {
				continue
			}
			discountPrice, err := strconv.ParseFloat(row[columnMap[DiscountPrice]], 64)
			if err != nil {
				continue
			}
			validity, err := time.Parse("2006-01-02 15:04:05", row[columnMap[Validity]])
			if err != nil {
				continue
			}
			var item = AdTbInfo{
				AdId:          DefaultAdId,
				Img:           row[columnMap[Img]],
				Link:          row[columnMap[Link]],
				ProductId:     row[columnMap[ProductID]],
				ProductName:   row[columnMap[ProductName]],
				OriginalPrice: originalPrice,
				DiscountPrice: originalPrice - discountPrice, // 券后价
				Validity:      validity.Unix(),
				CreateTime:    -1,
				UpdateTime:    -1,
			}
			fmt.Println(item.ProductId, "\t")
			list = append(list, item)

			fmt.Println(item, "\t")
		}
	}
	fmt.Println(list, "\t")
}

// locationTitle 根据标题名字找到具体类型在哪一列
func locationTitle(columnMap *map[string]int, colCell string, index int) {
	switch colCell {
	case ProductID:
		(*columnMap)[ProductID] = index
	case ProductName:
		(*columnMap)[ProductName] = index
	case Img:
		(*columnMap)[Img] = index
	case OriginPrice:
		(*columnMap)[OriginPrice] = index
	case DiscountPrice:
		(*columnMap)[DiscountPrice] = index
	case Validity:
		(*columnMap)[Validity] = index
	case Link:
		(*columnMap)[Link] = index
	}
}

func xlxs(path string) {
	xlFile, err := xlsx.OpenFile(path)
	//xlFile, err := xlsx.OpenBinary()
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	} else {
		fmt.Printf("excel: %#v\n", xlFile)

	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
