package auth

import (
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
TokenStoreInfo is model for the oauth_clients table
*/
type TokenStoreInfo struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
	Code      string    `db:"code"`
	Access    string    `db:"access"`
	Refresh   string    `db:"refresh"`
	Data      []byte    `db:"data"`
}

/*
NewTokenStore sets up the client store object
*/
func NewTokenStore(db *pg.DB) *TokenStore {
	return &TokenStore{db}
}

/*
GetByAccess return client details using id
*/
func (t *TokenStore) GetByAccess(Access string) (oauth2.TokenInfo, error) {
	oauthToken := TokenStoreInfo{Access: Access}
	err := t.db.Select(&oauthToken)
	if err != nil {
		fmt.Println("oc", err.Error())
		return nil, err
	}
	clientInfo := t.toTokenInfo(oauthToken)
	if err != nil {
		fmt.Println("ci", err.Error())
		return nil, err
	}
	return clientInfo, nil
}

func (t *TokenStore) toTokenInfo(data TokenStoreInfo) oauth2.TokenInfo {
	var tk models.Token
	tk.Access = data.Access
	tk.Code = data.Code
	tk.Refresh = data.Refresh
	// tk.AccessExpiresIn = data.ExpiresAt

	return &tk
}
