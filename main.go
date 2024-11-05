package main

import (
	"fmt"
	"os"

	bloc4_config "github.com/DoubleRiichi/BLOC4-Gateway/internal/config"
	bloc4_network "github.com/DoubleRiichi/BLOC4-Gateway/internal/network"
)

func main() {
	config, err := bloc4_config.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	states, err := bloc4_network.GetAPIsStatus(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := bloc4_config.ConfigIntoServer(config)

	for _, state := range states {
		fmt.Printf("%v : %v\n", state.Api.Host, state.Status)

		if state.Status == bloc4_network.UP {
			bloc4_network.ConstructSubdomainRoute(state.Api)
		}
	}

	fmt.Printf("%+v\n", config)

	server.ListenAndServe()
	fmt.Printf("Listening on " + server.Addr + "...")

}
