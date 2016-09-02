package django

import "github.com/flosch/pongo2"

// Config for django template engine
type Config struct {
	// Filters for pongo2, map[name of the filter] the filter function . The filters are auto register
	Filters map[string]pongo2.FilterFunction
	// Globals share context fields between templates. https://github.com/flosch/pongo2/issues/35
	Globals map[string]interface{}
	// DebugTemplates enables template debugging.
	// The verbose error messages will appear in browser instead of quiet passes with error code
	DebugTemplates bool
}

// DefaultConfig returns the default configuration for the django template engine
func DefaultConfig() Config {
	return Config{
		Filters:        make(map[string]pongo2.FilterFunction),
		Globals:        make(map[string]interface{}, 0),
		DebugTemplates: false,
	}
}
