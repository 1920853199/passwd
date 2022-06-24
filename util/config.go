package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

var configName = "config.json"

type Config struct {
	Addr  string `json:"addr"`
	Api   string `json:"api"`
	Token string `json:"token"`
}

func GetConfigPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".passwd")
}

func GetConfigFilePath() string {
	return path.Join(GetConfigPath(), configName)
}

// /usr/local/
func GetConfigByClient() (*Config, error) {
	data, err := ioutil.ReadFile(GetConfigFilePath())
	if err != nil {
		return nil, fmt.Errorf("配置读取失败：%s", err.Error())
	}
	conf := &Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, fmt.Errorf("配置解析数据失败：%s", err.Error())
	}
	return conf, nil
}
