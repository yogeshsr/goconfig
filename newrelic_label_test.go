package goconfig

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseEmptyLabel(t *testing.T) {
	labels, err := parseNewRelicLabels("")

	assert.Nil(t, labels)
	assert.EqualError(t, err, "unable to parse newrelic labels")
}

func TestParseValidLabelParts(t *testing.T) {
	labels, err := parseNewRelicLabels("team:team1;runtime:go;response_time:250ms")

	assert.Equal(t, 3, len(labels))
	assert.Equal(t, "team1", labels["team"])
	assert.Equal(t, "250ms", labels["response_time"])
	assert.NoError(t, err)
}

func TestParseValidLabelPartsWithoutResponseTime(t *testing.T) {
	labels, err := parseNewRelicLabels("foo:bar;key1:value1")

	assert.Equal(t, 2, len(labels))
	assert.Equal(t, "bar", labels["foo"])
	assert.Equal(t, "value1", labels["key1"])
	assert.NoError(t, err)
}

func TestInvalidLabelString(t *testing.T) {
	_, err := parseNewRelicLabels("foo:bar;go:o:car")
	assert.Error(t, err)
}

func TestInvalidResponseTimeLabelShouldPanic(t *testing.T) {
	assert.Panics(t, func() {parseNewRelicLabels("foo:bar;response_time:11ms")}, "should panic")
}




