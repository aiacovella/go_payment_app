package main

import (
	"gotutorial/handlers/payment"
	"gotutorial/handlers/utility"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.

	// Add handlers
	http.HandleFunc("/payments/create-payment-intent", payment.HandleCreatePaymentIntent)
	http.HandleFunc("/health/", utility.HandleHealth)

	// Start the web service
	log.Println("Starting web service...")
	err := http.ListenAndServe("localhost:4242", nil)

	if err != nil {
		log.Fatal(err)
	}

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
