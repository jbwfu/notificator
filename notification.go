package notificator

type Options struct {
	DefaultIcon string
	AppName     string
	OSXSender   string
}

const (
	UrNormal   = "normal"
	UrCritical = "critical"
)

type notifier interface {
	push(title string, text string, iconPath string, redirectUrl string) error
	pushCritical(title string, text string, iconPath string, redirectUrl string) error
}

type Notificator struct {
	notifier    notifier
	defaultIcon string
}

func (n Notificator) Push(urgency string, title string, text string, iconPath string, redirectUrl string) error {
	if n.notifier == nil {
		return nil
	}
	icon := n.defaultIcon

	if iconPath != "" {
		icon = iconPath
	}

	if urgency == UrCritical {
		return n.notifier.pushCritical(title, text, icon, redirectUrl)
	}

	return n.notifier.push(title, text, icon, redirectUrl)
}
