package handler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestConfigurationError(t *testing.T) {
	buffer := CaptureLogs(t)

	builder, err := Build(nil)
	require.Nil(t, builder)
	require.Error(t, err)

	assert.Contains(t, "AWS configuration cant be nil", err.Error())
	assert.Contains(t, "AWS configuration cant be nil", buffer.String())
}

/*********** Helpers *********/

func CaptureLogs(t *testing.T) *bytes.Buffer {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	t.Cleanup(func() {
		log.SetOutput(os.Stdout)
	})

	return &buffer
}
