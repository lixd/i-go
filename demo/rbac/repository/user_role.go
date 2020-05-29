package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"i-go/core/db/mongodb"
	"i-go/demo/rbac/model"
)

type rbacUserRole struct {
	coll *mongo.Collection
}

var RbacUserRole = &rbacUserRole{}

func (c *rbacUserRole) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.UserRole))
	}
	return c.coll
}

func (c *rbacUserRole) InsertMany(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
func (c *rbacUserRole) Query(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
