package payment

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	json_request := json.NewDecoder(request.Body).Decode(&CreatePaymentRequest)

	if request.Method != http.MethodPost {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	fmt.Println("Payment Intent Created")

}
