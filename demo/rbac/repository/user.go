package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"i-go/core/db/mongodb"
	"i-go/demo/rbac/model"
)

type rbacUser struct {
	coll *mongo.Collection
}

var RbacUser = &rbacUser{}

func (c *rbacUser) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.User))
	}
	return c.coll
}

func (c *rbacUser) InsertMany(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}

func (c *rbacUser) Query(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
