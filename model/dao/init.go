package dao

import (
	"fmt"
	"github.com/TakanashiRikka-Na/Rlog"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Client *redis.Client
var err error

func Init() {
	DB, err = InitDataBase()
	if err != nil {
		Rlog.Fatal("数据库链接失败", err)
	}
	Client, err = InitRedis()
	if err != nil {
		Rlog.Fatal("Redis链接失败", err)
	}
}
func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + viper.GetString("redis.port"),
		Password: viper.GetString("Redis_Password"),
		DB:       viper.GetInt("db"),
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func InitDataBase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql_passwd"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return DB, nil
}
