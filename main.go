package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/myip", MyIP)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, This is GO server!")
	fmt.Fprintf(w, "\n===========================================================================\n")
}

func MyIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my IP!")
	fmt.Fprintf(w, "\n===========================================================================\n")

	prg := "ip"
	arg1 := "addr"
	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintf(w, string(stdout))
}
