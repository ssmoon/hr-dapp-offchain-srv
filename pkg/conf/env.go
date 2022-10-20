package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// 相应设置配置
type Setting struct {
	RunMode    string      `yaml:"runMode"`
	Server     server      `yaml:"server"`
	Database   database    `yaml:"database"`
	Redis      redisConf   `yaml:"redis"`
	Log        logSetting  `yaml:"log"`
	WxAccounts []wxAccount `yaml:"wx"`
}

// 服务配置
type server struct {
	HTTPPort     int `yaml:"HTTPPort"`
	ReadTimeout  int `yaml:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout"`
}

// 数据库配置
type database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbName"`
	Port     int    `yaml:"port"`
}

type redisConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth string `yaml:"auth"`
}

type logSetting struct {
	Path   string `yaml:"path"`
	Level  string `yaml:"level"`
	Prefix string `yaml:"prefix"`
}

type wxAccount struct {
	AppID  string `yaml:"appId"`
	Secret string `yaml:"secret"`
	Token  string `yaml:"token"`
	AesKey string `yaml:"aesKey"`
	name   string `yaml:"name"`
}

var conf = &Setting{}

// 初始化方法
func InitEnv() {
	ex, _ := os.Executable()
	yamlFile, err := ioutil.ReadFile(filepath.Dir(ex) + "/resource/application_" + os.Getenv("ENV") + ".yml")
	if err != nil {
		panic("配置文件未找到")
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic("获得配置文件出错")
	}
}

// 获取配置  外部调用使用
func Conf() *Setting {
	return conf
}
