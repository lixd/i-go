package repository

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/db/mongodb"
	"i-go/job/constant"
	"i-go/job/model"
	"time"
)

type jobRepository struct {
	coll *mongo.Collection
}

var Job = jobRepository{}

func (c *jobRepository) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.JobInfo))
	}
	return c.coll
}

// Init
func (c *jobRepository) Init(taskCode, name, spec, desc string, duration int64) (model.JobInfo, error) {
	var info model.JobInfo

	var filter = bson.D{
		{"JobCode", taskCode},
	}
	var update = bson.D{
		{"$set", bson.D{
			{"JobName", name},
			{"CronSpec", spec},
			{"Desc", desc},
			{"MinDuration", duration},
			{"Status", jobstatus.Wait},
			{"CreateTime", time.Now().Unix()},
		}},
	}
	var opts = options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err := c.GetColl().FindOneAndUpdate(nil, filter, update, opts).Decode(&info)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "init job error"})
	}
	return info, nil
}

// FindAll 查询所有任务列表
func (c *jobRepository) FindAll() (result []model.JobInfo, err error) {
	filter := bson.D{}
	cursor, err := c.GetColl().Find(nil, filter)
	if err != nil {
		result = make([]model.JobInfo, 0)
		return
	}
	defer cursor.Close(nil)
	for cursor.Next(nil) {
		var jobInfo model.JobInfo
		err := cursor.Decode(&jobInfo)
		if err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "decode job error"})
			continue
		}
		result = append(result, jobInfo)
	}
	if result == nil {
		result = make([]model.JobInfo, 0)
	}
	return
}
