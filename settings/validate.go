package settings

import (
	"fmt"
	"os"
	// "errors"
)

func setupPaths(conf *Settings, force bool) error {
	// TODO: Find a better way of doing this
	var op string
	notExistMessage := "%s (%s) does not exist. Create? [Yn] "
	settings := map[string]string{
		"WatchDir":  conf.WatchDir,
		"ExportDir": conf.ExportDir,
	}

	var err error
	for k, v := range settings {
		_, err := os.Stat(v)
		if os.IsNotExist(err) {
			shouldCreate := true
			if !force {
				fmt.Printf(notExistMessage, k, v)
				fmt.Scanln(&op)
				switch op {
				case "Y":
				case "y":
				case "":
					shouldCreate = true
				default:
					shouldCreate = false
				}
			}
			if shouldCreate {
				err = os.Mkdir(v, 0o755)
				if err != nil {
					break
				}
			} else {
				return fmt.Errorf("%s (%s) does not exist, exiting", k, v)
			}
		} else {
			break
		}
	}
	return err
}
