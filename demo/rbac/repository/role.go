package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"i-go/core/db/mongodb"
	"i-go/demo/rbac/model"
)

type rbacRole struct {
	coll *mongo.Collection
}

var RbacRole = &rbacRole{}

func (c *rbacRole) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.Role))
	}
	return c.coll
}

func (c *rbacRole) InsertMany(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
func (c *rbacRole) Query(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
