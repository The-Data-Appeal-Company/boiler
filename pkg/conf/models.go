package conf

type Config struct {
	Source               SourceModel           `json:"source" yaml:"source"`
	Transformations      []TransformationModel `json:"transformations" yaml:"transformations"`
	RequestExecutorModel RequestExecutorModel  `json:"request_executor_model" yaml:"request_executor_model"`
}

type RequestExecutorModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}

type SourceModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}

type TransformationModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}
