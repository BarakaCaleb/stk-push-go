package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"stkpush/internal/auth"
	"stkpush/internal/stkpush"
)

type PaymentRequest struct {
	PhoneNumber string `json:"phone_number"`
	Amount      int    `json:"amount"`
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var paymentReq PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get access token
	token, err := auth.GetAccessToken()
	if err != nil {
		http.Error(w, "Failed to get access token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Initiate STK Push
	err = stkpush.InitiateSTKPush(token, paymentReq.PhoneNumber, paymentReq.Amount)
	if err != nil {
		http.Error(w, "Failed to initiate STK Push: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "STK Push initiated successfully")
}
