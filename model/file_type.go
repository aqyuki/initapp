package model

type (
	// FileType is used to identify file types
	FileType int
)

const (
	FileTypeJSON FileType = 1 // FileTypeJSON is a json file
	FileTypeYAML FileType = 2 // FileTypeYAML is a yaml file
)
