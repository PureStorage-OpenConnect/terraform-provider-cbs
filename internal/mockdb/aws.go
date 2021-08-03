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
	"database/sql"

	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var awsGetStackStmt, awsDelStackStmt, awsSetStackStmt, awsGetHostStackStmt *sql.Stmt
var awsGetSecretStmt, awsSetSecretStmt *sql.Stmt

func awsPrepare() {
	awsSetStackStmt = dbPrepSql(awsSetStackStmtSQL)
	awsGetStackStmt = dbPrepSql(awsGetStackStmtSQL)
	awsDelStackStmt = dbPrepSql(awsDelStackStmtSQL)
	awsGetHostStackStmt = dbPrepSql(awsGetHostStackStmtSQL)

	awsSetSecretStmt = dbPrepSql(awsSetSecretStmtSQL)
	awsGetSecretStmt = dbPrepSql(awsGetSecretStmtSQL)
}

var awsSetupStmts = []string{
	`CREATE TABLE IF NOT EXISTS awsStacks (
		"idOrName" TEXT NOT NULL PRIMARY KEY,
		"ip" TEXT NOT NULL,
		"data" BLOB
	);`,
	`CREATE TABLE IF NOT EXISTS awsSecrets (
		"secretId" TEXT NOT NULL PRIMARY KEY,
		"data" BLOB
	);`,
}

const awsSetStackStmtSQL = `INSERT INTO awsStacks(idOrName, ip, data) VALUES (?, ?, ?);`

func AWSSetStack(key string, ip string, value cloudformation.Stack) {
	execStmt(awsSetStackStmt, key, ip, toBytes(value))
}

const awsGetStackStmtSQL = `SELECT data FROM awsStacks WHERE idOrName = ?;`

func AWSGetStack(key string) *cloudformation.Stack {
	value := cloudformation.Stack{}
	if fromBytesRow(&value, awsGetStackStmt, key) != nil {
		return nil
	}
	return &value
}

const awsDelStackStmtSQL = `DELETE FROM awsStacks WHERE idOrName = ?;`

func AWSDelStack(key string) {
	execStmt(awsDelStackStmt, key)
}

const awsGetHostStackStmtSQL = `SELECT data FROM awsStacks WHERE ip = ?;`

func AWSGetHostStack(key string) *cloudformation.Stack {
	value := cloudformation.Stack{}
	if fromBytesRow(&value, awsGetHostStackStmt, key) != nil {
		return nil
	}
	return &value
}

const awsSetSecretStmtSQL = `INSERT INTO awsSecrets(secretId, data) VALUES (?, ?);`

func AWSSetSecret(key string, value string) {
	execStmt(awsSetSecretStmt, key, toBytes(value))
}

const awsGetSecretStmtSQL = `SELECT data FROM awsSecrets WHERE secretId = ?;`

func AWSGetSecret(key string) *string {
	var value string
	if fromBytesRow(&value, awsGetSecretStmt, key) != nil {
		return nil
	}
	return &value
}
