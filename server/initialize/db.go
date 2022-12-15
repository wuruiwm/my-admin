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

// Db 初始化Db
func Db() {
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

	conn, err := gorm.Open(mysql.Open(dbDsn()), gormConfig)
	if err != nil {
		panic("mysql conn error: " + err.Error())
	}
	global.Db = conn
	dbPool()
	dbPlugin()
}

func dbDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
		global.Config.Mysql.Username,
		global.Config.Mysql.Password,
		global.Config.Mysql.Host,
		global.Config.Mysql.Port,
		global.Config.Mysql.Database,
		global.Config.Mysql.Charset)
}

// dbPool 设置db连接池参数
func dbPool() {
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

func dbPlugin() {
	pluginList := []gorm.Plugin{
		&dbSqlLog{},
	}
	for _, v := range pluginList {
		err := global.Db.Use(v)
		if err != nil {
			panic(err)
		}
	}
}

type dbSqlLog struct {
}

func (l *dbSqlLog) Name() string {
	return "db_sql_log"
}

func (l *dbSqlLog) Initialize(db *gorm.DB) error {
	//执行sql之前
	_ = db.Callback().Create().Before("gorm:before_create").Register("before", l.Before)
	_ = db.Callback().Query().Before("gorm:query").Register("before", l.Before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register("before", l.Before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register("before", l.Before)
	_ = db.Callback().Row().Before("gorm:row").Register("before", l.Before)
	_ = db.Callback().Raw().Before("gorm:raw").Register("before", l.Before)

	//执行sql之后
	_ = db.Callback().Create().After("gorm:after_create").Register("after", l.After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after", l.After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after", l.After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after", l.After)
	_ = db.Callback().Row().After("gorm:row").Register("after", l.After)
	_ = db.Callback().Raw().After("gorm:raw").Register("after", l.After)
	return nil
}

func (l *dbSqlLog) Before(db *gorm.DB) {
	db.InstanceSet("start_time", time.Now())
	return
}

func (l *dbSqlLog) After(db *gorm.DB) {
	startTime, ok := db.InstanceGet("start_time")
	if !ok {
		return
	}
	t, ok := startTime.(time.Time)
	if !ok {
		return
	}
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	global.Logger.Info("sql",
		zap.String("sql", sql),
		zap.Int64("row", db.Statement.RowsAffected),
		zap.String("cost", util.TimeSince(t)),
	)
}
