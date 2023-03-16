package models

import (
	"time"

	"github.com/mumingv/gin-blog/dao"
)

type Category struct {
	Id      int
	Name    string
	Created time.Time
	Updated time.Time
}

// 获取所有类别
func CategoryList() (dataList []*Category, err error) {
	if err = dao.DB.Find(&dataList).Error; err != nil {
		return nil, err
	}
	return dataList, nil
}

func GetCategoryById(id int) (categoryList []*Category, err error) {
	db := dao.DB
	db = db.Where("id = ?", id)
	if err = db.Find(&categoryList).Error; err != nil {
		return nil, err
	}
	return categoryList, nil
}

// 更新类目
func UpdateCategory(category *Category) (err error) {
	err = dao.DB.Save(category).Error
	if err != nil {
		return err
	}
	return nil
}

// 根据ID删除类目
func DeleteCategory(id string) (err error) {
	if err = dao.DB.Where("id=?", id).Delete(&Category{}).Error; err != nil {
		return err
	}
	return nil
}
