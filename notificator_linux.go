//go:build linux

package notificator

import (
	"os/exec"
)

type linuxNotificator struct {
	AppName string
}

func New(o Options) *Notificator {
	var Notifier notifier

	Notifier = linuxNotificator{AppName: o.AppName}

	return &Notificator{notifier: Notifier, defaultIcon: o.DefaultIcon}
}

func (l linuxNotificator) push(title string, text string, iconPath string, redirectUrl string) error {
	return exec.Command("notify-send", "-a", l.AppName, "-i", iconPath, title, text).Run()
}

// Causes the notification to stick around until clicked.
func (l linuxNotificator) pushCritical(title string, text string, iconPath string, redirectUrl string) error {
	return exec.Command("notify-send", "-a", l.AppName, "-i", iconPath, title, text, "-u", "critical").Run()
}
