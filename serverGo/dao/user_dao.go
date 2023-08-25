package dao

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/pkg/errors"
	"serverGo/foundation/mysql"
	"serverGo/foundation/utils"
	"serverGo/model"
	"time"
)

func Login(user *model.User) (err error) {
	passwordParam := user.Password
	keyParamParam := user.Keyword
	sqlStr := "SELECT id, username, nickname, password, keyword, email, status, create_time, last_success_login_time, last_try_login_time FROM new_user WHERE username = ?"
	var userEntity model.UserEntity
	err = mysql.DB.Get(&userEntity, sqlStr, user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("该用户不存在")
		}
		return err
	}
	TranUserEntity(userEntity, user)
	//encryptPasswordParam := encryptPassword([]byte(passwordParam))
	encryptPasswordParam := passwordParam // 暂不加密 以后再改
	if encryptPasswordParam != user.Password {
		return errors.New("密码错误")
	}
	if keyParamParam != user.Keyword {
		return errors.New("口令错误")
	}
	sqlStr = "UPDATE new_user SET last_success_login_time = NOW(), last_try_login_time = NOW() WHERE id = ?"
	_, _ = mysql.DB.Exec(sqlStr, user.ID)
	return nil
}

func GetUserInfo(uId uint64) (user *model.User, err error) {
	sqlStr := "SELECT id, username,nickname, password, keyword, email, status, create_time, last_success_login_time, last_try_login_time FROM new_user WHERE id = ?"
	var userEntity model.UserEntity
	err = mysql.DB.Get(&userEntity, sqlStr, uId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("该用户不存在")
		}
		return user, err
	}
	user = &model.User{}
	TranUserEntity(userEntity, user)
	return
}

// encryptPassword 对密码进行加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte("sys-secret-Default"))
	return hex.EncodeToString(h.Sum(data))
}

func TranUserEntity(entity model.UserEntity, user *model.User) {
	user.ID = entity.ID
	user.Username = entity.Username
	user.Password = entity.Password
	user.Keyword = entity.Keyword
	if entity.Nickname.Valid {
		user.Nickname = entity.Nickname.String
	}
	if entity.Email.Valid {
		user.Email = entity.Email.String
	}
	user.Gender = entity.Gender
	user.Status = entity.Status
	if entity.CreateTime.Valid {
		user.CreateTime = utils.LocalTime(entity.CreateTime.Time)
	}
	if entity.LastSuccessLoginTime.Valid {
		user.LastSuccessLoginTime = utils.LocalTime(entity.LastSuccessLoginTime.Time)
	} else {
		user.CreateTime = utils.LocalTime(time.Time{}) // 通过 IsZero() 来判断是否是零值
	}
	if entity.LastTryLoginTime.Valid {
		user.LastTryLoginTime = utils.LocalTime(entity.LastTryLoginTime.Time)
	}
}
