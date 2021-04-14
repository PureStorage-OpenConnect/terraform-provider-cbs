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

package cbs

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	_ "github.com/mattn/go-sqlite3"
)

var mockDb *sql.DB
var mockDbAWSGetStackStmt, mockDbAWSDelStackStmt, mockDbAWSSetStackStmt *sql.Stmt
var mockDbAzureAppsGetStmt, mockDbAzureAppsDelStmt, mockDbAzureAppsSetStmt *sql.Stmt

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

	execSQL(`
		CREATE TABLE IF NOT EXISTS awsStacks (
			"idOrName" TEXT NOT NULL PRIMARY KEY,
			"data" BLOB
		);`)
	execSQL(`
		CREATE TABLE IF NOT EXISTS azureApps (
			"resourceGroupName" TEXT NOT NULL,
			"applicationName" TEXT NOT NULL,
			"data" BLOB,
			PRIMARY KEY ("resourceGroupName", "applicationName")
		);`)

	mockDbAWSSetStackStmt = dbPrepSql(mockDbAWSSetStackStmtSQL)
	mockDbAWSGetStackStmt = dbPrepSql(mockDbAWSGetStackStmtSQL)
	mockDbAWSDelStackStmt = dbPrepSql(mockDbAWSDelStackStmtSQL)

	mockDbAzureAppsSetStmt = dbPrepSql(mockDbAzureAppsSetStmtSQL)
	mockDbAzureAppsGetStmt = dbPrepSql(mockDbAzureAppsGetStmtSQL)
	mockDbAzureAppsDelStmt = dbPrepSql(mockDbAzureAppsDelStmtSQL)

	gob.Register(map[string]map[string]interface{}{})
	gob.Register(map[string]map[string]string{})
	gob.Register(map[string]string{})

}

const mockDbAWSSetStackStmtSQL = `INSERT INTO awsStacks(idOrName, data) VALUES (?, ?);`

func mockDbAWSSetStack(key string, value cloudformation.Stack) {
	execStmt(mockDbAWSSetStackStmt, key, toBytes(value))
}

const mockDbAWSGetStackStmtSQL = `SELECT data FROM awsStacks WHERE idOrName = ?;`

func mockDbAWSGetStack(key string) *cloudformation.Stack {
	value := cloudformation.Stack{}
	if fromBytesRow(&value, mockDbAWSGetStackStmt, key) != nil {
		return nil
	}
	return &value
}

const mockDbAWSDelStackStmtSQL = `DELETE FROM awsStacks WHERE idOrName = ?;`

func mockDbAWSDelStack(key string) {
	execStmt(mockDbAWSDelStackStmt, key)
}

const mockDbAzureAppsSetStmtSQL = `INSERT INTO azureApps(resourceGroupName, applicationName, data) VALUES (?, ?, ?);`

func mockDbAzureAppsSet(resourceGroupName string, applicationName string, value managedapplications.Application) {
	execStmt(mockDbAzureAppsSetStmt, resourceGroupName, applicationName, toBytes(value))
}

const mockDbAzureAppsGetStmtSQL = `SELECT data FROM azureApps WHERE resourceGroupName = ? and applicationName = ?;`

func mockDbAzureAppsGet(resourceGroupName string, applicationName string) *managedapplications.Application {
	value := managedapplications.Application{}
	if fromBytesRow(&value, mockDbAzureAppsGetStmt, resourceGroupName, applicationName) != nil {
		return nil
	}
	return &value
}

const mockDbAzureAppsDelStmtSQL = `DELETE FROM azureApps WHERE resourceGroupName = ? and applicationName = ?;`

func mockDbAzureAppsDel(resourceGroupName string, applicationName string) {
	execStmt(mockDbAzureAppsDelStmt, resourceGroupName, applicationName)
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

	err := row.Scan(&_bytes)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		fmt.Printf("%#v\n", err)
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
