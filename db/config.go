package db

import (
	"errors"

	"gorm.io/gorm"
)

type Config struct {
	BaseModel
}

func (c *Config) GetConfig() (*Config, error) {
	if err := GetDB().First(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// 主题 'light' or 'dark'
func (c *Config) SetConfig(theme string) error {
	if theme != "light" && theme != "dark" {
		return errors.New("theme must be 'light' or 'dark'")
	}
	var config Config
	if err := GetDB().Where("id = ?", 1).First(&config).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			GetDB().Create(&Config{})
			return nil
		}
		return err
	} else {
		if err := GetDB().Save(&config).Error; err != nil {
			return err
		}
	}
	return nil
}
