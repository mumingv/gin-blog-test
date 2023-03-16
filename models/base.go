package models

import "github.com/mumingv/gin-blog/dao"

func Count(value interface{}) (int, error) {

	var total int = 0
	dao.DB.Model(value).Count(&total)
	return total, nil
}
