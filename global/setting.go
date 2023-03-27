package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	// Logger 日志
	Logger *logger.Logger
)
