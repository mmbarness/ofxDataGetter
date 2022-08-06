package main

import (
	"encoding/csv"
	"log"
	"os"
)

func writeTransactionArrayToCSV(transactions []Transaction, fileName string, headers []string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write(headers)

	for _, transaction := range transactions {
		row := []string{transaction.UID, transaction.Amount, transaction.Name, transaction.Posted, transaction.TransactionType}
		w.Write(row)
	}
}

func writeBalancesToFile(balances []BalanceAndAccountInfo, headers []string) {

	file, err := os.Create("csvs/balances/balances.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write(headers)

	for _, balanceAndAcct := range balances {
		row := []string{balanceAndAcct.Balances.AcctNum, balanceAndAcct.Balances.Date, balanceAndAcct.Balances.Balance}
		w.Write(row)
	}

}
