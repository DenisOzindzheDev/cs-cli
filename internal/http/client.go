package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Nested struct {
	Token string `json:"client_token"`
}
type Auth struct {
	Object Nested `json:"auth"`
}

func login(roleid, secretid, vaultnamespace string) string {
	url := "https://vault.int.rolfcorp.ru/v1/auth/approle/login"
	method := "POST"
	fmt.Printf("Calling login with %s %s", roleid, secretid)
	payload := strings.NewReader(fmt.Sprintf(`{
   "role_id":  "%s",
   "secret_id": "%s"
}`, roleid, secretid))

	fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

	}

	req.Header.Add("X-Vault-Namespace", vaultnamespace)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}

	var reqBody Auth
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		fmt.Println(err)
	}
	token := reqBody.Object.Token
	// fmt.Printf("\n\n\n\n\n %s\n\n\n\n\n\n\n", string(token))

	return token
}

func ExtractData(role, secret, path, vaultnamespace string) {
	token := login(role, secret, vaultnamespace)
	// fmt.Printf("\n\n\n\n\n%s\n\n\n\n", token)
	url := fmt.Sprintf("https://vault.int.rolfcorp.ru/v1/%s", path)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)

	}
	req.Header.Set("X-Vault-Namespace", vaultnamespace)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Vault-Token", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	fmt.Println(string(body))

}