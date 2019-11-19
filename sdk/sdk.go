package sdk

import (
	"errors"
	"fmt"
)

type App struct {
	AppId string
	AppSecret string
	AppAccessToken string
}

type Tenant struct {
	TenantAccessToken string
}

func BuildInternalApp(appId, appSecret string) (*App, error){
	resp, err := GetAppAccessTokenInternal(appId, appSecret)
	if err != nil{
		return nil, err
	}
	if resp.Code != 0{
		return nil, errors.New(fmt.Sprintf("req err, code: %d, msg: %s", resp.Code, resp.Msg))
	}
	return &App{
		AppId: appId,
		AppSecret: appSecret,
		AppAccessToken: appSecret,
	}, nil
}

func BuildApp(appId, appSecret, appTicket string) (*App, error){
	resp, err := GetAppAccessToken(appId, appSecret, appTicket)
	if err != nil{
		return nil, err
	}
	if resp.Code != 0{
		return nil, errors.New(fmt.Sprintf("req err, code: %d, msg: %s", resp.Code, resp.Msg))
	}
	return &App{
		AppId: appId,
		AppSecret: appSecret,
		AppAccessToken: resp.AppAccessToken,
	}, nil
}

func BuildTenantInternal(appId, appSecret string) (*Tenant, error){
	resp, err := GetTenantAccessTokenInternal(appId, appSecret)
	if err != nil{
		return nil, err
	}
	if resp.Code != 0{
		return nil, errors.New(fmt.Sprintf("req err, code: %d, msg: %s", resp.Code, resp.Msg))
	}
	return &Tenant{
		TenantAccessToken: resp.TenantAccessToken,
	}, nil
}

func BuildTenant(appAccessToken, tenantKey string) (*Tenant, error){
	resp, err := GetTenantAccessToken(appAccessToken, tenantKey)
	if err != nil{
		return nil, err
	}
	if resp.Code != 0{
		return nil, errors.New(fmt.Sprintf("req err, code: %d, msg: %s", resp.Code, resp.Msg))
	}
	return &Tenant{
		TenantAccessToken: resp.TenantAccessToken,
	}, nil
}