package services

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// OAuth2Config menyimpan konfigurasi OAuth 2.0
var OAuth2Config *oauth2.Config

// InitOAuthConfig menginisialisasi konfigurasi OAuth2 dari viper
func InitOAuthConfig() {
	OAuth2Config = &oauth2.Config{
		ClientID:     viper.GetString("oauth.client_id"),
		ClientSecret: viper.GetString("oauth.client_secret"),
		RedirectURL:  viper.GetString("oauth.redirect_url"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  viper.GetString("oauth.auth_url"),
			TokenURL: viper.GetString("oauth.token_url"),
		},
		Scopes: []string{"read", "write"},
	}
}

// ExchangeCode menukar authorization code dengan access token
func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := OAuth2Config.Exchange(ctx, code)
	if err != nil {
		logrus.Errorf("Failed to exchange token: %v", err)
		return nil, err
	}
	logrus.Infof("Successfully exchanged token: %+v", token)
	return token, nil
}
