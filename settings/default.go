package settings

import (
	"github.com/adrg/xdg"
	"github.com/knadh/stuffbin"
	"log"
	"os"
)

var MAX_TEMPLATE_SIZE int = 100000000
var MAX_INSERT_SIZE int = 100000000

func newSettings() *Settings {
	var conf *Settings = new(Settings)
	conf.WatchDir = "./svg"
	conf.ExportDir = "./images"
	conf.NewTemplate, conf.InsertTemplate = getDefaultTemplates()
	conf.OnUpdate = []string{"inkscape", "--export-area-drawing", "--export-area-snap","--export-type=png", "-o", "${OUT}", "${IN}"}
	return conf
}

func getDefaultTemplates() (string, string) {
	path, err := os.Executable()
	if err != nil {
		log.Fatalf("error getting executable path: %v", err)
	}
	fs, err := stuffbin.UnStuff(path)
	if err != nil {
		log.Fatalf("error reading stuffed binary: %v", err)
	}

	var template, insert string
	for _, val := range []string{"default_template.svg", "default_insert.txt"} {
		f, err := fs.Get(val)
		if err != nil {
			log.Fatalf("error reading %s: %v", val, err)
		}
		template = insert
		insert = string(f.ReadBytes())
	}

	return template, insert
}

func getConfigPath() string {
	settingsPath, _ := xdg.ConfigFile("diagon_ally/settings.txt")
	return settingsPath
}
