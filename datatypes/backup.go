package datatypes

type BackupComponent string

const (
	ControlApp BackupComponent = "control-app"
	Data       BackupComponent = "data"
	All        BackupComponent = "all"
)

var BackupComponentMap = map[string]struct{}{
	string(ControlApp): {},
	string(Data):       {},
	string(All):        {},
}

type CreateStatus string

const (
	Completed  CreateStatus = "completed"
	InProgress CreateStatus = "in-progress"
	Failed     CreateStatus = "failed"
)

var CreateStatusMap = map[CreateStatus]struct{}{
	Completed:  {},
	InProgress: {},
	Failed:     {},
}
