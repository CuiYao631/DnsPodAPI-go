package usecase

import (
	"DDNS-c/controller/entity"
	"bytes"
	"encoding/xml"
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
	ListRecord() entity.Dnspod
	SetUpRecord(ddnsinfo entity.DdnsInfo) ([]byte, error)

	RemoveRecord() error
	DdnsRecord() error
	GetRecordInfo() error

	//*******待添加********
}
type usecase struct {
	http *http.Client
	host string
}

func MakeUsecase(http *http.Client, host string) Usecase {
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
func (uc *usecase) ListRecord() entity.Dnspod {

	url := "https://dnsapi.cn/Record.List"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("login_token", "240684,b61a5717428d23b9122efc6bf4237a67")
	_ = writer.WriteField("form", "json")
	_ = writer.WriteField("domain", "xcuitech.com")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)

	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	var response entity.Dnspod
	//解析返回数据
	jsonErr := xml.Unmarshal(body, &response)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return response
}
func (uc *usecase) SetUpRecord(ddnsinfo entity.DdnsInfo) ([]byte, error) {

	response := uc.ListRecord()
	RecordID := getRecordID(ddnsinfo.SubDomain, response)
	//url := uc.host + "/Record.Ddns"
	url := "https://dnsapi.cn/Record.Ddns"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("login_token", ddnsinfo.LoginToken)
	_ = writer.WriteField("form", ddnsinfo.Form)
	_ = writer.WriteField("domain", ddnsinfo.Domain)
	_ = writer.WriteField("record_id", RecordID)
	_ = writer.WriteField("sub_domain", ddnsinfo.SubDomain)
	_ = writer.WriteField("record_line", ddnsinfo.RecordLine)
	_ = writer.WriteField("value", ddnsinfo.Value)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := uc.http.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%s\n", bodyText)
	return bodyText, nil
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

func getRecordID(SubDomain string, dsp entity.Dnspod) string {
	list := dsp.Records.Item
	for i := 0; i < len(list); i++ {
		if list[i].Name == SubDomain {
			return list[i].ID
		}
	}
	return ""
}
