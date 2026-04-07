package notification

type NotificationEvent string

const (
	StateChange  NotificationEvent = "state-change"
	Restart      NotificationEvent = "restart"
	Start        NotificationEvent = "start"
	Stop         NotificationEvent = "stop"
	Backup       NotificationEvent = "backup"
	ConfigChange NotificationEvent = "config-change"
	Install      NotificationEvent = "install"
	Uninstall    NotificationEvent = "uninstall"
	Upgrade      NotificationEvent = "upgrade"
)

type Type string

const (
	NotificationType Type = "notification"
	EventType        Type = "event"
)
