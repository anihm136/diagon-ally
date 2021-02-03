package settings

import (
	"bufio"
	"diagon_ally/utils"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseConfig(configPath string) (userConf map[string]string, err error) {
	userConf = make(map[string]string)
	configFile, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("User config file not found, using defaults")
			err = nil
		} else {
			return
		}
	}
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if utils.Contains([]string{"WatchDir", "ExportDir", "NewTemplateFile", "InsertTemplateFile"}, parts[0]) == -1 {
			err = errors.New("Invalid key in config file")
			return
		}
		userConf[parts[0]], err = filepath.Abs(parts[1])
		if err != nil {
			return
		}
	}
	return
}

func mergeUserSettings(conf *Settings, userConf map[string]string) *Settings {
	// TODO: This is a very ugly way of doing this
	if val, ok := userConf["WatchDir"]; ok {
		conf.WatchDir = val
	}
	if val, ok := userConf["ExportDir"]; ok {
		conf.ExportDir = val
	}
	if val, ok := userConf["NewTemplateFile"]; ok {
		f, err := os.Open(val)
		if err != nil {
			log.Printf("NewTemplateFile (%s) does not exist, using default\n", val)
		} else {
			readBytes := make([]byte, MAX_TEMPLATE_SIZE)
			_, err = f.Read(readBytes)
			// TODO: Handle errors
			conf.NewTemplate = string(readBytes)
		}
	}
	if val, ok := userConf["InsertTemplateFile"]; ok {
		f, err := os.Open(val)
		if err != nil {
			log.Printf("InsertTemplateFile (%s) does not exist, using default\n", val)
		} else {
			readBytes := make([]byte, MAX_INSERT_SIZE)
			_, err = f.Read(readBytes)
			// TODO: Handle errors
			conf.InsertTemplate = string(readBytes)
		}
	}

	return conf
}

func GetSettings() (conf *Settings, err error) {
	conf = newSettings()
	configPath := getConfigPath()
	userConf, err := parseConfig(configPath)
	if err != nil {
		return
	}
	conf = mergeUserSettings(conf, userConf)
	err = setupPaths(conf, false)
	return
}
