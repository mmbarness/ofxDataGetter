package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserInfo struct {
	Uid      string `json:"uid"`
	UserId   string `json:"userId"`
	UserPass string `json:"userPass"`
	OrgId    string `json:"orgId"`
}

type BankData struct {
	FID    string `json:"fid"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	AppId  string `json:"appId"`
	AppVer string `json:"appVer"`
}

func useUserData() UserInfo {
	jsonFile, err := os.Open("configs/myInfo.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var userInfo UserInfo

	json.Unmarshal(byteValue, &userInfo)

	fmt.Println("userInfo: ", userInfo)

	return userInfo
}

func useBankData(id string) BankData {
	jsonFile, err := os.Open("configs/institutionInfo.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var bankData map[string]BankData

	json.Unmarshal(byteValue, &bankData)

	fmt.Println("bankData: ", bankData)

	return bankData[id]
}
