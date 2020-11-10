package transformation

import (
	"boiler/pkg/requests"
	"fmt"
	"strings"
	"time"
)

type RelativeDateShiftConfiguration struct {
	TargetFields []string
	RelativeTo   string
	DateFormat   string
}

type RelativeDateShiftTransformation struct {
	conf RelativeDateShiftConfiguration
}

func NewDateShift() {

}

func (r RelativeDateShiftTransformation) Apply(request requests.Request) (requests.Request, error) {
	date, err := valueOrResolveParam(r.conf.RelativeTo, request)
	if err != nil {
		return requests.Request{}, err
	}

	relative, err := time.Parse(r.conf.DateFormat, date)
	if err != nil {
		return requests.Request{}, err
	}

	timeDelta := truncDay(time.Now()).Sub(truncDay(relative))


}

func truncDay(val time.Time) time.Time {
	return time.Date(val.Year(), val.Month(), val.Day(), 0, 0, 0, 0, val.Location())
}

func valueOrResolveParam(key string, req requests.Request) (string, error) {
	if isVariable(key) {
		return getString(req.SourceParams, key)
	}
	return key, nil
}

func getString(values map[string]interface{}, key string) (string, error) {
	rawUrl, present := values[key]
	if !present {
		return "", fmt.Errorf("variable not present in source params: %s", key)
	}

	rawUrlStr, isStr := rawUrl.([]byte)
	if !isStr {
		return "", fmt.Errorf("value for source param %s must be a string", key)
	}

	return string(rawUrlStr), nil
}

func isVariable(val string) bool {
	const varPrefix = "$"
	return strings.HasPrefix(val, varPrefix)
}
