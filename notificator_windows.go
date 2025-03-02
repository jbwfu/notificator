//go:build windows

package notificator

import "os/exec"

type windowsNotificator struct{}

func New(o Options) *Notificator {
	var Notifier notifier

	Notifier = windowsNotificator{}

	return &Notificator{notifier: Notifier, defaultIcon: o.DefaultIcon}
}

func (w windowsNotificator) push(title string, text string, iconPath string, redirectUrl string) *exec.Cmd {
	return exec.Command("growlnotify", "/i:", iconPath, "/t:", title, text)
}

// Causes the notification to stick around until clicked.
func (w windowsNotificator) pushCritical(title string, text string, iconPath string, redirectUrl string) *exec.Cmd {
	return exec.Command("growlnotify", "/i:", iconPath, "/t:", title, text, "/s", "true", "/p", "2")
}
