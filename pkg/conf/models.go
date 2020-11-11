package conf

type Config struct {
	Source          SourceModel           `json:"source"`
	Transformations []TransformationModel `json:"transformations"`
}

type SourceModel struct {
	Type string
	Params map[string]interface{}
}

type TransformationModel struct {
	Type string
	Params map[string]interface{}
}
