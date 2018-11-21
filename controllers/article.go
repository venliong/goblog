package controllers

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//文章详情
func ArticleDetail(c *gin.Context) {
	id := c.Param("id")

	articleDetail, err := model.GetArticleDetial(id)
	if err != nil {
		fmt.Println(err)
	}

	//推荐文章列表
	recommendList, err := model.GetArticleRecommend(articleDetail.CategoryId)
	if err != nil {
		fmt.Println(err)
	}

	//获取点击排行
	hitsList, err := model.GetArticleHits(10, articleDetail.CategoryId)
	if err != nil {
		fmt.Println(err)
	}

	//获取相关文章
	aboutList, err := model.GetArticlListAbout(6, articleDetail)
	if err != nil {
		fmt.Println(err)
		aboutList = nil
	}

	/**
	获取上一篇文章
	*/
	prevArticle, err := model.GetArticlePrev(id, "prev")
	if err != nil {
		fmt.Println(err)
		prevArticle = nil
	}

	/**
	获取下一篇文章
	*/
	nextArticle, err := model.GetArticlePrev(id, "next")
	if err != nil {
		fmt.Println(err)
		nextArticle = nil
	}

	//更新文章点击量
	model.CounterArticle(id, 1)

	//文章标签
	var tags []string
	if articleDetail.Tags != "" {
		tags = strings.Split(articleDetail.Tags, ",")
	}

	c.HTML(http.StatusOK, "article/detail.html", gin.H{
		"aboutList":     aboutList,
		"articleDetail": articleDetail,
		"recommendList": recommendList,
		"hitsList":      hitsList,
		"prevArticle":   prevArticle,
		"nextArticle":   nextArticle,
		"tags":          tags,
	})
}

/**
修改文章的计数
*/
func ArticleCounter(c *gin.Context) {
	id := c.PostForm("id")
	source, err := strconv.Atoi(c.PostForm("source"))
	if err != nil {
		fmt.Println(err)
	}

	res, err := model.CounterArticle(id, uint(source))

	if res {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "操作成功!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "操作失败!",
		})
	}
}
