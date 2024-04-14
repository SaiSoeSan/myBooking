package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/SaiSoeSan/bookings/pkg/config"
	"github.com/SaiSoeSan/bookings/pkg/model"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

func RenderTempalate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}

	//get template requested from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from Template Cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template {}

	//get all the files named *.page.html from ./template
	pages, err := filepath.Glob("./template/*.page.html")

	if err != nil {
		return myCache, err
	}

	//range through all files 
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0  {
			ts, err = ts.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
 	}

	return myCache, nil
}

// var tc = make(map[string] *template.Template)

// func RenderTempalate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have template in the cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create a template
// 		log.Println("creating template and add into cache")
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println("error")
// 		}
// 	} else {
// 		//we have template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w,nil)

// 	if err != nil {
// 		log.Println("error")
// 	}

// }

// func createTemplateCache(t string ) error {
// 	templates := []string{
// 		fmt.Sprintf("./template/%s", t),
// 		"./template/base.layout.html",
// 	}

// 	//parse the template
// 	tmpl , err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return 	err
// 	}

// 	//add template to cache
// 	tc[t] = tmpl
// 	return nil
// }