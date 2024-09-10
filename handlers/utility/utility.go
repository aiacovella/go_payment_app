package utility

import (
	"log"
	"net/http"
)

func HandleHealth(writer http.ResponseWriter, request *http.Request) {

	var response []byte = []byte("{'status': 'ok'}")

	if _, err := writer.Write(response); err != nil {
		log.Fatalln("Error writing response:")
	}
	
}
