package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s! Your env variable is %s", r.URL.Path[1:], os.Getenv("ENV_EXAMPLE"))
}

func main() {
	// os.Setenv("TEST_ENV", "TESTSTESTSET")
	var test_env = os.Getenv("ENV_EXAMPLE")
	http.HandleFunc("/", helloServer)
	fmt.Println("Starting server on port 8080")
	fmt.Println(test_env)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
