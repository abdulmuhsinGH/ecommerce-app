package auth

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	uuid "github.com/satori/go.uuid"
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
	ID        uuid.UUID              `db:"id"`
	CreatedAt time.Time              `db:"created_at"`
	ExpiresAt time.Time              `db:"expires_at"`
	Code      string                 `db:"code"`
	Access    string                 `db:"access"`
	Refresh   string                 `db:"refresh"`
	Data      map[string]interface{} `db:"data"`
}

/*
NewTokenStore sets up the client store object
*/
func NewTokenStore(db *pg.DB) *TokenStore {
	return &TokenStore{db}
}

/*
Create inserts token inf
*/
func (t *TokenStore) Create(info oauth2.TokenInfo) error {

	oauthToken := &OauthToken{
	}
	fmt.Println(info.GetCode(), info.GetAccess(), info.GetRefresh(), info.GetAccessExpiresIn())

	if code := info.GetCode(); code != "" {
		fmt.Println("Hello")
		oauthToken.Code = code
		oauthToken.ExpiresAt = info.GetCodeCreateAt().Add(info.GetCodeExpiresIn())

	} else {
		fmt.Println("Hello World")
		oauthToken.Access = info.GetAccess()
		oauthToken.ExpiresAt = info.GetAccessCreateAt().Add(info.GetAccessExpiresIn())

		if refresh := info.GetRefresh(); refresh != "" {
			fmt.Println("Hello World 2")
			oauthToken.Refresh = info.GetRefresh()
			oauthToken.ExpiresAt = info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn())
		}
	}
	err := t.db.Model(&oauthToken).
		Where("oauth_tokens.access = ?", oauthToken.Access).
		Select()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(err.Error())
	return err

}

/*
GetByAccess return client details using id
*/
func (t *TokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	oauthToken := OauthToken{Access: access}
	err := t.db.Model(&oauthToken).
		Where("oauth_tokens.access = ?", oauthToken.Access).
		Select()
	if err != nil {
		fmt.Println("oc", err.Error())
		return nil, err
	}
	tokenInfo := t.toTokenInfo(oauthToken)
	if err != nil {
		fmt.Println("ci", err.Error())
		return nil, err
	}
	return tokenInfo, nil
}

/*
GetByCode return client details using id
*/
func (t *TokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	oauthToken := OauthToken{Code: code}
	err := t.db.Model(&oauthToken).
		Where("oauth_tokens.code = ?", oauthToken.Code).
		Select()
	if err != nil {
		fmt.Println("oc", err.Error())
		return nil, err
	}
	tokenInfo := t.toTokenInfo(oauthToken)
	if err != nil {
		fmt.Println("ci", err.Error())
		return nil, err
	}
	return tokenInfo, nil
}

/*
GetByRefresh return client details using id
*/
func (t *TokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	oauthToken := OauthToken{Refresh: refresh}
	err := t.db.Model(&oauthToken).
						Where("oauth_tokens.refresh = ?", oauthToken.Refresh).
						Select()
	if err != nil {
		fmt.Println("oc", err.Error())
		return nil, err
	}
	tokenInfo := t.toTokenInfo(oauthToken)
	if err != nil {
		fmt.Println("ci", err.Error())
		return nil, err
	}
	return tokenInfo, nil
}

/*
RemoveByAccess return client details using id
*/
func (t *TokenStore) RemoveByAccess(access string) error {
	oauthToken := OauthToken{Access: access}
	return t.db.Delete(&oauthToken)
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
	return t.db.Delete(&oauthToken)
}

func (t *TokenStore) toTokenInfo(data OauthToken) oauth2.TokenInfo {
	var tk models.Token
	tk.SetAccess(data.Access)
	tk.SetCode(data.Code)
	tk.SetRefresh(data.Refresh)
	tk.SetAccessExpiresIn(data.CreatedAt.Sub(data.ExpiresAt))

	return &tk
}
