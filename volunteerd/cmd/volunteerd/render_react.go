package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

// renderReactApp renders a react application.
func renderReactApp(manifestPath, path, entrypoint string) (template.HTML, template.CSS, error) {
	resolvedPath, err := resolveManifest(manifestPath, path)
	if err != nil {
		return "", "", err
	}
	vm := goja.New()

	// enable goja extensions
	registry := new(require.Registry)
	registry.Enable(vm)
	console.Enable(vm)

	script, err := ioutil.ReadFile(resolvedPath)
	if err != nil {
		return "", "", err
	}
	_, err = vm.RunScript("index.js", string(script))
	if err != nil {
		return "", "", err
	}

	renderFn := vm.Get("renderServerSide")
	_, ok := goja.AssertFunction(renderFn)
	if !ok {
		return "", "", fmt.Errorf("did not find renderServerSide in bundle")
	}

	var fn func() ([]string, error)
	err = vm.ExportTo(renderFn, &fn)
	if err != nil {
		return "", "", err
	}
	res, err := fn()
	if len(res) == 0 {
		return "", "", fmt.Errorf("got zero values from render function - it should return an array of [html, css]")
	} else if len(res) == 1 {
		return template.HTML(res[0]), "", err
	} else {
		return template.HTML(res[0]), template.CSS(res[1]), err
	}
}

func resolveManifest(manifestPath, path string) (string, error) {
	manifest := map[string]string{}
	f, err := os.Open(manifestPath)
	if err != nil {
		return "", err
	}
	if err := json.NewDecoder(f).Decode(&manifest); err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(manifestPath), manifest[path]), nil
}
