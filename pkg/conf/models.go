package conf

type Config struct {
	Source               SourceModel           `json:"source" yaml:"source"`
	Transformations      []TransformationModel `json:"transformations" yaml:"transformations"`
	RequestExecutorModel RequestExecutorModel  `json:"executor" yaml:"executor"`
}

type RequestExecutorModel struct {
	Type            string                 `json:"type" yaml:"type"`
	Concurrency     int                    `json:"concurrency"`
	ContinueOnError bool                   `json:"continue_on_error"`
	Params          map[string]interface{} `json:"params" yaml:"params"`
}

type SourceModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}

type TransformationModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}
