package server

import (
	"fmt"
	"net/http"

	"dockercodecompiler/compiler"
)

func compile(w http.ResponseWriter, req *http.Request) {
	runOutput := compiler.Compile([]string{"", "python3", "print('hi)", "hello"})
	fmt.Fprintf(w, runOutput)
}

func Serve() {

	http.HandleFunc("/compile", compile)

	http.ListenAndServe(":8090", nil)
}
