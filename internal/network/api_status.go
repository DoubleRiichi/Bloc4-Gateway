package bloc4_network

import (
	"fmt"
	"net"
	"time"

	bloc4_config "github.com/DoubleRiichi/BLOC4-Gateway/internal/config"
)

type status int8

const (
	DOWN status = iota
	UP
	REJECT
)

type ApiState struct {
	Api    bloc4_config.ApiCnf
	Status status
}

func PingAPI(api bloc4_config.ApiCnf) ApiState {
	host_port := api.Host + ":" + api.Port

	timeout := time.Duration(1 * time.Second)

	//add support for more protocols eventually?
	con, err := net.DialTimeout("tcp", host_port, timeout)

	if err != nil {
		return ApiState{api, DOWN}
	}

	con.Close()
	return ApiState{api, UP}
}

func GetAPIsStatus(config bloc4_config.Config) ([]ApiState, error) {

	if len(config.Apis) == 0 {
		return nil, fmt.Errorf("no API configuration found")
	}

	var states []ApiState

	for _, api_config := range config.Apis {
		states = append(states, PingAPI(api_config))
	}

	return states, nil
}
