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
TokenStoreInfo is model for the oauth_clients table
*/
type TokenStoreInfo struct {
	ID        uuid.UUID              `db:"id"`
	CreatedAt time.Time              `db:"created_at"`
	ExpiresIn time.Duration          `db:"expires_at"`
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
	oauthToken := TokenStoreInfo{
		Access:    info.GetAccess(),
		Code:      info.GetCode(),
		ExpiresIn: info.GetAccessExpiresIn(),
		Refresh:   info.GetRefresh(),
	}

	return t.db.Insert(&oauthToken)

}

/*
GetByAccess return client details using id
*/
func (t *TokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	oauthToken := TokenStoreInfo{Access: access}
	err := t.db.Select(&oauthToken)
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
	oauthToken := TokenStoreInfo{Code: code}
	err := t.db.Select(&oauthToken)
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
	oauthToken := TokenStoreInfo{Refresh: refresh}
	err := t.db.Select(&oauthToken)
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
	oauthToken := TokenStoreInfo{Access: access}
	return t.db.Delete(oauthToken)
}

/*
RemoveByCode return client details using id
*/
func (t *TokenStore) RemoveByCode(code string) error {
	oauthToken := TokenStoreInfo{Code: code}
	return t.db.Delete(oauthToken)
}

/*
RemoveByRefresh return client details using id
*/
func (t *TokenStore) RemoveByRefresh(refresh string) error {
	oauthToken := TokenStoreInfo{Refresh: refresh}
	return t.db.Delete(oauthToken)
}

func (t *TokenStore) toTokenInfo(data TokenStoreInfo) oauth2.TokenInfo {
	var tk models.Token
	tk.Access = data.Access
	tk.Code = data.Code
	tk.Refresh = data.Refresh
	// tk.AccessExpiresIn = data.ExpiresAt

	return &tk
}
