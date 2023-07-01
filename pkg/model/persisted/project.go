package persisted

type Project struct {
	HiddenFieldsModel
	Name    string
	Uuid    string
	ApiKeys []ApiKey
}
type ApiKey struct {
	HiddenFieldsModel
	ProjectID uint
	Project   Project
	Value     string
}
