package initapp

import (
	"errors"
	"os"

	"github.com/aqyuki/initapp/files"
	"github.com/aqyuki/initapp/model"
)

func decode[T interface{}](ft model.FileType, b []byte) (*T, error) {
	var data *T
	var err error
	switch ft {
	case model.FileTypeJSON:
		data, err = decodeJSON[T](b)
	case model.FileTypeYAML:
		data, err = decodeYAML[T](b)
	}
	if err != nil {
		return nil, err
	}
	return data, err
}

func encode[T interface{}](ft model.FileType, object *T) ([]byte, error) {
	var data []byte
	var err error
	switch ft {
	case model.FileTypeJSON:
		data, err = encodeJSON[T](object)
	case model.FileTypeYAML:
		data, err = encodeYAML[T](object)
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Initialize[T interface{}](cv *model.Converter) (*T, error) {
	raw, err := files.LoadFile(cv.Config.Path)
	if err != nil {
		return nil, err
	}
	return decode[T](cv.Config.Type, raw)
}

func InitializeWithDefault[T interface{}](cv *model.Converter, normal T) (*T, error) {
	raw, err := files.LoadFile(cv.Config.Path)
	if err == nil {
		return decode[T](cv.Config.Type, raw)
	}
	if errors.Is(err, os.ErrNotExist) {
		b, err := encode[T](cv.Config.Type, &normal)
		if err != nil {
			return nil, err
		}
		files.WriteFile(cv.Config.Path, b)
		return &normal, nil
	}
	return nil, err
}
