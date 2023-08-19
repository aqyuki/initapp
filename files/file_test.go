package files_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aqyuki/initapp/files"
)

func createTempFile(t *testing.T, path string, raw []byte) {
	f, err := os.Create(path)
	if err != nil {
		t.Errorf("could not create a file for testing")
	}
	defer f.Close()

	_, err = f.Write(raw)
	if err != nil {
		t.Errorf("could not write file to use testing")
	}

}

func TestLoadFile(t *testing.T) {

	target := "sample"
	raw := []byte(target)
	path := filepath.Join(t.TempDir(), "sample.txt")

	createTempFile(t, path, raw)

	data, err := files.LoadFile(path)
	if err != nil {
		t.Errorf("error returned, got -> %+v", err)
	}
	if !reflect.DeepEqual(data, raw) {
		t.Errorf("loaded data is different from expected, want -> %+v, got -> %+v", string(data), string(raw))
	}

	_, err = files.LoadFile(filepath.Join("path", "to", "your", "file.txt"))
	if err == nil {
		t.Errorf("did not return error")
	}
}
