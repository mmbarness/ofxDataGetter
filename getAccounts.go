package main

import (
	"fmt"
	"os"
	"time"
	"github.com/aclindsa/ofxgo"
)

func mapSlice[T any, M any](a []T, mapFn func(T) M) []M {
    mappedArr := make([]M, len(a))
    for i, ele := range a {
		newEle := mapFn(ele)
		print("newEle:  ", newEle)
        mappedArr[i] = newEle
    }
    return mappedArr
}

func getAccounts() []map[string]string {
	client, query := newRequest()

	uid, err := ofxgo.RandomUID()
	if err != nil {
		fmt.Println("Error creating uid for transaction:", err)
		os.Exit(1)
	}

	acctInfo := ofxgo.AcctInfoRequest{
		TrnUID:   *uid,
		DtAcctUp: ofxgo.Date{Time: time.Unix(0, 0)},
	}
	query.Signup = append(query.Signup, &acctInfo)

	response, err := client.Request(query)

	if err != nil {
		return []map[string]string{}
	}

	matchId := func(acct ofxgo.AcctInfo) map[string]string {
		if (acct.BankAcctInfo != nil) {
			bankInfo := map[string]string {"type": "bank", "acctId": acct.BankAcctInfo.BankAcctFrom.AcctID.String(), "bankId": acct.BankAcctInfo.BankAcctFrom.BankID.String(),  "acctType": acct.BankAcctInfo.BankAcctFrom.AcctType.String()}
			return bankInfo
		} else if (acct.CCAcctInfo != nil) {
			return map[string]string{"type": "credit_card", "acctId": acct.CCAcctInfo.CCAcctFrom.AcctID.String()}
		} else {
			return map[string]string{"type": "unknown", "id": "unknown"}
		}
	}

	if acctinfo, ok := response.Signup[0].(*ofxgo.AcctInfoResponse); ok {
		return mapSlice(acctinfo.AcctInfo, matchId)
	} else {
		return []map[string]string{}
	}
}
