package controller

import (
	"DDNS-c/controller/entity"
	"DDNS-c/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type Controller interface {
	//*******账户相关*******

	GetUserInfo(echo.Context) error
	GetUserLog(echo.Context) error

	//*******域名相关********
	CreateDomian(echo.Context) error
	ListDomain(echo.Context) error
	RemoveDomain(echo.Context) error
	SetDomainStatus(echo.Context) error
	GetDomainInfo(echo.Context) error
	GetDomainlog(echo.Context) error

	//*******记录相关********

	CreateRecord(echo.Context) error
	ListRecord(echo.Context) error
	SetUpRecord(echo.Context) error
	RemoveRecord(echo.Context) error
	DdnsRecord(echo.Context) error
	GetRecordInfo(echo.Context) error

	//*******待添加********
	//test
	Test(echo.Context) error
}
type controller struct {
	uc usecase.Usecase
}

func MakeController(uc usecase.Usecase) Controller {
	return &controller{uc: uc}
}
func (crtl *controller) GetUserInfo(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) GetUserLog(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) CreateDomian(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	var user entity.User
	user.Name = name
	user.Email = email
	return echo.NewHTTPError(http.StatusOK, user)
}
func (crtl *controller) ListDomain(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) RemoveDomain(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) SetDomainStatus(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) GetDomainInfo(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) GetDomainlog(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) CreateRecord(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) ListRecord(c echo.Context) error {

	res := crtl.uc.ListRecord()
	return echo.NewHTTPError(http.StatusOK, res)
}

func (crtl *controller) SetUpRecord(c echo.Context) error {
	var ddnsinfo entity.DdnsInfo
	ddnsinfo.LoginToken = c.QueryParam("login_token")
	ddnsinfo.Form = c.QueryParam("form")
	ddnsinfo.Domain = c.QueryParam("domain")
	ddnsinfo.SubDomain = c.QueryParam("sub_domain")
	ddnsinfo.RecordLine = c.QueryParam("record_line")
	ddnsinfo.Value = c.QueryParam("value")

	_, err := crtl.uc.SetUpRecord(ddnsinfo)
	if err != nil {
		return echo.NewHTTPError(400, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func (crtl *controller) RemoveRecord(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) DdnsRecord(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}

func (crtl *controller) GetRecordInfo(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) Test(c echo.Context) error{
	
	 return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
}
