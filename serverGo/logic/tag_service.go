package logic

import (
	"serverGo/dao"
	"serverGo/model"
)

// 获取内置TagList
func GetBuiltInTagList(tagParam *model.BuiltInTagForm) (tagList []model.BuiltInTag, err error) {
	list, err := dao.GetBuiltInTagList(tagParam)
	if err != nil {
		return nil, err
	}
	return list, nil
}
