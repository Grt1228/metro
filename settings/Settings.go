package settings

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `json:"name" mapstructure:"name"`
	Version      string `json:"version" mapstructure:"version"`
	Port         int    `json:"port" mapstructure:"port"`
	*MysqlConfig `json:"mysql" mapstructure:"mysql"`
	*RedisConfig `json:"redis" mapstructure:"redis"`
}

type MysqlConfig struct {
	Host     string `json:"port" mapstructure:"host"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Port     string `json:"port" mapstructure:"port"`
	Database string `json:"database" mapstructure:"database"`
}

type RedisConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

func Init() error {
	viper.SetConfigFile("./conf/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了~~~~~~~~~~~")
		viper.Unmarshal(&Conf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败了~~~~~~~~~~~~~, err: %v", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("解析配置文件失败了~~~~~~~~~~~~~~~, err:%v", err))
	}
	marshal, _ := json.Marshal(Conf)
	println(string(marshal))
	return err
}
