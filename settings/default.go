package settings

import (
	"github.com/adrg/xdg"
)

var (
	MAX_TEMPLATE_SIZE int = 100000000
	MAX_INSERT_SIZE   int = 100000000
)

func newSettings() *Settings {
	var conf *Settings = new(Settings)
	conf.WatchDir = "./svg"
	conf.ExportDir = "./images"
	conf.OnUpdate = []string{"inkscape", "--export-area-drawing", "--export-area-snap", "--export-type=png", "-o", "${OUT}", "${IN}"}
	return conf
}

func getConfigPath() string {
	settingsPath, _ := xdg.ConfigFile("diagon_ally/settings.txt")
	return settingsPath
}
