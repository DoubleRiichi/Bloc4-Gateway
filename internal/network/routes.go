package bloc4_network

import (
	"net/http"

	bloc4_config "github.com/DoubleRiichi/BLOC4-Gateway/internal/config"
)

func ConstructSubdomainRoute(api bloc4_config.ApiCnf) {

	//TODO: assign specifics handlers depending on API configuration
	api_wrapper := ApiHandler{config: api}
	http.HandleFunc("/"+api.GatewayName+"/", api_wrapper.simpleHandler)
}
