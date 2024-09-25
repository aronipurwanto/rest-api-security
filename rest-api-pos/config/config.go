package config

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	App struct {
		Port string
	}
	DB struct {
		User     string
		Password string
		Host     string
		Name     string
	}
	Keycloak struct {
		URL           string
		ClientID      string
		ClientSecret  string
		Realm         string
		TokenEndpoint string
	}
}

var AppConfig Config
var DB *gorm.DB
var OIDCVerifier *oidc.IDTokenVerifier

func LoadConfig() {
	// Viper untuk memuat file konfigurasi
	viper.SetConfigName("config/config") // Nama file konfigurasi (tanpa ekstensi)
	viper.SetConfigType("yaml")          // Tipe file konfigurasi (yaml)
	viper.AddConfigPath(".")             // Direktori tempat mencari file (direktori saat ini)

	// Coba baca file konfigurasi
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal file konfigurasi ke struct
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// Inisialisasi koneksi database dengan GORM
	dsn := AppConfig.DB.User + ":" + AppConfig.DB.Password + "@tcp(" + AppConfig.DB.Host + ")/" + AppConfig.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Inisialisasi OIDC verifier untuk Keycloak
	provider, err := oidc.NewProvider(context.Background(), AppConfig.Keycloak.URL+"/realms/"+AppConfig.Keycloak.Realm)
	if err != nil {
		log.Fatalf("Failed to initialize OIDC provider: %v", err)
	}
	OIDCVerifier = provider.Verifier(&oidc.Config{ClientID: AppConfig.Keycloak.ClientID})
}
