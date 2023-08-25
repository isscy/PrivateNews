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
type User struct {
	ID                   uint64          `json:"id" db:"id"`
	Username             string          `json:"username" db:"username"`
	Nickname             string          `json:"nickname" db:"nickname"`
	Password             string          `json:"password" db:"password"`
	Keyword              string          `json:"keyword" db:"keyword"`
	Email                string          `json:"email" db:"email"`            // 邮箱
	Gender               int             `json:"gender" db:"gender"`          // 性别
	Status               string          `json:"status" db:"status"`          // 1-启用 2-禁用
	CreateTime           utils.LocalTime `json:"createTime" db:"create_time"` // 创建日期
	LastSuccessLoginTime utils.LocalTime `json:"lastSuccessLoginTime" db:"last_success_login_time"`
	LastTryLoginTime     utils.LocalTime `json:"lastTryLoginTime" db:"last_try_login_time"`
	/*CreateTime           time.Time `json:"createTime" db:"create_time"` // 创建日期
	LastSuccessLoginTime time.Time `json:"lastSuccessLoginTime" db:"last_success_login_time"`
	LastTryLoginTime     time.Time `json:"lastTryLoginTime" db:"last_try_login_time"`*/

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

/*
=========================================================================================================================
====================================== 				DB映射类						=============================================
==========================================================================================================================
*/
type UserEntity struct {
	ID                   uint64         `db:"id"`
	Username             string         `db:"username"`
	Nickname             sql.NullString `db:"nickname"` // 邮箱
	Password             string         `db:"password"`
	Keyword              string         `db:"keyword"`
	Email                sql.NullString `db:"email"`       // 邮箱
	Gender               int            `db:"gender"`      // 性别
	Status               string         `db:"status"`      // 1-启用 2-禁用
	CreateTime           sql.NullTime   `db:"create_time"` // 创建日期
	LastSuccessLoginTime sql.NullTime   `db:"last_success_login_time"`
	LastTryLoginTime     sql.NullTime   `db:"last_try_login_time"`
}

/*=========================================================================================================================
====================================== 				Form					=============================================
==========================================================================================================================*/
// LoginForm 登录请求参数
type LoginForm struct {
	UserName string `json:"username" binding:"required" msg:"用户名必填"`
	Password string `json:"password" binding:"required" msg:"密码必填"`
	Keyword  string `json:"keyword" binding:"required" msg:"口令必填"` // 口令
}

// RegisterForm 注册请求参数
type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`  // 用户名
	Email           string `json:"email" binding:"required"`     // 邮箱
	Gender          int    `json:"gender" binding:"oneof=0 1 2"` // 性别 0:未知 1:男 2:女
	Password        string `json:"password" binding:"required"`  // 密码
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
