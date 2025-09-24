package notification

type Event string

const (
	StateChange  Event = "state-change"
	Restart      Event = "restart"
	Backup       Event = "backup"
	ConfigChange Event = "config-change"
	Install      Event = "install"
	Uninstall    Event = "uninstall"
	Upgrade      Event = "upgrade"
)
