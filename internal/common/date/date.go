package date

import (
	"fmt"
	"strings"
)

func DisplayTime(seconds int64, granularity ...int) string {
	gran := 2
	if len(granularity) > 0 && granularity[0] > 0 {
		gran = granularity[0]
	}

	if seconds <= 0 {
		return "unknown"
	}

	intervals := []struct {
		Name  string
		Value int64
	}{
		{"week", 604800},
		{"day", 86400},
		{"hour", 3600},
		{"minute", 60},
		{"second", 1},
	}

	result := []string{}
	for _, interval := range intervals {
		if seconds >= interval.Value {
			count := seconds / interval.Value
			seconds %= interval.Value

			name := interval.Name
			if count > 1 {
				name += "s"
			}
			result = append(result, fmt.Sprintf("%d %s", count, name))

			if len(result) >= gran {
				break
			}
		}
	}

	return strings.Join(result, ", ")
}
