package models

import (
	"blog_web/config"
	"blog_web/database"
	"fmt"
	"log"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

// 添加文章
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	if err != nil {
		log.Println("插入文章失败")
		return i, err
	}
	SetArticleRowsNum()
	return i, err

}

// 插入文章
func insertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into article(title, tags, short, content, author, createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

// 按页查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	fmt.Println("---------->page", page)

	return QueryArticleWithPage(page, config.NUM)

}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticlesWithCon(sql)
}

// 根据条件查询文章
func QueryArticlesWithCon(sql1 string) ([]Article, error) {
	sql := "select id, title, tags, short, content, author, createtime from article" + sql1
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

// 翻页
var articleRowsNum = 0

func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

// 设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}

// 查询文章总条数
func QueryArticleRowNum() int {
	row := database.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)

	return num
}

//  查询文章
func QueryArticleWithId(id int) Article {
	row := database.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

//  修改数据
func UpdateArticle(article Article) (int64, error) {
	return database.ModifyDB("update article set title=?, tags=?, short=?, content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//删除文章
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArticle(artID)
	SetArticleRowsNum()
	return i, err

}

func deleteArticleWithArticle(artID int) (int64, error) {
	return database.ModifyDB("delete from article where id=?", artID)
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := database.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号
通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]Article, error) {

	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticlesWithCon(sql)
}
