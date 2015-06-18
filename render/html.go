package render

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	tpl = template.New("")
)

func ParseHTMLTemplates(root string, funcMap template.FuncMap) {

	tpl.Funcs(funcMap)

	filepath.Walk(root, func(path string, fi os.FileInfo, _ error) error {

		if fi == nil || fi.IsDir() {

			return nil
		}

		if name, err := filepath.Rel(root, path); err == nil {

			if content, err := ioutil.ReadFile(path); err == nil {

				template.Must(tpl.New(name).Parse(string(content)))

			} else {

				panic(err)
			}

		} else {

			panic(err)
		}

		return nil
	})
}

func HTML(rw http.ResponseWriter, name string, data interface{}) error {

	return tpl.ExecuteTemplate(rw, name, data)
}
