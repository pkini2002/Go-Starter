package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template Cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}

	// Get requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Error: Template not found in cache:", tmpl)
	}

	buf := new(bytes.Buffer)
	
	err = t.Execute(buf, nil) 

	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error){
	myCache := map[string]*template.Template{}

	// get all the files name *.page.gohtml from the templates directory
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.go.html
	for _, page := range pages {
		name := filepath.Base(page) // get the file name
		ts, err := template.New(name).ParseFiles(page) // parse the file

		if err != nil {
			return myCache, err
		}

		// Identify the base layout file
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts // add the template to the cache
	}
	return myCache, nil
}


