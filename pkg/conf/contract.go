package conf

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"hr-dapp/srv/pkg/contracts"
	"log"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// 相应设置配置
type ContractAddress struct {
	OwnerAddr  string       `json:"ownerAddr"`
	OwnerKey   string       `json:"ownerKey"`
	ProxyAddr  ProxyAddress `json:"proxyAddr"`
	Network    string       `json:"network"`
	Updated    string       `json:"updated"`
	MemberAddr string       `json:"memberAddr"`
	MemberKey  string       `json:"memberKey"`
	ChainId    int64        `json:"chainId"`
}

// 服务配置
type ProxyAddress struct {
	Proxy          string `json:"proxy"`
	Admin          string `json:"admin"`
	Implementation string `json:"implementation"`
}

var contractConfig = &ContractAddress{}

// 初始化方法
func InitContractConfig() {
	ex, _ := os.Executable()
	jsonFile, err := os.ReadFile(filepath.Dir(ex) + "/resource/deployed_contracts_dev.json")
	if err != nil {
		panic("配置文件未找到")
	}

	err = json.Unmarshal([]byte(jsonFile), contractConfig)
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

func GetOwnerKey() string {
	return contractConfig.OwnerKey
}

func GetOwnerAddr() common.Address {
	return common.HexToAddress(contractConfig.OwnerAddr)
}

func GetCurrentUserAddr() common.Address {
	return common.HexToAddress(contractConfig.MemberAddr)
}

func GetCurrentUserKey() string {
	return contractConfig.MemberKey
}

func GetTransactionOpts(key string, from common.Address) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := GetEthClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(contractConfig.ChainId))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	auth.From = from

	return auth
}

func GetCallOpts(from common.Address) *bind.CallOpts {
	return &bind.CallOpts{From: from}
}

func ConvertStringToByte32(str string) [32]byte {
	var result [32]byte
	copy(result[:], str)
	return result
}

func Sha3StringToByte32(str string) [32]byte {
	var result [32]byte
	copy(result[:], solsha3.SoliditySHA3(solsha3.String(str)))
	return result
}

func ComputeWorkerHash(birthAt uint16, graduatedAt uint16, collegeCode string) string {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint16", "uint16", "string"},

		// values
		[]interface{}{
			birthAt,
			graduatedAt,
			collegeCode,
		},
	)
	return hex.EncodeToString(hash)
}
