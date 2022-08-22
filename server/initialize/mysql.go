package initialize

import (
	"app/global"
	"app/util"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//Mysql 初始化mysql连接池
func Mysql() {
	var gormConfig *gorm.Config
	var logLevel logger.LogLevel

	if global.Config.Debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	gormConfig = &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		SkipDefaultTransaction: true, //禁用默认写操作都使用事务 提高性能
	}

	conn, err := gorm.Open(mysql.Open(getMysqlConnString()), gormConfig)
	if err != nil {
		panic("mysql conn error: " + err.Error())
	}
	global.Db = conn
	setConnPool()
	setPlugin()
}

//getMysqlConnString 获取mysql连接字符串
func getMysqlConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
		global.Config.Mysql.Username,
		global.Config.Mysql.Password,
		global.Config.Mysql.Host,
		global.Config.Mysql.Port,
		global.Config.Mysql.Database,
		global.Config.Mysql.Charset)
}

//setMysqlSetting 设置mysql连接池参数
func setConnPool() {
	sqlDB, err := global.Db.DB()
	if err != nil {
		panic("mysql set coon pool error: " + err.Error())
	}
	//用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConn)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConn)
	//设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func setPlugin() {
	pluginList := []gorm.Plugin{
		&GormLog{},
	}
	for _, v := range pluginList {
		err := global.Db.Use(v)
		if err != nil {
			panic(err)
		}
	}
}

type GormLog struct {
}

func (g *GormLog) Name() string {
	return "gorm_log"
}

func (g *GormLog) Initialize(db *gorm.DB) error {
	//执行sql之前
	_ = db.Callback().Create().Before("gorm:before_create").Register("before", g.Before)
	_ = db.Callback().Query().Before("gorm:query").Register("before", g.Before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register("before", g.Before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register("before", g.Before)
	_ = db.Callback().Row().Before("gorm:row").Register("before", g.Before)
	_ = db.Callback().Raw().Before("gorm:raw").Register("before", g.Before)

	//执行sql之后
	_ = db.Callback().Create().After("gorm:after_create").Register("after", g.After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after", g.After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after", g.After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after", g.After)
	_ = db.Callback().Row().After("gorm:row").Register("after", g.After)
	_ = db.Callback().Raw().After("gorm:raw").Register("after", g.After)
	return nil
}

func (g *GormLog) Before(db *gorm.DB) {
	db.Set("start_time", time.Now())
	return
}

func (g *GormLog) After(db *gorm.DB) {
	startTime, ok := db.Get("start_time")
	if !ok {
		return
	}
	t, ok := startTime.(time.Time)
	if !ok {
		return
	}
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	global.Logger.Info("sql",
		zap.String("time", util.GetDate()),
		zap.String("sql", sql),
		zap.Int64("row", db.Statement.RowsAffected),
		zap.Duration("cost", time.Since(t)),
	)
}
