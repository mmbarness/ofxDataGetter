package main

import (
	"fmt"
	"os"
	"github.com/aclindsa/ofxgo"
)

func newRequest() (ofxgo.Client, *ofxgo.Request) {

	userInfo := useUserData()
	bankData := useBankData(userInfo.OrgId)

	ver, err := ofxgo.NewOfxVersion("102")
	if err != nil {
		fmt.Println("Error creating new OfxVersion enum:", err)
		os.Exit(1)
	}
	var client = ofxgo.GetClient(bankData.Url,
		&ofxgo.BasicClient{
			AppID:          bankData.AppId,
			AppVer:         bankData.AppVer,
			SpecVersion:    ver,
			NoIndent:       true,
			CarriageReturn: true,
		})

	var query ofxgo.Request
	query.URL = bankData.Url
	query.Signon.ClientUID = ofxgo.UID(userInfo.Uid)
	query.Signon.UserID = ofxgo.String(userInfo.UserId)
	query.Signon.UserPass = ofxgo.String(userInfo.UserPass)
	query.Signon.Org = ofxgo.String(userInfo.OrgId)
	query.Signon.Fid = ofxgo.String(bankData.FID)

	return client, &query
}
