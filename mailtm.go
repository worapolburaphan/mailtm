package mailtm

import (
	"net/http"
)

type Service string

const (
	MailTm Service = "https://api.mail.tm"
)

type MailClient struct {
	http    *http.Client
	service Service
}

type MailClientOptions struct {
	Http    *http.Client
	Service Service
}

func NewWithOption(options *MailClientOptions) (*MailClient, error) {
	if options == nil {
		options = &MailClientOptions{}
	}
	if options.Http == nil {
		options.Http = &http.Client{}
	}
	if options.Service == "" {
		options.Service = MailTm
	}
	return &MailClient{http: options.Http, service: options.Service}, nil
}

func New() (*MailClient, error) {
	return &MailClient{service: MailTm, http: &http.Client{}}, nil
}
