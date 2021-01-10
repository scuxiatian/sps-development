package utils

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"sps-template-server/global"
	"time"
)

//@function: GetWriteSyncer
//@description: zap logger中加入file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer() (zapcore.WriteSyncer, error)  {
	fileWriter, err := rotatelogs.New(
		path.Join(global.SdConfig.Zap.Director, "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour))
	if global.SdConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}