// Get the access token from safaricom
package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stkpush/internal/config"
)

func GetAccessToken() (string, error) {
	consumerKey := config.GetEnv("CONSUMER_KEY")
	consumerSecret := config.GetEnv("CONSUMER_SECRET")

	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	credentials := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	req.Header.Add("Authorization", "Basic "+credentials)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("could not get access token")
	}
	return token, nil
}
