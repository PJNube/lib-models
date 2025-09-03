package datatypes

type NotificationEvent string

const (
	NotificationEvent_StateChange  NotificationEvent = "state-change"
	NotificationEvent_Restart      NotificationEvent = "restart"
	NotificationEvent_Backup       NotificationEvent = "backup"
	NotificationEvent_ConfigChange NotificationEvent = "config-change"
	NotificationEvent_Install      NotificationEvent = "install"
	NotificationEvent_Uninstall    NotificationEvent = "uninstall"
	NotificationEvent_Upgrade      NotificationEvent = "upgrade"
)
