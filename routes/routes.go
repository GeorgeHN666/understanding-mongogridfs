package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/GeorgeHN666/understanding-mongogridfs/handlers"
	"github.com/gorilla/mux"
)

func Routes() {

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.UploadFileEndPoint)
	router.HandleFunc("/getImage", handlers.ServeImage)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server Listening in PORT:::", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))

}
