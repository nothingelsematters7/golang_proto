package config

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config contains global settings
type Config struct {
	AMQP_HOST  string
	AMQP_USER  string
	AMQP_PASS  string
	AMQP_VHOST string
	AMQP_PORT  int

	REDIS_HOST string
	REDIS_PASS string
	REDIS_PORT int
	REDIS_DB   int64
}

// AMQPUrl returns AMQP URL for connecting to Rabbit MQ
func (conf Config) AMQPUrl() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d%s", conf.AMQP_USER, conf.AMQP_PASS, conf.AMQP_HOST,
		conf.AMQP_PORT, conf.AMQP_VHOST)
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
