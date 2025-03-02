//go:build darwin

package notificator

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type osxNotificator struct {
	AppName string
	Sender  string
}

func New(o Options) *Notificator {
	return &Notificator{
		notifier:    osxNotificator{AppName: o.AppName, Sender: o.OSXSender},
		defaultIcon: o.DefaultIcon,
	}
}

func (o osxNotificator) push(title string, text string, iconPath string, redirectUrl string) error {
	// Checks if terminal-notifier exists, and is accessible.

	// if terminal-notifier exists, use it.
	// else, fall back to osascript. (Mavericks and later.)
	if CheckTermNotif() {
		if redirectUrl != "" {
			return exec.Command("terminal-notifier", "-title", o.AppName, "-message", text, "-subtitle", title, "-contentImage", iconPath, "-open", redirectUrl).Run()
		}
		return exec.Command("terminal-notifier", "-title", o.AppName, "-message", text, "-subtitle", title, "-contentImage", iconPath, "-sender", o.Sender).Run()
	} else if CheckMacOSVersion() {
		title = strings.ReplaceAll(title, `"`, `\"`)
		text = strings.ReplaceAll(text, `"`, `\"`)

		notification := fmt.Sprintf("display notification \"%s\" with title \"%s\" subtitle \"%s\"", text, o.AppName, title)
		return exec.Command("osascript", "-e", notification).Run()
	}

	// finally falls back to growlnotify.

	return exec.Command("growlnotify", "-n", o.AppName, "--image", iconPath, "-m", title, "--url", redirectUrl).Run()
}

// Causes the notification to stick around until clicked.
func (o osxNotificator) pushCritical(title string, text string, iconPath string, redirectUrl string) error {
	// same function as above...
	if CheckTermNotif() {
		// timeout set to 30 seconds, to show the importance of the notification
		if redirectUrl != "" {
			return exec.Command("terminal-notifier", "-title", o.AppName, "-message", text, "-subtitle", title, "-contentImage", iconPath, "-timeout", "30", "-open", redirectUrl).Run()
		}

		return exec.Command("terminal-notifier", "-title", o.AppName, "-message", text, "-subtitle", title, "-contentImage", iconPath, "-timeout", "30", "-sender", o.Sender).Run()
	} else if CheckMacOSVersion() {
		notification := fmt.Sprintf("display notification \"%s\" with title \"%s\" subtitle \"%s\"", text, o.AppName, title)
		return exec.Command("osascript", "-e", notification).Run()
	}

	return exec.Command("growlnotify", "-n", o.AppName, "--image", iconPath, "-m", title, "--url", redirectUrl).Run()
}

// Helper function for macOS

func CheckTermNotif() bool {
	// Checks if terminal-notifier exists, and is accessible.
	if err := exec.Command("which", "terminal-notifier"); err != nil {
		return false
	}
	// no error, so return true. (terminal-notifier exists)
	return true
}

func CheckMacOSVersion() bool {
	// Checks if the version of macOS is 10.9 or Higher (osascript support for notifications.)

	cmd := exec.Command("sw_vers", "-productVersion")
	check, _ := cmd.Output()

	version := strings.Split(strings.TrimSpace(string(check)), ".")

	// semantic versioning of macOS

	if len(version) < 2 {
		return false
	}

	major, _ := strconv.Atoi(version[0])
	minor, _ := strconv.Atoi(version[1])

	if major < 10 {
		return false
	} else if major == 10 && minor < 9 {
		return false
	} else {
		return true
	}
}
