package usecase

import (
	"DDNS-c/controller/entity"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type Usecase interface {
	//*******账户相关*******

	GetUserInfo() error
	GetUserLog() error

	//*******域名相关********
	CreateDomian() error
	ListDomain() error
	RemoveDomain() error
	SetDomainStatus() error
	GetDomainInfo() error
	GetDomainlog() error

	//*******记录相关********

	CreateRecord() error
	ListRecord() error
	SetUpRecord(ddnsinfo entity.DdnsInfo) error
	RemoveRecord() error
	DdnsRecord() error
	GetRecordInfo() error

	//*******待添加********
}
type usecase struct {
	http *http.Client
	host string
}

func MakeUsecase(http *http.Client, host string) *usecase {
	return &usecase{http: http, host: host}
}
func (uc *usecase) GetUserInfo() error {

	return nil
}
func (uc *usecase) GetUserLog() error {
	return nil
}
func (uc *usecase) CreateDomian() error {
	return nil
}
func (uc *usecase) ListDomain() error {
	return nil
}
func (uc *usecase) RemoveDomain() error {
	return nil
}
func (uc *usecase) SetDomainStatus() error {
	return nil
}
func (uc *usecase) GetDomainInfo() error {
	return nil
}
func (uc *usecase) GetDomainlog() error {
	return nil
}
func (uc *usecase) CreateRecord() error {
	return nil
}
func (uc *usecase) ListRecord() error {
	return nil
}
func (uc *usecase) SetUpRecord(ddnsinfo entity.DdnsInfo) error {

	url := uc.host + "/Record.Ddns"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("login_token", ddnsinfo.LoginToken)
	_ = writer.WriteField("form", ddnsinfo.Form)
	_ = writer.WriteField("domain_id", ddnsinfo.Domain)
	_ = writer.WriteField("record_id", ddnsinfo.RecordID)
	_ = writer.WriteField("sub_domain", ddnsinfo.SubDomain)
	_ = writer.WriteField("record_line", ddnsinfo.RecordLine)
	_ = writer.WriteField("value", ddnsinfo.Value)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := uc.http.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}
func (uc *usecase) RemoveRecord() error {
	return nil
}
func (uc *usecase) DdnsRecord() error {
	return nil
}
func (uc *usecase) GetRecordInfo() error {
	return nil
}
