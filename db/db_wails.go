package db

import (
	"dayz-server-tools/logger"
	"sync"

	"gorm.io/gorm"
)

var (
	orm     *Orm
	ormOnce sync.Once
)

type Orm struct {
	Db *gorm.DB
}

func GetOrm() *Orm {
	ormOnce.Do(func() {
		orm = &Orm{
			Db: GetDB(),
		}
	})
	return orm
}

func (o *Orm) DBBridgeTest() error {
	logger.Info("正在测试数据库连接")
	return nil
}
