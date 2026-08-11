package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Azure AKS GitOps!")
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Azure DevOps | Docker | Kubernetes | Helm | Argo CD")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/courses", coursesHandler)

	fmt.Println("Server started on port 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}