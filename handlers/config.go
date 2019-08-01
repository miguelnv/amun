package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Response - responses
type Response struct {
	Path        string
	Headers     map[string]string `yaml:"headers"`
	QueryParams map[string]string `yaml:"query"`
	ContentType string            `yaml:"contentType"`
	Template    string            `yaml:"template"`
	RawTemplate []byte            `yaml:"-"`
}

// Cfg - yaml file
type Cfg struct {
	Responses []Response
}

// ReadConfig read configuration file
func ReadConfig(cfgFilePath string) Cfg {
	y := Cfg{}

	filename, err := filepath.Abs(cfgFilePath)
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		log.Fatalf("yamlFile.Get unmarshal   #%v ", err)
	}

	log.Printf("yamlFile.struct   #%v ", y)

	y.applyRawTemplate()

	return y
}

func (cfg *Cfg) applyRawTemplate() {
	for i := range cfg.Responses {
		cfg.Responses[i].RawTemplate = []byte(cfg.Responses[i].Template)
	}
}

func (r *Response) ContainsHeaders(rHds *http.Header) bool {
	for k, v := range r.Headers {
		if rHds.Get(k) != v {
			return false
		}
	}
	return true
}

func (r *Response) ContainsParams(params url.Values) bool {
	for k, v := range r.QueryParams {
		if params.Get(k) != v {
			return false
		}
	}
	return true
}
