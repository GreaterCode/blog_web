package models

import (
	"blog_web/config"
	"blog_web/utils"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	content    string
	Author     string
	CreateTime string

	// 查看地址
	Link string

	//修改地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type HomeFoolterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

type TagLink struct {
	TagName string
	TagUrl  string
}

// 首页显示内容

func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		fmt.Println(art.Title)
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLink(art.Tags)
		fmt.Println("tag---->" + art.Tags)
		homeParam.Short = art.Short
		homeParam.content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/show/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		// 解析模板文件，进行变量插入
		t, _ := template.ParseFiles("views/home_block.html")

		// 开辟缓存
		buffer := bytes.Buffer{}

		// html数据替换
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

// 将tags 字符串转换成首页模板所需要操作的数据结构
func createTagsLink(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&") //一篇文章多个tag 用&分隔开
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?atg=" + tag})
	}
	return tagLink
}

// 翻页
func ConfigHomeFooterPageCode(page int) HomeFoolterPageCode {
	pageCode := HomeFoolterPageCode{}
	num := GetArticleRowsNum()
	//计算总页数
	allPageNum := (num-1)/config.NUM + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	// 如果当前页数大于总页数，下一页不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	// 如果当前页数大于总页数，下一页不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/？page=" + strconv.Itoa(page+1)

	return pageCode
}
