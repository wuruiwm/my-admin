package global

import (
	"app/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
)
