package model

import (
	"database/sql"
	"serverGo/foundation/utils"
)

/*
=========================================================================================================================
====================================== 				实体类						=============================================
==========================================================================================================================
*/
type BuiltInTag struct {
	ID         uint64           `json:"id" db:"id"`
	TagKey     string           `json:"tagKey" db:"tag_key"`
	TagName    string           `json:"tagName" db:"tag_name"`
	Status     string           `json:"status" db:"status"` // 1-启用 2-禁用
	SortNo     int64            `json:"sortNo" db:"sort_no"`
	CreateTime utils.LocalTime  `json:"createTime" db:"create_time"` // 创建日期
	Items      []BuiltInTagItem `json:"items"`                       // 子表
}

type BuiltInTagItem struct {
	ID         uint64          `json:"id" db:"id"`
	TagId      string          `json:"tagId" db:"tag_id"`
	TagKey     uint64          `json:"tagKey" db:"tag_key"` // 本类下唯一标识
	TagName    string          `json:"tagName" db:"tag_name"`
	Status     string          `json:"status" db:"status"` // 1-启用 2-禁用
	SortNo     int64           `json:"sortNo" db:"sort_no"`
	CreateTime utils.LocalTime `json:"createTime" db:"create_time"` // 创建日期
}

/*
=========================================================================================================================
====================================== 				DB映射类						=============================================
==========================================================================================================================
*/
type BuiltInTagEntity struct {
	ID         uint64        `db:"id"`
	TagKey     string        `db:"tag_key"`
	TagName    string        `db:"tag_name"`
	Status     string        `db:"status"` // 1-启用 2-禁用
	SortNo     sql.NullInt64 `db:"sort_no"`
	CreateTime sql.NullTime  `db:"create_time"` // 创建日期
}
type BuiltInTagItemEntity struct {
	ID         uint64        `json:"id" db:"id"`
	TagId      string        `json:"tagId" db:"tag_id"`
	TagKey     uint64        `json:"tagKey" db:"tag_key"` // 本类下唯一标识
	TagName    string        `json:"tagName" db:"tag_name"`
	Status     string        `json:"status" db:"status"` // 1-启用 2-禁用
	SortNo     sql.NullInt64 `json:"sortNo" db:"sort_no"`
	CreateTime sql.NullTime  `json:"createTime" db:"create_time"` // 创建日期
}

/*=========================================================================================================================
====================================== 				Form					=============================================
==========================================================================================================================*/
// BuiltInTagForm 查询请求参数
type BuiltInTagForm struct {
	TagKey  string `json:"tagKey"`
	TagName string `json:"tagName"`
	Status  string `json:"status"` // 口令
	TagId   string `json:"tagId"`  // 口令
}
