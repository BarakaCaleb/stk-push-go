package stkpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stkpush/internal/config"
	"time"
)

func InitiateSTKPush(accessToken string) error {
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

	shortcode := config.GetEnv("BUSINESS_SHORTCODE")
	passkey := config.GetEnv("PASSKEY")
	phoneNumber := "254795430038" // Replace or pass dynamically
	amount := 1
	callbackURL := config.GetEnv("CALLBACK_URL")

	timestamp := time.Now().Format("20060102150405") // yyyyMMddHHmmss
	password := base64.StdEncoding.EncodeToString([]byte(shortcode + passkey + timestamp))

	payload := map[string]interface{}{
		"BusinessShortCode": shortcode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            shortcode,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       callbackURL,
		"AccountReference":  "TestPayment",
		"TransactionDesc":   "Payment for goods",
	}

	payloadBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, ioutil.NopCloser(bytes.NewReader(payloadBytes)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("STK Push Response:", string(body))

	return nil
}
