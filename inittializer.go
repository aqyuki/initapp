package initapp

import (
	"github.com/aqyuki/initapp/files"
	"github.com/aqyuki/initapp/model"
)

func Initialize[T interface{}](cv *model.Converter) (*T, error) {
	raw, err := files.LoadFile(cv.Config.Path)
	if err != nil {
		return nil, err
	}

	var data *T
	switch cv.Config.Type {
	case model.FileTypeJSON:
		data, err = decodeJSON[T](raw)
	case model.FileTypeYAML:
		data, err = decodeYAML[T](raw)
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
