package conf

import (
	"hr-dapp/srv/pkg/contracts"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
)

// 相应设置配置
type ContractAddress struct {
	OwnerAddr string       `yaml:"ownerAddr"`
	ProxyAddr ProxyAddress `yaml:"proxyAddr"`
	Network   string       `yaml:"network"`
	Updated   string       `yaml:"updated"`
}

// 服务配置
type ProxyAddress struct {
	Proxy          string `yaml:"proxy"`
	Admin          string `yaml:"admin"`
	Implementation string `yaml:"implementation"`
}

var contractConfig = &ContractAddress{}

// 初始化方法
func InitContractConfig() {
	ex, _ := os.Executable()
	yamlFile, err := ioutil.ReadFile(filepath.Dir(ex) + "/resource/deployed_contracts_" + os.Getenv("ENV") + ".yml")
	if err != nil {
		panic("配置文件未找到")
	}

	err = yaml.Unmarshal(yamlFile, contractConfig)
	if err != nil {
		panic("获得配置文件出错")
	}
}

func GetEthClient() *ethclient.Client {
	client, err := ethclient.Dial(contractConfig.Network)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetFacadeContract() *contracts.Facade {
	instance, err := contracts.NewFacade(GetProxyAddr(), GetEthClient())
	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func GetProxyAddr() common.Address {
	return common.HexToAddress(contractConfig.ProxyAddr.Proxy)
}

func GetOwnerAddr() common.Address {
	return common.HexToAddress(contractConfig.ProxyAddr.Admin)
}
