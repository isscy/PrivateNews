package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"serverGo/foundation/resp"
	"serverGo/foundation/utils"
	"serverGo/logic"
	"serverGo/model"
)

// 发表一篇文章
func PublishOneArticleController(c *gin.Context) {
	// 获取参数 及 参数验证
	var articleForm *model.ArticleForm
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		zap.L().Error("发表文章的请求参数错误！", zap.Error(err))
		errMsg := utils.GetValigMsg(err, articleForm)
		fmt.Println("errMsg:  ", errMsg)
		resp.Fails(c, resp.InvalidParams, errMsg)
		return
	}

	article, err := logic.AddOneArticle(articleForm)
	if err != nil {
		zap.L().Error("查询内置Tag的请求参数错误：", zap.Error(err))
		resp.Fails(c, resp.BizCheckFail, err.Error())
		return
	}
	resp.Succ(c, article)
}

// 发表一篇文章
func GetArticlePageController(c *gin.Context) {
	// 获取参数 及 参数验证
	var articleForm *model.ArticleSearchForm
	_ = c.ShouldBindJSON(&articleForm)
	if articleForm.CurrentPage == 0 {
		articleForm.CurrentPage = 1
	}
	if articleForm.PageSize == 0 {
		articleForm.PageSize = 10
	}
	page, err := logic.GetArticlePage(articleForm)
	if err != nil {
		zap.L().Error("查询内置Tag的请求参数错误：", zap.Error(err))
		resp.Fails(c, resp.BizCheckFail, err.Error())
		return
	}
	resp.Succ(c, page)
}
