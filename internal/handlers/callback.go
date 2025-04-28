package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Callback received:", string(body))
	w.WriteHeader(http.StatusOK)
}
