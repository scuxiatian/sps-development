package datas

import (
	"gorm.io/gorm"
	"os"
	"sps-template-server/global"
	"sps-template-server/model"
	"time"
)

var Authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
}

func InitSysAuthority(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			global.SdLog.Warn("sys_authorities表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		global.SdLog.Error("[Mysql]--> sys_authorities 表的初始数据失败,err: %v\n")
		os.Exit(0)
	}
}
