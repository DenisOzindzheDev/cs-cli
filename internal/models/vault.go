package models

type VaultSecret struct {
	HC_VAULT_ROLE_ID   string
	HC_VAULT_SECRET_ID string
}

// login request structures
type Body struct {
	Token string `json:"client_token"`
}
type Auth struct {
	Object Body `json:"auth"`
}
