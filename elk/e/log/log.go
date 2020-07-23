package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-go/core/conf"
	"i-go/core/db/elasticsearch"
	"i-go/utils"
	"io"
	"os"
	"strconv"
	"strings"
)

var es *elastic.Client

const PathNormal = "D:/wlinno/document/featuremodel/样本/4.29 正常轨迹10W样本.log"

func Init() {
	err := conf.Init("./conf/config.yml")
	if err != nil {
		panic(err)
	}
	elasticsearch.Init()
	es = elasticsearch.ESClient
}

func main() {
	Init()
	records := LoadFromFile(PathNormal)
	Insert2ESBulk(records)
}

func LoadFromFile(path string) []PathRecord {
	var records = make([]PathRecord, 0, 100000)
	open, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(open)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("read err: ", err)
				continue
			}
		}
		// 去掉"\n"
		line = strings.ReplaceAll(line, "\n", "")
		var r PathRecord
		err = json.Unmarshal([]byte(line), &r)
		if err != nil {
			fmt.Println("Unmarshal err: ", err)
		} else {
			records = append(records, r)
		}
	}
	return records
}

// Upsert
func Upsert(body *PathRecord) error {
	_, err := es.Update().
		Index("path_v1").
		Id(body.Id.Hex()).
		Doc(&body).
		// ES 检测到 文档无变化则不执行任何操作
		DetectNoop(true).
		// 文档不存在则创建
		DocAsUpsert(true).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"CallerName": utils.Caller(), "Scenes": "ES:Index"}).Error(err)
	}
	return err
}

func Insert2ESBulk(list []PathRecord) {
	var bulkReq *elastic.BulkIndexRequest
	var count int
	var bulk *elastic.BulkService
	bulk = es.Bulk()
	for i, v := range list {
		if v.Vid == "" {
			fmt.Println("记录为空")
			continue
		}
		if count > 500 {
			_, err := bulk.Do(context.Background())
			if err != nil {
				logrus.WithFields(logrus.Fields{"CallerName": utils.Caller(), "Scenes": "写入失败"}).Error(err)
			}
			fmt.Println("成功写入 500 条记录")
			count = 0
			bulk = es.Bulk()
		}
		bulkReq = elastic.NewBulkIndexRequest().Index("path_v1").Id(strconv.Itoa(i)).Doc(&v)
		bulk.Add(bulkReq)
		count++
	}
}

type PathRecord struct {
	Id                 primitive.ObjectID `bson:"_id"`
	Vid                string             `bson:"Vid"`
	CellId             string             `bson:"CellId"`
	DrawTimeString     string             `bson:"DrawTimeString"`
	Device             string             `bson:"Device"`
	DF                 string             `bson:"DF"`
	IP                 string             `bson:"IP"`
	VerifyPath         string             `bson:"VerifyPath"`
	DateTime           int64              `bson:"DateTime"`
	LocusId            string             `bson:"LocusId"`
	AdId               string             `bson:"AdId"`
	ImgType            int                `bson:"ImgType"` //
	Similarity         float64            `bson:"Similarity"`
	FeatureModelScore  int                `bson:"FeatureModelScore"`
	ValidateResult     bool               `bson:"ValidateResult"`
	SimilarityPassLine int                `bson:"SimilarityPassLine"`
	FeaturePassLine    int                `bson:"FeaturePassLine"`
	DeductionScore     float64            `bson:"DeductionScore"`
	Frequency          float64            `bson:"Frequency"`
	ValidateKnock      string             `bson:"ValidateKnock"`
	Url                string             `bson:"Url"`
	Referer            string             `bson:"Referer"`
	UserAgent          string             `bson:"UserAgent"`
	Cookie             string             `bson:"Cookie"`
}
