package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"i-go/core/db/mongodb"
	"i-go/demo/rbac/model"
)

type rbacModuleRole struct {
	coll *mongo.Collection
}

var RbacModuleRole = &rbacModuleRole{}

func (c *rbacModuleRole) GetColl() *mongo.Collection {
	if c.coll == nil {
		c.coll = mongodb.GetJobCollection(new(model.UserRole))
	}
	return c.coll
}

func (c *rbacModuleRole) InsertMany(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
func (c *rbacModuleRole) Query(docs []interface{}) error {
	_, err := c.GetColl().InsertMany(nil, docs)
	return err
}
