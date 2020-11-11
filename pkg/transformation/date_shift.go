package transformation

import (
	"boiler/pkg/requests"
	"fmt"
	"strings"
	"time"
)

const (
	TransformationRelativeDateShift = "relative-time-shift"
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
	relative, err := r.parseOrResolveParamTime(r.conf.RelativeTo, r.conf.DateFormat, request)
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

func (r RelativeDateShiftTransformation) parseOrResolveParamTime(key string, layout string, req requests.Request) (time.Time, error) {
	if isVariable(key) {
		return r.getTime(req.SourceParams, key[1:])
	}

	return time.Parse(layout, key)
}

func (r RelativeDateShiftTransformation) getTime(values map[string]interface{}, key string) (time.Time, error) {
	rawUrl, present := values[key]
	if !present {
		return time.Time{}, fmt.Errorf("variable not present in source params: %s", key)
	}

	tim, isTime := rawUrl.(time.Time)
	if isTime {
		return tim, nil
	}

	str, isStr := rawUrl.(string)
	if isStr {
		return time.Parse(r.conf.DateFormat, str)
	}

	return time.Time{}, fmt.Errorf("value for source param %s must be a time.Time or string compatible with %s layout", key, r.conf.DateFormat)
}

func isVariable(val string) bool {
	const varPrefix = "$"
	return strings.HasPrefix(val, varPrefix)
}
