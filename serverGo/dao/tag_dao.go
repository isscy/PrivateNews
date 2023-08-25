package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"serverGo/foundation/mysql"
	"serverGo/foundation/utils"
	"serverGo/model"
)

func GetBuiltInTagList(formParam *model.BuiltInTagForm) (list []model.BuiltInTag, err error) {
	var args []string = make([]string, 4)
	sqlStr := "SELECT id,tag_key,tag_name,`status`,sort_no, create_time  FROM new_built_in_tag a WHERE 1=1 "
	if formParam.TagName != "" {
		sqlStr = sqlStr + " AND a.tag_name =  " + formParam.TagName
		args = append(args, formParam.TagName)
	}
	if formParam.TagKey != "" {
		sqlStr = sqlStr + " AND a.tag_key = ? " + formParam.TagKey
		args = append(args, formParam.TagKey)
	}

	var tagEntityList []model.BuiltInTagEntity // = make([]model.BuiltInTagEntity, 0)
	err = mysql.DB.Select(&tagEntityList, sqlStr)
	if err != nil {
		fmt.Println(errors.WithStack(err))
		if err == sql.ErrNoRows {
			return list, errors.New("该Tag不存在")
		}
		return list, err
	}
	var tagList = make([]model.BuiltInTag, 0)
	if len(tagEntityList) > 0 {
		for _, value := range tagEntityList {
			sqlStr = "SELECT *  FROM new_built_in_tag_item WHERE tag_id = ?"
			var tagItemEntityList []model.BuiltInTagItemEntity
			err = mysql.DB.Select(&tagItemEntityList, sqlStr, value.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					continue
				}
				return list, err
			}
			tag := model.BuiltInTag{}
			TranTagEntity(value, &tag)
			var tagItemList = make([]model.BuiltInTagItem, 0)
			for _, subValue := range tagItemEntityList {
				tagItem := model.BuiltInTagItem{}
				TranTagItemEntity(subValue, &tagItem)
				tagItemList = append(tagItemList, tagItem)
			}
			tag.Items = tagItemList
			tagList = append(tagList, tag)
		}
	}
	return tagList, err
}

func TranTagEntity(entity model.BuiltInTagEntity, tag *model.BuiltInTag) {
	tag.ID = entity.ID
	tag.TagKey = entity.TagKey
	tag.TagName = entity.TagName
	tag.Status = entity.Status
	if entity.SortNo.Valid {
		tag.SortNo = entity.SortNo.Int64
	}
	if entity.CreateTime.Valid {
		tag.CreateTime = utils.LocalTime(entity.CreateTime.Time)
	}
}

func TranTagItemEntity(entity model.BuiltInTagItemEntity, tagItem *model.BuiltInTagItem) {
	tagItem.ID = entity.ID
	tagItem.TagId = entity.TagId
	tagItem.TagKey = entity.TagKey
	tagItem.TagName = entity.TagName
	tagItem.Status = entity.Status
	if entity.SortNo.Valid {
		tagItem.SortNo = entity.SortNo.Int64
	}
	if entity.CreateTime.Valid {
		tagItem.CreateTime = utils.LocalTime(entity.CreateTime.Time)
	}
}
