package render

import (
	"bytes"
	"github.com/OusManDiouf/go-booking-app/pkg/config"
	"github.com/OusManDiouf/go-booking-app/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	pathToTemplateFiles = "./templates/"
)

var functions = template.FuncMap{}

// appConfig will hold an instance of appConfig for this package
var appConfig *config.AppConfig

// NewTemplates Get the appConfig from global configuration to this package
func NewTemplates(ac *config.AppConfig) {
	appConfig = ac
}
func AddDefaultData(tmplData *models.TemplateData) *models.TemplateData {

	// No default data for now, just foward the given data
	return tmplData
}

// Template render template using html
func Template(writer http.ResponseWriter, tmpl string, tmplData *models.TemplateData) {

	var templateCache map[string]*template.Template

	if appConfig.UseCache {
		// Get the template cache from the appConfig
		templateCache = appConfig.TemplateCache
	} else {
		// rebuild the cache on each request (onDevMode)
		templateCache, _ = CreateTemplateCache()
	}

	parsedTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
		return
	}
	buffer := new(bytes.Buffer)

	tmplData = AddDefaultData(tmplData) // default + passed in data
	_ = parsedTemplate.Execute(buffer, tmplData)
	_, errBuffer := buffer.WriteTo(writer)
	if errBuffer != nil {
		log.Println(errBuffer)
	}

}

// CreateTemplateCache create a cached templated in a Map DS
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// Get all pages
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		log.Println(err)
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, errTs := template.New(name).Funcs(functions).ParseFiles(page)
		if errTs != nil {
			log.Println(errTs)
			return templateCache, errTs
		}

		matches, errMatching := filepath.Glob(pathToTemplateFiles + "*.page.gohtml")
		if errMatching != nil {
			log.Println(errMatching)
			return templateCache, errMatching
		}

		if len(matches) > 0 {
			templateSet, errParseGlob := templateSet.ParseGlob(pathToTemplateFiles + "*.layout.gohtml")
			if errParseGlob != nil {
				log.Println(errParseGlob)
				return templateCache, errParseGlob
			}

			templateCache[name] = templateSet
		}
	}
	return templateCache, nil
}
