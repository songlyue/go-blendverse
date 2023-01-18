package initialize

import (
	"blendverse/global"
	"blendverse/model/example"
	"blendverse/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitGorm() *gorm.DB {
	p := global.GVA_CONFIG.Pgsql
	if p.DbName == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig)); err != nil {
		return nil
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(10)
		sqlDb.SetMaxIdleConns(10)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.User{},
		example.Example{},
	)
	if err != nil {
		global.GVA_LOG.Error("初始化数据库失败", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
