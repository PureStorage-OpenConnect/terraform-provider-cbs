// +build mock

/*

	Copyright 2021, Pure Storage Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.


	This file contains code to support on-disk persistence for mocked versions
	of the cloud service calls.  Persistence for mocks enables mocked tests
	over multiple terraform command invocations, such as manual testing.

*/

package mockdb

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mattn/go-sqlite3"
)

var mockDb *sql.DB

func init() {

	{
		mockDbPath := os.Getenv("MOCK_DB_PATH")
		err := os.MkdirAll(filepath.Dir(mockDbPath), 0755)
		if err != nil {
			panic(err)
		}
		db, err := sql.Open("sqlite3", mockDbPath)
		if err != nil {
			fmt.Printf("Could not open DB path:%#v err:%#v\n", mockDbPath, err)
			panic(err)
		}
		mockDb = db
	}

	for _, stmt := range awsSetupStmts {
		execSQL(stmt)
	}
	for _, stmt := range azureSetupStmts {
		execSQL(stmt)
	}

	awsPrepare()
	azurePrepare()

	gob.Register(map[string]map[string]interface{}{})
	gob.Register(map[string]map[string]string{})
	gob.Register(map[string]string{})
}

func dbPrepSql(sql string) *sql.Stmt {
	statement, err := mockDb.Prepare(sql)
	if err != nil {
		fmt.Printf("dbPrepSql `%s` %#v\n", sql, err)
		panic(nil)
	}
	return statement
}

func execStmt(stmt *sql.Stmt, queryParameters ...interface{}) {
	_, err := stmt.Exec(queryParameters...)
	if err != nil {
		fmt.Printf("%#v\n", err)
		fmt.Printf("HINT: If mock SQL statements are failing and you aren't actively iterating on them, then it is likely that you just have an old schema, and need to run `make reset-mock-db`\n")
		panic(err)
	}
}

func execSQL(sql string) {
	execStmt(dbPrepSql(sql))
}

func toBytes(value interface{}) []byte {
	var blob bytes.Buffer
	enc := gob.NewEncoder(&blob)
	err := enc.Encode(value)
	if err != nil {
		fmt.Printf("%#v\n", err)
		panic(err)
	}
	return blob.Bytes()
}

func fromBytesRow(value interface{}, stmt *sql.Stmt, queryParameters ...interface{}) error {
	row := stmt.QueryRow(queryParameters...)
	var _bytes []uint8

retry:
	err := row.Scan(&_bytes)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		fmt.Printf("%#v\n", err)

		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrBusy {
				fmt.Println("Sqlite database locked, will retry")
				time.Sleep(100 * time.Millisecond)
				goto retry
			}
		}
		panic(err)
	}
	blob := bytes.NewBuffer(_bytes)
	dec := gob.NewDecoder(blob)
	err = dec.Decode(value)
	if err != nil {
		fmt.Printf("%#v\n", err)
		panic(err)
	}
	return nil
}
