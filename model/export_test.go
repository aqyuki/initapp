package model_test

import (
	"path/filepath"
	"testing"
)

func createTestPath(t *testing.T) string {
	t.Helper()
	return filepath.Join("test", "sample")
}
