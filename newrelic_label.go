package goconfig

import (
	"errors"
	"fmt"
	"strings"
)

var allowedRespTimes = map[string]bool{
	"15ms":    true,
	"30ms":    true,
	"60ms":    true,
	"125ms":   true,
	"250ms":   true,
	"500ms":   true,
	"1000ms":  true,
	"2000ms":  true,
	"5000ms":  true,
	"10000ms": true,
	"30000ms": true,
	"60000ms": true,
}

func parseNewRelicLabels(labelStr string) (map[string]string, error) {
	labels := map[string]string{}
	for _, labelParts := range strings.Split(labelStr, ";") {
		if labelParts == "" {
			return nil, errors.New("unable to parse newrelic labels")
		}
		parts := strings.Split(labelParts, ":")
		if len(parts) != 2 {
			return nil, errors.New(fmt.Sprintf("Invalid newrelic labels configuration: %s", labelParts))
		}
		labels[parts[0]] = parts[1]
	}

	if labels["response_time"] != "" && !allowedRespTimes[labels["response_time"]] {
		panic(errors.New(fmt.Sprintf("invalid new relic label for response_time. allowed values are %v", allowedRespTimes)))
	}
	return labels, nil
}
