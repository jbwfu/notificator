package main

import (
	"log"

	"github.com/go-musicfox/notificator"
)

var notify *notificator.Notificator

func main() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "My test App",
		OSXSender:   "com.netease.163music",
	})

	notify.Push(notificator.UrNormal, "title", "text", "/home/user/icon.png", "https://github.com/go-musicfox/go-musicfox")

	// Check errors
	err := notify.Push(notificator.UrCritical, "error", "ops =(", "/home/user/icon.png", "https://github.com/go-musicfox/go-musicfox")

	if err != nil {
		log.Fatal(err)
	}
}
