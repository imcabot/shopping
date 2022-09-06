package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var cfgReader *configReader

type (
	ConfigUration struct {
		DatabaseSettings
		JwtSettings
	}
	//数据库配置
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	//Jwt配置
	JwtSettings struct {
		SecretKey string
	}
	//reader
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

func GetAllConfigValues(configFile string) (configuration *ConfigUration, err error) {
	newConfigReader(configFile)
	if err := cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取失败：%s", err)
		return nil, err
	}
	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败：%s", err)
		return nil, err
	}
	return configuration, nil
}

//实例化configReader
func newConfigReader(configFile string) {
	v := viper.GetViper()

	//v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}
