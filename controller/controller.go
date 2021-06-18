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
}
type controller struct {
	uc usecase.Usecase
}

func MakeController(uc usecase.Usecase) *controller {
	return &controller{uc: uc}
}
func (crtl *controller) GetUserInfo(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) GetUserLog(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
}
func (crtl *controller) CreateDomian(echo.Context) error {

	return echo.NewHTTPError(http.StatusOK)
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

	return echo.NewHTTPError(http.StatusOK)
}

func (crtl *controller) SetUpRecord(c echo.Context) error {

	var ddnsinfo entity.DdnsInfo
	if err := c.Bind(&ddnsinfo); err != nil {
		return echo.NewHTTPError(400, "Parameter Faild")
	}
	crtl.uc.SetUpRecord(ddnsinfo)
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
