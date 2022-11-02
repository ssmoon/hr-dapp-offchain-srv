package services

import (
	"encoding/json"
	"hr-dapp/srv/pkg/conf"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func GrantUserPrivilege(userAddr common.Address) error {
	facade := conf.GetFacadeContract()
	tx, err := facade.CreateUser(conf.GetTransactionOpts(conf.GetOwnerKey(), conf.GetOwnerAddr()), userAddr)
	manifestJson, _ := json.MarshalIndent(tx, "", "    ")
	log.Println(string(manifestJson))
	return err
}
