package services

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// OAuth2Config menyimpan konfigurasi OAuth 2.0
var OAuth2Config *oauth2.Config

// InitOAuthConfig menginisialisasi konfigurasi OAuth2 dari viper
func InitOAuthConfig() {
	OAuth2Config = &oauth2.Config{
		ClientID:     viper.GetString("oauth.client_id"),
		ClientSecret: viper.GetString("oauth.client_secret"),
		RedirectURL:  viper.GetString("oauth.redirect_url"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"openid", "profile", "email"},
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
