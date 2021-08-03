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

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
)

var azureAppsGetStmt, azureAppsDelStmt, azureAppsSetStmt *sql.Stmt
var azureSecretGetStmt, azureSecretSetStmt, azureSecretDelStmt *sql.Stmt

func azurePrepare() {

	azureAppsSetStmt = dbPrepSql(azureAppsSetStmtSQL)
	azureAppsGetStmt = dbPrepSql(azureAppsGetStmtSQL)
	azureAppsDelStmt = dbPrepSql(azureAppsDelStmtSQL)

	azureSecretSetStmt = dbPrepSql(azureSecretSetStmtSQL)
	azureSecretGetStmt = dbPrepSql(azureSecretGetStmtSQL)
	azureSecretDelStmt = dbPrepSql(azureSecretDelStmtSQL)
}

var azureSetupStmts = []string{
	`CREATE TABLE IF NOT EXISTS azureApps (
		"resourceGroupName" TEXT NOT NULL,
		"applicationName" TEXT NOT NULL,
		"data" BLOB,
		PRIMARY KEY ("resourceGroupName", "applicationName")
	);`,
	`CREATE TABLE IF NOT EXISTS azureSecrets (
		"key" TEXT NOT NULL PRIMARY KEY,
		"data" BLOB
	);`,
}

const azureAppsSetStmtSQL = `INSERT INTO azureApps(resourceGroupName, applicationName, data) VALUES (?, ?, ?);`

func AzureAppsSet(resourceGroupName string, applicationName string, value managedapplications.Application) {
	execStmt(azureAppsSetStmt, resourceGroupName, applicationName, toBytes(value))
}

const azureAppsGetStmtSQL = `SELECT data FROM azureApps WHERE resourceGroupName = ? and applicationName = ?;`

func AzureAppsGet(resourceGroupName string, applicationName string) *managedapplications.Application {
	value := managedapplications.Application{}
	if fromBytesRow(&value, azureAppsGetStmt, resourceGroupName, applicationName) != nil {
		return nil
	}
	return &value
}

const azureAppsDelStmtSQL = `DELETE FROM azureApps WHERE resourceGroupName = ? and applicationName = ?;`

func AzureAppsDel(resourceGroupName string, applicationName string) {
	execStmt(azureAppsDelStmt, resourceGroupName, applicationName)
}

const azureSecretSetStmtSQL = `INSERT INTO azureSecrets(key, data) VALUES (?, ?);`

func AzureSecretSet(vaultId string, secretID string, value string) {
	execStmt(azureSecretSetStmt, vaultId+"-"+secretID, toBytes(value))
}

const azureSecretGetStmtSQL = `SELECT data FROM azureSecrets WHERE key = ?;`

func AzureSecretGet(vaultId string, secretID string) *string {
	var value string
	if fromBytesRow(&value, azureSecretGetStmt, vaultId+"-"+secretID) != nil {
		return nil
	}
	return &value
}

const azureSecretDelStmtSQL = `DELETE FROM azureSecrets WHERE key = ?;`

func AzureSecretDel(vaultId string, secretID string) {
	execStmt(azureSecretDelStmt, vaultId+"-"+secretID)
}
