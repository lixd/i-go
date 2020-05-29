package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"i-go/core/db/mongodb"
	"i-go/demo/rbac/model"
)

type rbacModule struct {
	coll *mongo.Collection
}

var RbacModule = &rbacModule{}

func (c *rbacModule) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.Module))
	}
	return c.coll
}

func (c *rbacModule) InsertMany(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
func (c *rbacModule) Query(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
