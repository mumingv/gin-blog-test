package models

import "github.com/mumingv/gin-blog/dao"

type Config struct {
	Id    int
	Name  string
	Value string
}

// ConfigList
func ConfigList() (config []*Config, err error) {
	if err = dao.DB.Find(&config).Error; err != nil {
		return nil, err
	}
	return config, nil
}

// UpdateConfig
func UpdateConfig(config *Config) (err error) {
	err = dao.DB.Save(config).Error
	return err
}
