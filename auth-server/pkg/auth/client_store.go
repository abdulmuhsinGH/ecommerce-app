package auth

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

/*
ClientStore is an interface for this
*/
type ClientStore struct {
	db *pg.DB
}

/*
OauthClient is model for the oauth_clients table
*/
type OauthClient struct {
	ID     string                 `db:"id"`
	Secret string                 `db:"secret"`
	Domain string                 `db:"domain"`
	Data   map[string]interface{} `db:"data"`
}

/*
NewClientStore sets up the client store object
*/
func NewClientStore(db *pg.DB) *ClientStore {
	return &ClientStore{db}
}

/*
GetByID return client details using id
*/
func (c *ClientStore) GetByID(ID string) (oauth2.ClientInfo, error) {
	oauthClient := OauthClient{ID: ID}
	err := c.db.Select(&oauthClient)
	if err != nil {
		fmt.Println("oc", err.Error())
		return nil, err
	}
	fmt.Println(oauthClient)
	clientInfo := c.toClientInfo(oauthClient)
	if err != nil {
		fmt.Println("ci", err.Error())
		return nil, err
	}
	return clientInfo, nil
}

func (c *ClientStore) toClientInfo(data OauthClient) oauth2.ClientInfo {
	var cm models.Client
	cm.ID = data.ID
	cm.Secret = data.Secret
	cm.Domain = data.Domain
	//err := json.Unmarshal(data, &cm)
	return &cm
}
