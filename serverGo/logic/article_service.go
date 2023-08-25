package logic

import (
	"serverGo/dao"
	"serverGo/foundation/resp"
	"serverGo/foundation/utils"
	"serverGo/model"
	"strings"
	"time"
)

// 新增一篇文章
func AddOneArticle(form *model.ArticleForm) (article *model.Article, err error) {
	article = &model.Article{
		Title:         form.Title,
		Topic:         form.Topic,
		Content:       form.Content,
		Tags:          "",
		International: form.International,
		Region:        form.Region,
		EventType:     form.EventType,
		Status:        "1",
		CreateTime:    utils.LocalTime(time.Now()),
	}
	if form.TagList != nil && len(form.TagList) > 0 {
		tags := strings.Join(form.TagList, ",")
		article.Tags = tags
	}
	err = dao.InsertOneArticle(article)
	return article, err
	/*	if err := dao.InsertOneArticle(article); err != nil {
			return
		}
		return*/
}

// 新增一篇文章
func GetArticlePage(form *model.ArticleSearchForm) (list resp.Paged[[]model.Article], err error) {
	return dao.SelectArticlePage(form)
}
