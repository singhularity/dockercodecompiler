package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockercodecompiler/compiler"
)

type CompilerParams struct {
	Language string
	Code     string
	Stdin    string
}

func compile(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var compilerParams CompilerParams
		err := json.NewDecoder(req.Body).Decode(&compilerParams)
		if err == nil {
			runParams := append([]string{""}, compilerParams.Language, compilerParams.Code, compilerParams.Stdin)
			fmt.Print(runParams)
			runOutput := compiler.Compile(runParams)
			w.Write([]byte(runOutput))
		} else {
			w.Write([]byte("Failed to read compiler parameter"))
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported Method"))
	}
}

// Serve Start webserver
func Serve() {

	http.HandleFunc("/api/compile", compile)

	http.ListenAndServe(":8090", nil)
}
