package django

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"fmt"

	"github.com/flosch/pongo2"
)

type (
	// Engine the pongo2 engine
	Engine struct {
		Config        Config
		templateCache map[string]*pongo2.Template
		mu            sync.Mutex
	}
)

const (
	templateErrorMessage = `
	<html><body>
	<h2>Error in template</h2>
	<h3>%s</h3>
	</body></html>`
)

// New creates and returns a Pongo template engine
func New(cfg ...Config) *Engine {
	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	// cuz mergo has a little bug on maps
	if c.Globals == nil {
		c.Globals = make(map[string]interface{}, 0)
	}
	if c.Filters == nil {
		c.Filters = make(map[string]pongo2.FilterFunction, 0)
	}

	return &Engine{Config: c, templateCache: make(map[string]*pongo2.Template)}
}

// Funcs should returns the helper funcs
func (p *Engine) Funcs() map[string]interface{} {
	return p.Config.Globals
}

// LoadDirectory builds the templates
func (p *Engine) LoadDirectory(dir string, extension string) (templateErr error) {
	fsLoader, err := pongo2.NewLocalFileSystemLoader(dir) // I see that this doesn't read the content if already parsed, so do it manually via filepath.Walk
	if err != nil {
		return err
	}

	set := pongo2.NewSet("", fsLoader)
	set.Globals = getPongoContext(p.Config.Globals)
	// Walk the supplied directory and compile any files that match our extension list.
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// Fix same-extension-dirs bug: some dir might be named to: "users.tmpl", "local.html".
		// These dirs should be excluded as they are not valid golang templates, but files under
		// them should be treat as normal.
		// If is a dir, return immediately (dir is not a valid golang template).
		if info == nil || info.IsDir() {

		} else {

			rel, err := filepath.Rel(dir, path)
			if err != nil {
				templateErr = err
				return err
			}

			ext := filepath.Ext(rel)
			if ext == extension {

				buf, err := ioutil.ReadFile(path)
				if err != nil {
					templateErr = err
					return err
				}
				if err != nil {
					templateErr = err
					return err
				}
				name := filepath.ToSlash(rel)
				p.templateCache[name], templateErr = set.FromString(string(buf))

				if templateErr != nil && p.Config.DebugTemplates {
					p.templateCache[name], _ = set.FromString(
						fmt.Sprintf(templateErrorMessage, templateErr.Error()))
				}
			}

		}
		return nil
	})

	return
}

// LoadAssets loads the templates by binary
func (p *Engine) LoadAssets(virtualDirectory string, virtualExtension string, assetFn func(name string) ([]byte, error), namesFn func() []string) error {
	var templateErr error
	/*fsLoader, err := pongo2.NewLocalFileSystemLoader(virtualDirectory)
	if err != nil {
		return err
	}*/
	set := pongo2.NewSet("", pongo2.DefaultLoader) //I'm not 100% sure that works on assets but we will seee
	set.Globals = getPongoContext(p.Config.Globals)

	if len(virtualDirectory) > 0 {
		if virtualDirectory[0] == '.' { // first check for .wrong
			virtualDirectory = virtualDirectory[1:]
		}
		if virtualDirectory[0] == '/' || virtualDirectory[0] == os.PathSeparator { // second check for /something, (or ./something if we had dot on 0 it will be removed
			virtualDirectory = virtualDirectory[1:]
		}
	}

	names := namesFn()
	for _, path := range names {
		if !strings.HasPrefix(path, virtualDirectory) {
			continue
		}

		rel, err := filepath.Rel(virtualDirectory, path)
		if err != nil {
			templateErr = err
			return err
		}

		ext := filepath.Ext(rel)
		if ext == virtualExtension {

			buf, err := assetFn(path)
			if err != nil {
				templateErr = err
				return err
			}
			name := filepath.ToSlash(rel)
			p.templateCache[name], err = set.FromString(string(buf))
			if err != nil {
				templateErr = err
				return err
			}
		}
	}
	return templateErr
}

// getPongoContext returns the pongo2.Context from map[string]interface{} or from pongo2.Context, used internaly
func getPongoContext(templateData interface{}) pongo2.Context {
	if templateData == nil {
		return nil
	}

	if contextData, isPongoContext := templateData.(pongo2.Context); isPongoContext {
		return contextData
	}

	return templateData.(map[string]interface{})
}

func (p *Engine) fromCache(relativeName string) *pongo2.Template {
	p.mu.Lock() // defer is slow

	tmpl, ok := p.templateCache[relativeName]

	if ok {
		p.mu.Unlock()
		return tmpl
	}
	p.mu.Unlock()
	return nil
}

// ExecuteWriter executes a templates and write its results to the out writer
// layout here is useless
func (p *Engine) ExecuteWriter(out io.Writer, name string, binding interface{}, options ...map[string]interface{}) error {
	if tmpl := p.fromCache(name); tmpl != nil {
		return tmpl.ExecuteWriter(getPongoContext(binding), out)
	}

	return fmt.Errorf("[IRIS TEMPLATES] Template with name %s doesn't exists in the dir", name)
}
