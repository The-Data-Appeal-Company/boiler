package conf

type Config struct {
	Source               SourceModel           `json:"source" yaml:"source"`
	Transformations      []TransformationModel `json:"transformations" yaml:"transformations"`
	RequestExecutorModel RequestExecutorModel  `json:"executor" yaml:"executor"`
}

type RequestExecutorBudgetModel struct {
	Time string `json:"time" yaml:"time"`
}

type RequestExecutorConfigurationModel struct {
	Concurrency     int  `json:"concurrency" yaml:"concurrency"`
	ContinueOnError bool `json:"continue_on_error" yaml:"continue_on_error"`
}

type RequestExecutorModel struct {
	Type   string                            `json:"type" yaml:"type"`
	Budget RequestExecutorBudgetModel        `json:"budget" yaml:"budget"`
	Config RequestExecutorConfigurationModel `json:"configuration" yaml:"configuration"`
	Params map[string]interface{}            `json:"params" yaml:"params"`
}

type SourceModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}

type TransformationModel struct {
	Type   string                 `json:"type" yaml:"type"`
	Params map[string]interface{} `json:"params" yaml:"params"`
}
