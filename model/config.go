package model

type (
	// Config is the struct to set config of Initializer instance
	Config struct {
		Type FileType // Type is a type of configure file type.
		Path string   // Path is a path of save configure file
	}
)

// NewConfig create config instance
func NewConfig(ft FileType, path string) *Config {
	var cnf *Config = new(Config)

	// Check supported file type.
	if ft == FileTypeJSON {
		cnf.Type = FileTypeJSON
	} else if ft == FileTypeYAML {
		cnf.Type = FileTypeYAML
	} else {
		// Panic exit if a file format ID not supported by the library is specified.
		panic("Unexpected file format ID specified")
	}

	// Check path already set
	if path == "" {
		panic("Path not specified")
	}
	cnf.Path = path

	return cnf
}
