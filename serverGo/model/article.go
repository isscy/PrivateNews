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
type Article struct {
	ID                int64           `json:"id" db:"id"`
	Title             string          `json:"title" db:"title"`
	Topic             string          `json:"topic" db:"topic"`
	Content           string          `json:"content" db:"content"`
	Tags              string          `json:"tags" db:"tags"`
	International     int             `json:"international" db:"international"`
	Region            int             `json:"region" db:"region"`
	EventType         int             `json:"eventType" db:"event_type"`
	Status            string          `json:"status" db:"status"`          // 1-启用 2-禁用
	CreateTime        utils.LocalTime `json:"createTime" db:"create_time"` // 创建日期
	InternationalName string          `json:"internationalName" db:"internationalName"`
	RegionName        string          `json:"regionName" db:"regionName"`
	EventTypeName     string          `json:"eventTypeName" db:"event_typeName"`
}

/*
=========================================================================================================================
====================================== 				DB映射类						=============================================
==========================================================================================================================
*/
type ArticleEntity struct {
	ID                int64        `db:"id"`
	Title             string       `db:"title"`
	Topic             string       `db:"topic"`
	Content           string       `db:"content"`
	Tags              string       `db:"tags"`
	International     int          `db:"international"`
	InternationalName string       `db:"internationalName"`
	Region            int          `db:"region"`
	RegionName        string       `db:"regionName"`
	EventType         int          `db:"event_type"`
	EventTypeName     string       `db:"event_typeName"`
	Status            string       `db:"status"`      // 1-启用 2-禁用
	CreateTime        sql.NullTime `db:"create_time"` // 创建日期
}

/*=========================================================================================================================
====================================== 				Form					=============================================
==========================================================================================================================*/
// 新增时候的Form
type ArticleForm struct {
	ID            uint64   `json:"id" db:"id"`
	Title         string   `json:"title" db:"title" binding:"required" msg:"标题必填"`
	Topic         string   `json:"topic" db:"topic" binding:"required" msg:"主题必填"`
	Content       string   `json:"content" db:"content" binding:"required" msg:"内容必填"`
	TagList       []string `json:"tagList" db:"tags"`
	International int      `json:"international" db:"international" `
	Region        int      `json:"region" db:"region"`
	EventType     int      `json:"eventType" db:"event_type"`
}

// 查询时候的form
type ArticleSearchForm struct {
	Title         string `json:"title"`
	Topic         string `json:"topic"`
	Content       string `json:"content"`
	Tag           string `json:"tags"`
	International int    `json:"international"`
	Region        int    `json:"region"`
	EventType     int    `json:"eventType"`
	Status        string `json:"status"` // 1-启用 2-禁用
	CurrentPage   int    `json:"currentPage"`
	PageSize      int    `json:"pageSize"`
	SearchValue   string `json:"searchValue"`
}
