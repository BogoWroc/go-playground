package api

type CreateLabReq struct {
	CreatedBy    string
	StackName    string
	TemplateName string
	EnvName      string
	InitLab      bool
}

type LabModel struct {
	StackId   string
	StackName string
	Ip        string
	Status    string
	CreatedBy string
	CreatedAt string
}

type LabManagerApi interface {
	CreateLab(req CreateLabReq) (bool, error)
	FetchAllLabs() ([]LabModel, error)
	FetchLabsFor(createdBy string) ([]LabModel, error)
	TerminateLab(stackId string) (bool, error)
}
