package main

import (
	"github.com/aclindsa/ofxgo"
)

type Balance struct {
	AcctNum string
	Balance string
	Date    string
}

func bankBalance(resp *ofxgo.Response) Balance {
	balance := Balance{}
	if stmt, ok := resp.Bank[0].(*ofxgo.StatementResponse); ok {
		balance.AcctNum = stmt.BankAcctFrom.AcctID.String()
		balance.Balance = stmt.BalAmt.String()
		balance.Date = stmt.AvailDtAsOf.UTC().Format("2006-01-02")
	}
	return balance
}

func creditCardBalance(resp *ofxgo.Response) Balance {
	balance := Balance{}
	if stmt, ok := resp.CreditCard[0].(*ofxgo.CCStatementResponse); ok {
		balance.AcctNum = stmt.CCAcctFrom.AcctID.String()
		balance.Balance = stmt.BalAmt.String()
		balance.Date = stmt.AvailDtAsOf.UTC().Format("2006-01-02")
		return balance
	}
	return balance
}
