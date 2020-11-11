package transformation

import (
	"boiler/pkg/requests"
	"fmt"
	"strings"
	"time"
)


const (
	TransformationRelativeDateShift = "RelativeDateShiftTransformation"
)
type RelativeDateShiftConfiguration struct {
	TargetFields   []string
	RelativeTo     string
	DateFormat     string
	RelativeTimeFn func() time.Time
}

type RelativeDateShiftTransformation struct {
	conf RelativeDateShiftConfiguration
}

func NewDateShift(conf RelativeDateShiftConfiguration) RelativeDateShiftTransformation {
	return RelativeDateShiftTransformation{conf: conf}
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

	relativeTime := truncDay(r.conf.RelativeTimeFn())
	timeDelta := relativeTime.Sub(truncDay(relative))

	for _, paramName := range r.conf.TargetFields {
		value, present := request.Params[paramName]
		if !present {
			return requests.Request{}, fmt.Errorf("no query parameter with name %s found in request %s", paramName, request.String())
		}

		if len(value) == 0 {
			return requests.Request{}, fmt.Errorf("query parameter with name %s has more than 1 value", paramName)
		}

		parsedParam, err := time.Parse(r.conf.DateFormat, value[0])
		if len(value) == 0 {
			return requests.Request{}, fmt.Errorf("query parameter with name %s has more than 1 value", paramName)
		}

		if err != nil {
			return requests.Request{}, err
		}

		parsedParam = parsedParam.Add(timeDelta)

		formattedParam := parsedParam.Format(r.conf.DateFormat)
		request.Params[paramName] = []string{formattedParam}
	}

	return request, nil
}

func truncDay(val time.Time) time.Time {
	return time.Date(val.Year(), val.Month(), val.Day(), 0, 0, 0, 0, val.Location())
}

func valueOrResolveParam(key string, req requests.Request) (string, error) {
	if isVariable(key) {
		return getString(req.SourceParams, key[1:])
	}
	return key, nil
}

func getString(values map[string]interface{}, key string) (string, error) {
	rawUrl, present := values[key]
	if !present {
		return "", fmt.Errorf("variable not present in source params: %s", key)
	}

	rawUrlStr, isStr := rawUrl.(string)
	if !isStr {
		return "", fmt.Errorf("value for source param %s must be a string", key)
	}

	return rawUrlStr, nil
}

func isVariable(val string) bool {
	const varPrefix = "$"
	return strings.HasPrefix(val, varPrefix)
}
