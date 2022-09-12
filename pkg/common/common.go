package common

type Client struct {
	PK           string `json:"PK"`
	SK           string `json:"SK"`
	SMSLimit     string `json:"SMSLimit"`
	BusinessName string `json:"BusinessName"`
}

type ClientData struct {
	BusinessID   string `json:"BusinessID"`
	Email        string `json:"Email"`
	IsAdmin      bool   `json:"IsAdmin"`
	Enabled      bool   `json:"Enabled"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	SMSLimit     int    `json:"SMSLimit"`
	BusinessName string `json:"BusinessName"`
}
