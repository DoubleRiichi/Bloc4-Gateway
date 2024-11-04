package main

import (
	"fmt"
	"net/http"
	"os"

	bloc4_config "github.com/DoubleRiichi/BLOC4-Gateway/internal/config"
)

func main() {
	config, err := bloc4_config.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bloc_server := bloc4_config.ConfigIntoServer(config)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

	})
	bloc_server.ListenAndServe()

	fmt.Printf("%+v\n", config)
}
