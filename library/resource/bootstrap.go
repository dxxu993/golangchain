package resource

import (
	"context"
	"time"

	"github.com/dxxu993/golangchain/library/conf"
	"github.com/dxxu993/golangchain/library/logger"

	"github.com/patrickmn/go-cache"
	gocache "github.com/patrickmn/go-cache"
	logrus "github.com/sirupsen/logrus" // gorm mysql 驱动包
	"gorm.io/gorm"                      // gorm
)

var Logger *logrus.Logger
var AppConf *conf.AppConf
var DbConn *gorm.DB
var LocalCache *gocache.Cache

func Bootstrap(ctx context.Context) {
	Logger = logger.InitLogger()
	AppConf = conf.InitAppConf()
}

func InitLocalCache() {
	// 设置默认超时时间和清理时间
	LocalCache = cache.New(5*time.Minute, 10*time.Minute)
}
