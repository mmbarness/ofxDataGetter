package main

import (
	"github.com/aclindsa/ofxgo"
)

type TransactionAndAccountInfo struct {
	Transactions []Transaction
	AccountInfo  map[string]string
}

type StatementAndAccountInfo struct {
	Statement *ofxgo.Response
	AccountInfo map[string]string
}

type BalanceAndAccountInfo struct {
	Balances Balance
	AccountInfo  map[string]string
}

func main() {

	accounts := getAccounts()

	statements := []StatementAndAccountInfo{}

	for _, account := range accounts {
		var stmtAndAcct StatementAndAccountInfo
		stmtAndAcct.AccountInfo = account
		if account["type"] == "bank" {
			stmt := makeBankStatementRequest(account["acctId"], account["bankId"], account["acctType"])
			stmtAndAcct.Statement = stmt
		} else if account["type"] == "credit_card" {
			stmt := makeCreditCardStatementRequest(account["acctId"])
			stmtAndAcct.Statement = stmt
		}
		statements = append(statements, stmtAndAcct)
	}

	transactions := []TransactionAndAccountInfo{}

	for _, stmtAndAcct := range statements {
		var transactionsAndAccount TransactionAndAccountInfo
		transactionsAndAccount.AccountInfo = stmtAndAcct.AccountInfo
		if stmtAndAcct.AccountInfo["type"] == "bank" {
			transactionsAndAccount.Transactions = parseBankTractions(stmtAndAcct.Statement)
		} else if stmtAndAcct.AccountInfo["type"] == "credit_card" {
			transactionsAndAccount.Transactions = parseCreditTransactions(stmtAndAcct.Statement)
		}
		transactions = append(transactions, transactionsAndAccount)
	}

	balances := []BalanceAndAccountInfo{}

	for _, stmtAndAcct := range statements {
		var balanceAndAcct BalanceAndAccountInfo
		balanceAndAcct.AccountInfo = stmtAndAcct.AccountInfo
		if stmtAndAcct.AccountInfo["type"] == "bank" {
			balanceAndAcct.Balances = bankBalance(stmtAndAcct.Statement)
		} else if stmtAndAcct.AccountInfo["type"] == "credit_card" {
			balanceAndAcct.Balances = creditCardBalance(stmtAndAcct.Statement)
		}
		balances = append(balances, balanceAndAcct)
	}

	transactionHeaders := []string{"UID", "amount", "name", "posted", "transactionType"}
	balanceHeaders := []string{"account number","date", "balance"}

	for _, allInfo := range transactions {
		writeTransactionArrayToCSV(allInfo.Transactions, "csvs/transactions/"+allInfo.AccountInfo["acctId"]+".csv", transactionHeaders)
	}

	writeBalancesToFile(balances, balanceHeaders)

}
