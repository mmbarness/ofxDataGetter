package main

import (
	"fmt"
	"github.com/aclindsa/ofxgo"
)

func makeCreditCardStatementRequest(acctId string) *ofxgo.Response {
	client, query := newRequest()

	uid, err := ofxgo.RandomUID()

	if err != nil {
		fmt.Println("Error creating uid for transaction:", err)
		return &ofxgo.Response{}
	}

	statementRequest := ofxgo.CCStatementRequest{
		TrnUID: *uid,
		CCAcctFrom: ofxgo.CCAcct{
			AcctID: ofxgo.String(acctId),
		},
		Include: true,
	}

	query.CreditCard = append(query.CreditCard, &statementRequest)

	response, err := client.Request(query)

	if err != nil {
		fmt.Println("Error requesting account statement:", err)
	}

	return response; 
}

func makeBankStatementRequest(acctId, bankId, acctType string) *ofxgo.Response {
	client, query := newRequest()

	fmt.Println("acctType: ", acctType)

	acctTypeEnum, acctTypeError := ofxgo.NewAcctType(acctType)
	uid, uidError := ofxgo.RandomUID()

	if uidError != nil {
		fmt.Println("Error creating uid for transaction:", uidError)
		return &ofxgo.Response{}
	} else if acctTypeError != nil {
		fmt.Println("Error handling account type", acctTypeError)
		return &ofxgo.Response{}
	}

	statementRequest := ofxgo.StatementRequest{
		TrnUID: *uid,
		BankAcctFrom: ofxgo.BankAcct{
			BankID:   ofxgo.String(bankId),
			AcctID:   ofxgo.String(acctId),
			AcctType: acctTypeEnum,
		},
		Include: true,
	}

	query.Bank = append(query.Bank, &statementRequest)

	response, err := client.Request(query)

	if err != nil {
		fmt.Println("Error requesting account statement:", err)
		return &ofxgo.Response{}
	}

	return response
}