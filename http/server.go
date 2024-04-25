package http

import (
	"fmt"
	"net/http"
)

// Port we listen on.
const (
	HttpPort string = ":8080"
	InfoPath        = "/info"
)


func Home(w http.ResponseWriter, r *http.Request) { // Handler function.
	fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) { // Handler function.
	fmt.Fprintf(w, "Info page")
}

//func main() {
//	log.Println("Starting our simple http server.")
//
//	// Registering our handler functions, and creating paths.
//	http.HandleFunc("/", Home)
//	http.HandleFunc(InfoPath, Info)
//
//	log.Println("Started on port", HttpPort)
//	fmt.Println("To close connection CTRL+C :-)")
//
//	// Spinning up the server.
//	err := http.ListenAndServe(HttpPort, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
