package dao

import (
	"QrCodeLogin/pkg/log"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	orm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
)

func CreateDSN(username, password, host, dbName string, timeout time.Duration) (string, error) {
	// [username:password@protocol(ip:port)/dbName]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbName)

	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", errors.New(fmt.Sprintf("%s, config error for host: %s", err.Error(), host))
	}

	if timeout > 0 {
		config.Timeout = timeout      // Dial timeout 连接超时时间
		config.ReadTimeout = timeout  // I/O read timeout 读超时时间
		config.WriteTimeout = timeout // I/O write timeout 写超时时间
	}

	return config.FormatDSN(), nil
}

func InitMySQLInstance(confPath string) (err error) {
	// 获取MySQL配置信息
	conf := struct {
		Username string        `yaml:"username"`
		Password string        `yaml:"password"`
		Host     string        `yaml:"host"`
		DbName   string        `yaml:"dbName"`
		Timeout  time.Duration `yaml:"timeout"`
	}{}

	body, err := ioutil.ReadFile(confPath + "/mysql.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(body, &conf)
	if err != nil {
		return
	}

	// 初始化MySQL对象
	dsn, err := CreateDSN(conf.Username, conf.Password, conf.Host, conf.DbName, conf.Timeout)
	if err != nil {
		return
	}
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	if err = conn.Ping(); err != nil {
		return
	}

	conn.SetMaxOpenConns(100)              // 设置最大的连接数
	conn.SetMaxIdleConns(50)               // 设置连接池中最大的空闲连接数量
	conn.SetConnMaxLifetime(7 * time.Hour) // 设置单个连接的最大生命周期

	db, err = gorm.Open(orm.New(orm.Config{
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Info("MySQL数据库启动成功")
	return
}

func GetMySQLInstance() *gorm.DB {
	return db
}
