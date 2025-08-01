package config

import "github.com/kelseyhightower/envconfig"

type (
	Specification struct {
		Database  DatabaseSpecification
		Gmailsmtp GmailsmtpSpecification
	}

	DatabaseSpecification struct {
		Host     string `envconfig:"HOST" required:"true"`
		Port     string `envconfig:"PORT" required:"true"`
		User     string `envconfig:"USER" required:"true"`
		Password string `envconfig:"PASSWORD" required:"true"`
		DbName   string `envconfig:"DBNAME" required:"true"`
		SslMode  string `envconfig:"SSLMODE" required:"true"`
	}
	GmailsmtpSpecification struct {
		Email    string `envconfig:"GMAIL" required:"true"`
		Password string `envconfig:"GMAILPASSWORD" required:"true"`
	}
)

var Settings Specification

func Init() {
	if err := envconfig.Process("", &Settings); err != nil {
		panic(err.Error())
	}
}
