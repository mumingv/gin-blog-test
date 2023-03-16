package models

import "github.com/mumingv/gin-blog/dao"

func Count(value interface{}) (int, error) {

	var total int64 = 0
	dao.DB.Model(value).Count(&total)
	return int(total), nil
}
