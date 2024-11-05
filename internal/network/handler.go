package bloc4_network

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	bloc4_config "github.com/DoubleRiichi/BLOC4-Gateway/internal/config"
)

// Wrap a given api config in order to access its values within the general Handler
type ApiHandler struct {
	config bloc4_config.ApiCnf
}

func (api *ApiHandler) simpleHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	fmt.Printf("Received request %v : %v \n\n %v\n", r.URL.RequestURI(), r.Header, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//this might not be needed actually lmao
	url_right := strings.Split(r.URL.RequestURI(), "/"+api.config.GatewayName+"/")[1]
	fmt.Println(r.URL.RequestURI())
	// exemple where api.gatewayName = test, api.host = randomapi.net and api.port = 443 and request is 127.0.0.1/test/movies/get
	// url_right = movies/get
	// url = randomapi.net:443/movies/get
	// WARNING: Added "http://" because of "first path segment in URL cannot contain colon" error
	url := api.config.Protocol + "://" + api.config.Host + ":" + api.config.Port + "/" + url_right

	forwardedReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	forwardedReq.Header = r.Header.Clone()

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	// proxyReq.Header = make(http.Header)
	// for h, val := range req.Header {
	//     proxyReq.Header[h] = val
	// }
	client := http.Client{}
	fmt.Printf("Forwarding request %v \n", forwardedReq.URL.RequestURI())

	resp, err := client.Do(forwardedReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)

		return
	}

	//TODO: should go into its own function
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	resp_body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp_body)

	defer resp.Body.Close()
}
