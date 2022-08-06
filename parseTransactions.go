package main

import (
	"fmt"

	"github.com/aclindsa/ofxgo"
)

type Transaction struct {
	UID             string
	Amount          string
	Name            string
	Posted          string
	TransactionType string
}

func parseBankTractions(resp *ofxgo.Response) []Transaction {
	var transactions []Transaction

	if len(resp.Bank) < 1 {
		fmt.Println("No banking messages received")
	}

	if stmt, ok := resp.Bank[0].(*ofxgo.StatementResponse); ok {

		for _, tran := range stmt.BankTranList.Transactions {

			transactionObject := Transaction{
				UID:             tran.FiTID.String(),
				Amount:          tran.TrnAmt.String(),
				Name:            tran.Name.String(),
				Posted:          tran.DtPosted.UTC().Format("2006-01-02"),
				TransactionType: tran.TrnType.String(),
			}

			transactions = append(transactions, transactionObject)
		}
	}

	return transactions
}

func parseCreditTransactions(resp *ofxgo.Response) []Transaction {
	var transactions []Transaction

	if len(resp.CreditCard) < 1 {
		fmt.Println("No banking messages received")
		return transactions
	}

	if stmt, ok := resp.CreditCard[0].(*ofxgo.CCStatementResponse); ok {

		for _, tran := range stmt.BankTranList.Transactions {
			var name string
			if len(tran.Name) > 0 {
				name = string(tran.Name)
			} else {
				name = string(tran.Payee.Name)
			}

			if len(tran.Memo) > 0 {
				name = name + " - " + string(tran.Memo)
			}

			transactionObject := Transaction{
				UID:             tran.FiTID.String(),
				Amount:          tran.TrnAmt.String(),
				Name:            name,
				Posted:          tran.DtPosted.UTC().Format("2006-01-02"),
				TransactionType: tran.TrnType.String(),
			}

			transactions = append(transactions, transactionObject)

		}
	}

	return transactions
}
