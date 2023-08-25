package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"serverGo/foundation/mysql"
	"serverGo/foundation/resp"
	"serverGo/foundation/utils"
	"serverGo/model"
	"strconv"
)

func InsertOneArticle(article *model.Article) (err error) {

	sqlStr := "INSERT INTO new_article (`title`, `topic`, `content`, `tags`, `international`, `region`, `event_type`, `status`, `create_time`) VALUES ( :Title, :Topic, :Content, :Tags, :International, :Region, :EventType, :Status, :CreateTime);"
	result, err := mysql.DB.NamedExec(sqlStr,
		map[string]interface{}{
			"Title":         article.Title,
			"Topic":         article.Topic,
			"Content":       article.Content,
			"Tags":          article.Tags,
			"International": article.International,
			"Region":        article.Region,
			"EventType":     article.EventType,
			"Status":        article.Status,
			"CreateTime":    article.CreateTime.AsStr(),
		})
	if err != nil {
		fmt.Println(errors.WithStack(err))
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(errors.WithStack(err))
		return
	}
	article.ID = id
	return
}

func SelectArticlePage(formParam *model.ArticleSearchForm) (list resp.Paged[[]model.Article], err error) {
	sqlStr := "SELECT a.id,a.title,a.topic,a.tags,a.content,a.international,a.region,a.event_type, a.status,a.create_time,b1.tag_name internationalName,b2.tag_name regionName,b3.tag_name event_typeName FROM new_article a LEFT JOIN new_built_in_tag_item b1 ON a.international=b1.tag_key LEFT JOIN new_built_in_tag_item b2 ON a.region=b2.tag_key LEFT JOIN new_built_in_tag_item b3 ON a.event_type=b3.tag_key WHERE 1=1 "
	if formParam.SearchValue != "" {
		like := "%" + formParam.SearchValue + "%"
		appendStr := fmt.Sprintf(" AND(a.content LIKE '%s' OR a.title LIKE '%s' OR a.topic LIKE '%s')", like, like, like)
		sqlStr = sqlStr + appendStr
	}
	if formParam.International != 0 {
		sqlStr = sqlStr + " AND a.international=" + strconv.Itoa(formParam.International)
	}
	if formParam.Region != 0 {
		sqlStr = sqlStr + " AND a.region=" + strconv.Itoa(formParam.Region)
	}
	if formParam.EventType != 0 {
		sqlStr = sqlStr + " AND a.event_type=" + strconv.Itoa(formParam.EventType)
	}
	offset := (formParam.CurrentPage - 1) * formParam.PageSize
	countSql := "SELECT COUNT(1) FROM ( " + sqlStr + ") tmp"
	rows, err := mysql.DB.Query(countSql)
	if err != nil {
		fmt.Println(errors.WithStack(err))
		return
	}
	var totalCount int
	for rows.Next() {
		err = rows.Scan(&totalCount)
		if err != nil {
			fmt.Println(errors.WithStack(err))
			return
		}
	}
	pageCount := (totalCount / formParam.PageSize) + 1
	sqlStr = sqlStr + " LIMIT ? OFFSET ?"
	var entityList []model.ArticleEntity
	err = mysql.DB.Select(&entityList, sqlStr, formParam.PageSize, offset)
	if err != nil {
		fmt.Println(errors.WithStack(err))
		if err == sql.ErrNoRows {
			return list, errors.New("该Tag不存在")
		}
		return list, err
	}
	rowDataList := make([]model.Article, 0)
	if len(entityList) > 0 {
		for _, value := range entityList {
			article := model.Article{}
			TranArticleEntity(value, &article)
			rowDataList = append(rowDataList, article)
		}
	}
	list = resp.Paged[[]model.Article]{
		CurrentPage: formParam.CurrentPage,
		PageSize:    formParam.PageSize,
		TotalCount:  totalCount,
		PageCount:   pageCount,
		RowDataList: rowDataList,
	}
	return
}

func TranArticleEntity(entity model.ArticleEntity, article *model.Article) {
	article.ID = entity.ID
	article.Title = entity.Title
	article.Topic = entity.Topic
	article.Content = entity.Content
	article.Tags = entity.Tags
	article.International = entity.International
	article.Region = entity.Region
	article.EventType = entity.EventType
	article.Status = entity.Status
	if entity.CreateTime.Valid {
		article.CreateTime = utils.LocalTime(entity.CreateTime.Time)
	}
	article.InternationalName = entity.InternationalName
	article.RegionName = entity.RegionName
	article.EventTypeName = entity.EventTypeName
}
