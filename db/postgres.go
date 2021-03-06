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

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Instance *sql.DB
}

func CreateDBConfPostgres(host string, port int, user, password, dbname string) *DB {
	return &DB{host, port, user, password, dbname, &sql.DB{}}
}

func (db *DB) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBname)

	dbinstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = dbinstance.Ping()
	if err != nil {
		return err
	}

	dbinstance.SetMaxIdleConns(15)
	db.Instance = dbinstance

	log.Println("Connected to Postgres successfully")
	return nil
}

func (db *DB) Init() error {

	// create simple table with transaction id, block hash, block number
	// keys a and b (for simple chaincode)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password)

	dbinstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = dbinstance.Ping()
	if err != nil {
		return err
	}

	_, err = dbinstance.Exec(`
        CREATE DATABASE txs
    	`)
	if err != nil {
		return err
	}

	// close old connection and reconnect to new database
	dbinstance.Close()
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, "txs")

	dbinstance, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	_, err = dbinstance.Exec(`
        CREATE TABLE txs (
    		txid TEXT PRIMARY KEY,
    		blockhash TEXT,
    	    blocknum INT,
    	    payload BYTEA
    	)`)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Insert(txid, blockhash string, blocknum uint64, payload string) error {
	query := `INSERT INTO public.txs (txid, blockhash, blocknum, payload) VALUES ($1, $2, $3, $4);`
	_, err := db.Instance.Exec(query, txid, blockhash, blocknum, payload)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) QueryBlockByHash(hash string) ([]Tx, error) {
	//query := fmt.Sprintf(`
	//    SELECT (txid, blockhash, blocknum, payload) FROM public.txs
	//    WHERE blockhash='%s';`, hash)
	//
	//var (
	//	txid, blockhash string
	//	blocknum        uint64
	//	payload         string
	//)
	//
	//err := db.Instance.QueryRow(query).Scan(&txid, &blockhash, &blocknum, &payload)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return Tx{txid, blockhash, blocknum, payload}, nil
	return nil, nil
}

func (db *DB) QueryAll() ([]Tx, error) {
	arr := []Tx{}

	query := fmt.Sprintf(`
        SELECT * FROM public.txs;`)
	rows, err := db.Instance.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		var (
			txid, blockhash string
			blocknum        uint64
			payload         string
		)

		err := rows.Scan(&txid, &blockhash, &blocknum, &payload)
		if err != nil {
			return nil, err
		}
		arr = append(arr, Tx{txid, blockhash, blocknum, payload})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return arr, nil
}

// dummy stub
func (db *DB) GetByTxId(filter string) ([]Tx, error) {
	return nil, nil
}

// dummy stub
func (db *DB) GetByBlocknum(filter uint64) ([]Tx, error) {
	return nil, nil
}

// dummy stub
func (db *DB) GetBlockInfoByPayload(filter string) ([]Tx, error) {
	return nil, nil
}
