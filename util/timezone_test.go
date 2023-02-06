package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSetGlobalTimezoneAsUTC(t *testing.T) {
	err := SetGlobalTimezoneAsUTC()
	require.NoError(t, err)

	location, err := time.LoadLocation("Local")
	require.NoError(t, err)

	assert.Equal(t, "UTC", location.String())
}