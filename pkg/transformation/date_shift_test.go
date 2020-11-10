package transformation

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDateShiftRelativeToSameDate(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$request_date",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-01")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"2020-01-02"},
			"to":   {"2020-01-05"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	modifiedReq, err := shifter.Apply(req)
	require.NoError(t, err)

	require.Equal(t, modifiedReq.Params["from"][0], "2020-01-02")
	require.Equal(t, modifiedReq.Params["to"][0], "2020-01-05")
}

func TestDateShiftRelativeToNextDay(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$request_date",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-02")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"2020-01-02"},
			"to":   {"2020-01-05"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	modifiedReq, err := shifter.Apply(req)
	require.NoError(t, err)

	require.Equal(t, modifiedReq.Params["from"][0], "2020-01-03")
	require.Equal(t, modifiedReq.Params["to"][0], "2020-01-06")
}


func TestDateShiftRelativeAfter5Days(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$request_date",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-06")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"2020-01-02"},
			"to":   {"2020-01-05"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	modifiedReq, err := shifter.Apply(req)
	require.NoError(t, err)

	require.Equal(t, modifiedReq.Params["from"][0], "2020-01-07")
	require.Equal(t, modifiedReq.Params["to"][0], "2020-01-10")
}

func TestDateShiftRelativeToNotpresent(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$not_present",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-01")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"2020-01-02"},
			"to":   {"2020-01-05"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	_, err := shifter.Apply(req)
	require.Error(t, err)
}

func TestDateShiftRelativeToQueryParamNotPresent(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$request_date",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-01")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"2020-01-02"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	_, err := shifter.Apply(req)
	require.Error(t, err)
}

func TestDateShiftRelativeToQueryParamInvalidDate(t *testing.T) {
	const format = "2006-01-02"

	shifter := NewDateShift(RelativeDateShiftConfiguration{
		TargetFields: []string{"to", "from"},
		RelativeTo:   "$request_date",
		DateFormat:   format,
		RelativeTimeFn: func() time.Time {
			ts, err := time.Parse(format, "2020-01-01")
			require.NoError(t, err)
			return ts
		},
	})

	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"from": {"monday"},
			"to": {"2020-01-02"},
		},
		SourceParams: map[string]interface{}{
			"request_date": "2020-01-01",
		},
	}

	_, err := shifter.Apply(req)
	require.Error(t, err)
}