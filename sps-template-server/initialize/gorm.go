package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/workflow"
)

//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB)  {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		err = db.Migrator().CreateTable(&model.CasbinRule{})
	}
	err = db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
		model.SysApi{},
		model.JwtBlacklist{},
		model.FileUploadAndDownload{},
		model.WorkflowProcess{},
		model.WorkflowNode{},
		model.WorkflowEdge{},
		model.WorkflowStartPoint{},
		model.WorkflowEndPoint{},
		model.WorkflowMove{},
		workflow.ExaWfLeave{})
	if err != nil {
		global.SdLog.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.SdLog.Info("register table success")

	// 初始化数据，仅第一次启动程序时用，初始化完数据后需注释掉
	initMysqlTables(db)
	global.SdLog.Info("initialize table success")
}

//@function: initMysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func initMysqlTables(db *gorm.DB)  {
	//datas.InitSysUser(db)
	//datas.InitSysAuthority(db)
	//datas.InitSysBaseMenus(db)
	//datas.InitAuthorityMenu(db)
	//datas.InitSysAuthorityMenus(db)
	//datas.InitSysApi(db)
	//datas.InitCasbinModel(db)
}

//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB  {
	m := global.SdConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.SdLog.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig(mod bool) *gorm.Config  {
	switch global.SdConfig.Mysql.LogZap {
	case "Silent":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Error":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Error),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Warn":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Warn),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	case "Info":
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	default:
		if mod {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Info),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		} else {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Silent),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}
	}
}
