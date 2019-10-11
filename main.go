package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	log.Println("serving port http://localhost:9000")

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	out := `
	<html>
		<head>
			<title>Service Deployment</title>
			<style>
				%s
			</style>
		</head>
		<body style="background-color: %s">
			<div class="content">
				<h1>%s</h1>	
			</div>
		</body>
	</html>
	`
	color, message := getDeployment()
	fmt.Fprintln(w, fmt.Sprintf(out, style, color, message))
}

func getDeployment() (string, string) {
	deployment := os.Getenv("DEPLOYMENT")
	switch deployment {
	case "green":
		return "#008000", "🌻<br />Green Deployment!"
	case "blue":
		return "#00BFFF", "🌤<br />Blue Deployment!"
	default:
		return "white", "🦓<br />Black and White Deployment!"
	}
}

var style = `
body {
	margin: auto;
	padding: 25px;
	text-align: center;
}

.content {
    margin: auto;
    padding-top: 5%;
}

.content h1 {
	font-size: 7rem;
}
`
