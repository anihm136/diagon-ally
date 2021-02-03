package watcher

import (
	"diagon_ally/settings"
	"github.com/rjeczalik/notify"
	"fmt"
)

func WatchDir(c chan notify.EventInfo, path string) {
	go notify.Watch(path, c, notify.All)
}

func OnUpdate(c chan notify.EventInfo, conf *settings.Settings) {
	for {
		if ev := <-c; ev != nil {
			fmt.Println(ev.Event(), ev.Path())
		}
	}
}
