package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sps-template-server/config"
	"time"
)

var (
	SdDB    *gorm.DB
	SdConfig config.Server
	SdVp   *viper.Viper
	SdLog	*zap.Logger
)

type SdModel struct {
	ID			uint			`gorm:"primarykey"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"-"`
} 