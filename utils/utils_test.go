package utils

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSearchBrowserLauncher(t *testing.T) {
	browser := searchBrowserLauncher("darwin")
	assert.Equal(t, "open", browser)

	browser = searchBrowserLauncher("windows")
	assert.Equal(t, "cmd /c start", browser)
}
