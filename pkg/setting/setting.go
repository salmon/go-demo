package setting

import (
	"time"

	"io/ioutil"
	"github.com/sirupsen/logrus"
	"github.com/ghodss/yaml"
)

var Config AppMarketConfig

type HttpConfig struct {
	HTTPPort     int           `json:"port,default=9999"`
	TLS          bool          `json:"tls"`
	TlsKey       string        `json:"tlsKey"`
	TlsCert      string        `json:"tlsCert"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
}

type MongoConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	Timeout  int    `json:"timeout"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type TaskConfig struct {
	Broker          string       `json:"broker"`
	DefaultQueue    string       `json:"default_queue"`
	ResultBackend   string       `json:"result_backend"`
	ResultsExpireIn int          `json:"results_expire_in"`
}

type AppMarketConfig struct {
	Debug bool `json:"debug"`

	TaskConfig *TaskConfig `json:"taskConfig"`
	RedisConfig *RedisConfig `json:"redisConfig"`
	MongoConfig *MongoConfig `json:"mongoConfig"`
	HttpConfig *HttpConfig `json:"serverConfig"`
}

// StartResyncReleaseCaches sets values from the environment.
func InitConfig(configPath string) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		logrus.Fatalf("Read config file faild! %s\n", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		logrus.Fatalf("Unmarshal config file faild! %s\n", err.Error())
	}
}
