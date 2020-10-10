package Services

import (
	"GoRestBasics/Models"
	"github.com/ahmetb/go-linq"
)

var articles = []Models.Article{
	{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func GetArticles() []Models.Article {
	return articles
}

func GetArticle(id int) Models.Article {
	var filteredArticles []Models.Article
	var result Models.Article

	linq.From(articles).WhereT(func(a Models.Article) bool {
		return a.Id == id
	}).Select(func(c interface{}) interface{} {
		return c.(Models.Article)
	}).ToSlice(&filteredArticles)

	if filteredArticles != nil {
		result = filteredArticles[0]
	}

	return result
}
