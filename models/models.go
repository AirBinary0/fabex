/*
   Copyright 2019 Vadim Inshakov

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package models

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/vadiminshakov/fabex/db"
)

type Fabex struct {
	Db            db.DbManager
	ChannelClient *channel.Client
	LedgerClient  *ledger.Client
}

type DB struct {
	Host       string
	Port       int
	Dbuser     string
	Dbsecret   string
	Dbname     string
	Collection string
}

type GRPCServer struct {
	Host string
	Port string
}

type Fabric struct {
	User              string
	Secret            string
	Org               string
	Channel           string
	ConnectionProfile string
}

type Config struct {
	DB
	Fabric
	GRPCServer
}

type Chaincode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Block struct {
	BlockHash string `json:"blockhash"`
	Blocknum  uint64 `json:"blocknum"`
	Tx        []Tx   `json:"txs"`
}

type Tx struct {
	Txid string `json:"txid"`
	KW   []KW
}

type KW struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
