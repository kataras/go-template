package pug

import "gopkg.in/kataras/go-template.v0/html"

// Pug is the 'jade', same configs as the html engine

// Config for pug template engine
type Config html.Config

// DefaultConfig returns the default configuration for the pug(jade) template engine
func DefaultConfig() Config {
	return Config(html.DefaultConfig())
}
