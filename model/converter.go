package model

type (
	// Converter initialize application when application stated running
	Converter struct {
		Target interface{} // Target is a structures to parse some files
		Config *Config     // Config is a configuration of Initializer
	}
)

// NewConverter create Converter instance
func NewConverter(target interface{}, config *Config) *Converter {
	if target == nil {
		panic("target structure is nil")
	}
	if config == nil {
		panic("config is nil")
	}
	return &Converter{Config: config, Target: target}
}
