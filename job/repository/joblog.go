package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/core/db/mongodb"
	"i-go/job/model"
	"i-go/utils"
	"i-go/utils/itime"
)

type logRepository struct {
	coll *mongo.Collection
}

var Log = logRepository{}

const (
	_ = iota
	Ready
	Working
	Failed
)

func (c *logRepository) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.JobLog))
	}
	return c.coll
}

func (c *logRepository) InsertJobLog(log *model.JobLog) error {
	log.Id = primitive.NewObjectID()
	_, err := c.GetColl().InsertOne(nil, log)
	if err != nil {
		logrus.Error("service log error,", err)
	}
	return err
}
func (c *logRepository) UpdateEndTime(id string, end int64) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{
		{"_id", _id},
	}
	update := bson.D{
		{"EndTime", end},
	}
	_, err = c.GetColl().UpdateOne(context.Background(), filter, update)
	return err
}

// ClearLogs 删除50条之前的日志
func (c *logRepository) ClearLogs(names []string, t time.Time) {

	for _, name := range names {
		var filter = bson.D{
			{"JobName", name},
		}
		var sort = bson.D{
			{"_id", -1},
		}
		var opts = options.FindOne().SetSort(sort).SetSkip(50)
		doc := c.GetColl().FindOne(nil, filter, opts)
		var m model.JobLog
		if err := doc.Decode(&m); err != nil {
			logrus.Error(" JobName: error:", name, err)
			continue
		}
		filter = bson.D{
			{"JobName", name},
			{"$lte", bson.D{
				{"_id", m.Id},
			}},
		}
		logrus.Info("task: ", name, m.Id.Hex())
		_, _ = c.GetColl().DeleteMany(nil, filter)
	}
}

func (c *logRepository) FindPage(filter bson.D, skip int64, limit int64) (result []model.JobLog, total int64, err error) {
	if filter == nil {
		filter = bson.D{}
	}
	sort := bson.D{
		{"StartTime", -1},
	}
	cursor, err := c.GetColl().Find(context.Background(), filter, options.Find().SetSort(sort).SetSkip(skip).SetLimit(limit))
	if err != nil {
		result = make([]model.JobLog, 0)
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller()}).Warn(err)
		return
	}
	total, err = c.GetColl().CountDocuments(context.Background(), filter)
	defer cursor.Close(context.Background())
	if err != nil {
		result = make([]model.JobLog, 0)
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller()}).Warn(err)
		return
	}
	for cursor.Next(context.Background()) {
		var log model.JobLog
		err := cursor.Decode(&log)
		if err != nil {
			logrus.WithFields(logrus.Fields{"Caller": utils.Caller()}).Warn(err)
			continue
		}
		result = append(result, log)
	}
	return
}

func (c *logRepository) FindError() (total int64, err error) {
	start := time.Now().Unix()
	end := itime.GetZeroTime(time.Now().Unix())
	filter := bson.D{
		{"StartTime", bson.D{
			{"$gte", start},
			{"$lte", end},
		}},
	}
	total, err = c.GetColl().CountDocuments(context.Background(), filter)
	if err != nil {
		total = 0
		return
	}
	return
}

func (c *logRepository) CleanWeekBeforeLog() {
	start := itime.GetZeroTime(time.Now().AddDate(0, 0, -6).Unix())
	filter := bson.D{
		{"StartTime", bson.D{
			{"$lte", start},
		}},
	}
	_, err := c.GetColl().DeleteMany(context.Background(), filter)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller()}).Warn(err)
	}
}
