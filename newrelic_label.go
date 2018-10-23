package goconfig

import (
	"errors"
	"fmt"
	"strings"
)

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

	return labels, nil
}
