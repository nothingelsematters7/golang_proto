package config

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config contains global settings
type Config struct {
	AMQP_HOST      string
	AMQP_USER      string
	AMQP_PASS      string
	AMQP_VHOST     string
	AMQP_PORT      int

	REDIS_HOST     string
	REDIS_PASS     string
	REDIS_PORT     int
	REDIS_DB       int64

	MYSQL_USER     string
	MYSQL_PASS     string
	MYSQL_DATABASE string

	OANDA_KEY      string
}

// AMQPUrl returns AMQP URL for connecting to Rabbit MQ
func (conf Config) AMQPUrl() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d%s", conf.AMQP_USER, conf.AMQP_PASS, conf.AMQP_HOST,
		conf.AMQP_PORT, conf.AMQP_VHOST)
}

func (conf Config) MysqlArgs() string {
	return fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MYSQL_USER, conf.MYSQL_PASS, conf.MYSQL_DATABASE)
}

// Conf variable is global config instance
var Conf Config

func init() {
	configPath, _ := filepath.Abs("config/config.toml")
	_, err := toml.DecodeFile(configPath, &Conf)
	if err != nil {
		panic("Error while decoding config.toml")
	}
}
