package settings

import (
	"bufio"
	"diagon_ally/utils"
	"errors"
	"fmt"
	"log"
	"os"
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
		if utils.Contains(
			[]string{
				"WatchDir",
				"ExportDir",
				"OnUpdate",
			},
			parts[0]) == -1 {
			err = errors.New("Invalid key in config file")
			return
		}
		userConf[parts[0]] = parts[1]
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
	if val, ok := userConf["OnUpdate"]; ok {
		args := strings.Split(val, " ")
		conf.OnUpdate = args
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

func (settings *Settings) UpdateFlags(source string, dest string, onUpdate string) {
	if source != "" {
		settings.WatchDir = source
	}
	if dest != "" {
		settings.ExportDir = dest
	}
	if onUpdate != "" {
		settings.OnUpdate = strings.Split(onUpdate, " ")
	}
}
