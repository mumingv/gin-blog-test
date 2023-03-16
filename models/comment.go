package models

import (
	"time"

	"github.com/mumingv/gin-blog/dao"
)

type Comment struct {
	Id       int
	Username string
	Content  string
	Created  time.Time
	PostId   int
	Ip       string
}

// 获取当前文章的评价
func GetCommentById(post_id int) (dataList []*Comment, err error) {

	// 设置当前文章的ID
	db := dao.DB.Debug().Where("post_id = ?", post_id)
	// 评论内容按照创建时间逆序
	db = db.Order("created desc")
	if err = db.Find(&dataList).Error; err != nil {
		return nil, err
	}
	return
}

// 插入评价
func CreateComment(comment *Comment) (err error) {
	if err = dao.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
