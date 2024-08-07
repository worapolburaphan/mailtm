package mailtm

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Domain struct {
	ID        string    `json:"id"`
	TLD       string    `json:"domain"`
	IsActive  bool      `json:"isActive"`
	IsPrivate bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var cacheDomains []Domain

func (c *MailClient) GetDomains() ([]Domain, error) {
	var response []Domain

	if len(cacheDomains) > 0 {
		return cacheDomains, nil
	}

	req, err := http.NewRequest("GET", string(c.service)+"/domains", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	cacheDomains = response

	return response, nil
}

func (c *MailClient) GetDomainByID(id string) (*Domain, error) {
	var response Domain

	req, err := http.NewRequest("GET", string(c.service)+"/domains/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return &response, nil
}
