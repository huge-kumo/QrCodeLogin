package core

import (
	"QrCodeLogin/internal/dao"
	"QrCodeLogin/internal/http"
	"QrCodeLogin/pkg/log"
)

func InitApp(confPath string) (closeFunc func(), err error) {
	// 初始化MySQL数据库
	if err = dao.InitMySQLInstance(confPath); err != nil {
		return
	}

	// 数据库Redis数据库
	if err = dao.InitRedisInstance(confPath); err != nil {
		return
	}

	// 初始化HTTP服务
	if err = http.InitHttpService(confPath); err != nil {
		return
	}

	return
}

func closeFunc() {
	defer func() {
		_ = log.Close()
	}()

}
