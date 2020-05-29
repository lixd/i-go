package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// module 模块 权限
type Module struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"Name"`
	Parent string             `bson:"Parent"`

	CreateTime time.Duration `bson:"CreateTime"`
	UpdateTime time.Duration `bson:"UpdateTime"`
}

func (*Module) GetCollectionName() string {
	return "V_Module"
}

// User 用户
type User struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"Name"`
	Password   string             `bson:"Password"`
	Role       []string           `bson:"Role"`
	CreateTime time.Duration      `bson:"CreateTime"`
	UpdateTime time.Duration      `bson:"UpdateTime"`
}

func (*User) GetCollectionName() string {
	return "V_User"
}

// Role 角色
type Role struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"Name"`
	CreateTime time.Duration      `bson:"CreateTime"`
	UpdateTime time.Duration      `bson:"UpdateTime"`
}

func (*Role) GetCollectionName() string {
	return "V_Role"
}

// ModuleRole 权限-角色
type ModuleRole struct {
	Id         primitive.ObjectID `bson:"_id"`
	ModuleId   string             `bson:"ModuleId"`
	RoleId     string             `bson:"RoleId"`
	Name       string             `bson:"Name"`
	CreateTime time.Duration      `bson:"CreateTime"`
	UpdateTime time.Duration      `bson:"UpdateTime"`
}

func (*ModuleRole) GetCollectionName() string {
	return "V_ModuleRole"
}

// UserRole 用户-角色
type UserRole struct {
	Id         primitive.ObjectID `bson:"_id"`
	UserId     string             `bson:"UserId"`
	RoleId     string             `bson:"RoleId"`
	Code       string             `bson:"Code"`
	Name       string             `bson:"Name"`
	CreateTime time.Duration      `bson:"CreateTime"`
	UpdateTime time.Duration      `bson:"UpdateTime"`
}

func (*UserRole) GetCollectionName() string {
	return "V_UserRole"
}

// LoginLog 登录日志
type LoginLog struct {
	Id         primitive.ObjectID `bson:"_id"`
	UserId     string             `bson:"UserId"`
	UserName   string             `bson:"UserName"`
	IP         string             `bson:"Code"`
	CreateTime time.Duration      `bson:"CreateTime"`
}

func (*LoginLog) GetCollectionName() string {
	return "V_LoginLog"
}
