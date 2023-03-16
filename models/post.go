package models

import (
	"time"

	"github.com/mumingv/gin-blog/dao"
)

type Post struct {
	Id         int
	UserId     int
	Title      string
	Url        string
	Content    string
	Tags       string
	Views      int
	IsTop      int
	Created    time.Time
	Updated    time.Time
	CategoryId int
	Status     int
	Types      int
	Info       string
	Image      string
}

// admin 商品列表
func GetArtileList(offset int, pagesize int) (articleList []*Post, err error) {

	db := dao.DB
	db = db.Offset(offset).Limit(pagesize).Order(" is_top desc , created desc")

	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func GetDetailById(id int) (post *Post, err error) {
	post = new(Post)

	if err = dao.DB.Debug().Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return
}

// 商品添加
func CreatePost(post *Post) (err error) {
	err = dao.DB.Create(&post).Error
	return
}

// 商品更新
func UpdatePost(post *Post) (err error) {
	err = dao.DB.Save(post).Error
	return
}

// 根据ID删除类目
func DeletePost(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Post{}).Error
	return
}

func GetAllArticle(keyword string, cate_id int, actionName string, page int, pagesize int, sort_type int) (articleList []*Post, err error) {

	if page < 1 {
		page = 1
	}
	db := dao.DB
	// 设置过滤字段
	if cate_id > 0 {
		db = db.Where("category_id = ?", cate_id)
	}
	// 搜索关键词（采用模糊匹配）
	if len(keyword) > 0 {
		db = db.Where("title like ?", "%"+keyword+"%")
	}
	if actionName == "resource" { // 资源列表 标签设置0
		db = db.Where("types = ?", 0)
	} else {
		db = db.Where("types = ?", 1)
	}
	if actionName == "home" { // 文章首页标签
		db = db.Where("is_top = ?", 1)
	}
	// https: //blog.csdn.net/linux_player_c/article/details/82351934
	// 设置分页
	offset := (page - 1) * pagesize
	db = db.Offset(offset).Limit(pagesize).Order("views desc")

	// 提供两种排序方式
	// 默认情况：文章创建时间排序
	// 热榜文章：按照浏览器排序
	if sort_type == 1 {
		db = db.Order("views desc")
	} else {
		db = db.Order("created desc")
	}
	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return
}

// 获取公告信息
func GetNotice(cate_id int) (articleList []*Post, err error) {
	db := dao.DB
	db = db.Where("category_id = ?", cate_id) // 公告类别字段
	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return
}
