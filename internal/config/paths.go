package bloc4_config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	gopath "path"
	"path/filepath"
)

const DEFAULT_CONFIG_PATH = "/config/"
const SERVER_CONFIG_PATH = DEFAULT_CONFIG_PATH + "server.yml"
const APIS_CONFIG_PATH = "/config/app"

func getWorkingDirectory() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		//should be replaced by a logger function when implemented
		return path, fmt.Errorf("error Trying to get current Working Directory :\n %w", err)
	}

	return path, nil
}

func checkFileExists(path string) bool {
	_, error := os.Stat(path)
	return !errors.Is(error, os.ErrNotExist)
}

func isAPIconfigFile(name string) bool {
	result, _ := filepath.Match("api_*.yaml", name)
	return result
}

func isServerConfigFile(name string) bool {
	result, _ := filepath.Match("server.yaml", name)
	return result
}

func loadFile(path string) (*[]byte, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return &file, fmt.Errorf("could not read file %v :\n %w", path, err)

	}

	return &file, nil
}

// search for the relevant folders and construct a dictionary containing .yaml files to be loaded
func SearchConfigFiles(path string) ([]configPair, error) {

	var found_configs []configPair

	executable_path, err := os.Executable()

	if err != nil {
		fmt.Println(err.Error())
	}

	executable_dir := gopath.Dir(executable_path)

	if !checkFileExists(executable_dir + DEFAULT_CONFIG_PATH) {
		return found_configs, errors.New("config path " + executable_dir + DEFAULT_CONFIG_PATH + "not found.")
	}

	filepath.Walk(executable_dir, func(path string, file fs.FileInfo, err error) error {
		if err != nil {

			return fmt.Errorf("error locating config files... %v :\n %w", path, err)
		}

		if isAPIconfigFile(file.Name()) {
			found_configs = append(found_configs, configPair{kind: cAPI, path: path})
		}

		if isServerConfigFile(file.Name()) {
			found_configs = append(found_configs, configPair{kind: cSERVER, path: path})
		}

		return nil
	})

	return found_configs, nil
}
