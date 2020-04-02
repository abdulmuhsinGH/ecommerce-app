package authstore

import (
	"encoding/json"
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
	db         *pg.DB
	gcDisabled bool
	gcInterval time.Duration
	ticker     *time.Ticker
}

/*
OauthToken is model for the oauth_clients table
*/
type OauthToken struct {
	ID        uuid.UUID `db:"id"`
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
	store := &TokenStore{
		db:         db,
		gcInterval: 10 * time.Minute,
	}

	/* if !store.gcDisabled {
		store.ticker = time.NewTicker(store.gcInterval)
		go store.gc()
	} */

	return store, nil

}

// Close close the store
func (t *TokenStore) Close() error {
	if !t.gcDisabled {
		t.ticker.Stop()
	}
	return nil
}

func (t *TokenStore) clean() {
	now := time.Now()
	ot := &OauthToken{
		ExpiresAt: now,
	}
	_, err := t.db.Model(ot).Where("expires_at<= ?expires_at").Delete() // Exec(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE expires_at <= $1", s.tableName), now)
	if err != nil {
		fmt.Println("Error deleting tokens")
	}
}

func (t *TokenStore) gc() {
	for range t.ticker.C {
		t.clean()
	}
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
	fmt.Println(item.Data)
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
	fmt.Println("c", code)
	oauthToken := OauthToken{
		Code: code,
	}

	err := t.db.Model(&oauthToken).
		Where("code = ?", code).
		Limit(1).
		Select()
	if err != nil {
		fmt.Println("coc", err.Error())
		return nil, err
	}
	fmt.Println(oauthToken.ID)
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
