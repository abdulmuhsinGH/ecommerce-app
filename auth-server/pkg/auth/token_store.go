package auth

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

/*
TokenStore is an interface for this
*/
type TokenStore struct {
	db *pg.DB
}

/*
OauthToken is model for the oauth_clients table
*/
type OauthToken struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ExpiresAt time.Time `db:"expires_at"`
	Code      string    `db:"code"`
	Access    string    `db:"access"`
	Refresh   string    `db:"refresh"`
	Data      string    `db:"data"`
}

/*
NewTokenStore sets up the client store object
*/
func NewTokenStore(db *pg.DB) (*TokenStore, error) {
	return &TokenStore{db}, nil
}

/*
Create inserts token inf
*/
func (t *TokenStore) Create(info oauth2.TokenInfo) error {
	fmt.Println("hello create")
	tk, err := json.Marshal(info)
	if err != nil {
		return err
	}
	item := &OauthToken{
		Data: string(tk),
	}

	if code := info.GetCode(); code != "" {
		item.Code = code
		item.ExpiresAt = info.GetCodeCreateAt().Add(info.GetCodeExpiresIn())
	} else {
		item.Access = info.GetAccess()
		item.ExpiresAt = info.GetAccessCreateAt().Add(info.GetAccessExpiresIn())

		if refresh := info.GetRefresh(); refresh != "" {
			item.Refresh = info.GetRefresh()
			item.ExpiresAt = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn())
		}
	}

	return t.db.Insert(item)

}

/*
GetByAccess return client details using id
*/
func (t *TokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	fmt.Println(access)
	oauthToken := OauthToken{}
	err := t.db.Model(&oauthToken).
		Where("access = ?", access).
		Select()
	if err != nil {
		fmt.Println("aoc", err.Error())
		return nil, err
	}
	return t.toTokenInfo(oauthToken), nil

}

/*
GetByCode return client details using id
*/
func (t *TokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	fmt.Println(code)
	var oauthToken OauthToken

	err := t.db.Model(&oauthToken).
		Where("code = ?", code).
		Select()
	if err != nil {
		//fmt.Println("coc", err.Error())
		return nil, err
	}
	return t.toTokenInfo(oauthToken), nil

}

/*
GetByRefresh return client details using id
*/
func (t *TokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	fmt.Println(refresh)
	oauthToken := OauthToken{}
	err := t.db.Model(&oauthToken).
		Where("refresh = ?", refresh).
		Select()
	if err != nil {
		fmt.Println("roc", err.Error())
		return nil, err
	}
	return t.toTokenInfo(oauthToken), nil
}

/*
RemoveByAccess return client details using id
*/
func (t *TokenStore) RemoveByAccess(access string) error {
	oauthToken := OauthToken{Access: access}
	_, err := t.db.Model(&oauthToken).Where("access= ?access").Delete()
	return err
}

/*
RemoveByCode return client details using id
*/
func (t *TokenStore) RemoveByCode(code string) error {
	oauthToken := OauthToken{Code: code}
	return t.db.Delete(&oauthToken)
}

/*
RemoveByRefresh return client details using id
*/
func (t *TokenStore) RemoveByRefresh(refresh string) error {
	oauthToken := OauthToken{Refresh: refresh}
	_, err := t.db.Model(&oauthToken).Where("refresh= ?refresh").Delete()
	return err
}

func (t *TokenStore) toTokenInfo(data OauthToken) oauth2.TokenInfo {
	var tm models.Token
	err := json.Unmarshal([]byte(data.Data), &tm)
	if err != nil {
		return nil
	}
	return &tm

}
