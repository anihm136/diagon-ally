package server

import (
	"diagon_ally/settings"
	"diagon_ally/utils"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rjeczalik/notify"
)

func runCommand(command exec.Cmd) {
	err := command.Run()
	// log.Println("Executing command", command)
	if err != nil {
		log.Printf("Command %s exited with error: %v\n", command.Args[0], err)
	}
}

func OnUpdate(c chan notify.EventInfo, conf *settings.Settings) {
	for {
		// ev := <-c
		// log.Println(ev.Event())
		switch ev := <-c; ev.Event() {
		case notify.Write:
			fallthrough
		case notify.Create:
			go func() {
				path := ev.Path()
				args := utils.Replace(conf.OnUpdate, "${IN}", path)
				filename := filepath.Base(path)
				extension := filepath.Ext(filename)
				filenameWithoutExtension := strings.TrimSuffix(filename, extension)
				outPath := filepath.Join(conf.ExportDir, filenameWithoutExtension)
				// TODO: allow other formats
				outPath = outPath + ".png"
				args = utils.Replace(args, "${OUT}", outPath)
				cmd := exec.Command(args[0], args[1:]...)
				runCommand(*cmd)
			}()
		default:
		}
	}
}
