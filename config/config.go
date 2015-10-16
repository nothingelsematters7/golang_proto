package config

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config contains info about for connecting to RabbitMQ
type Config struct {
	Host  string
	User  string
	Pass  string
	Vhost string
	Port  int
}

// AMQPUrl returns AMQP URL for connecting to Rabbit MQ
func (conf Config) AMQPUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", conf.User, conf.Pass, conf.Host,
		conf.Port, conf.Vhost)
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
