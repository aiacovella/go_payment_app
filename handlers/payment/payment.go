package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"gotutorial/models"
	"net/http"
	"os"
)

func HandleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	var req models.CreatePaymentRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductID)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(paymentIntent.ClientSecret)

	var response struct {
		ClientSecret string `json:"client_secret"`
	}

	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	if _, err := writer.Write(buf.Bytes()); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func calculateOrderAmount(productID string) int64 {

	switch productID {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}

	return 0
}
