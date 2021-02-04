package server

import (
	"github.com/rjeczalik/notify"
)

func WatchDir(c chan notify.EventInfo, path string) {
	go notify.Watch(path, c, notify.All)
}
