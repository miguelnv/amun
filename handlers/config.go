package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Mapping - mapping
type Mapping struct {
	Path        string
	Headers     map[string]string `yaml:"headers"`
	QueryParams map[string]string `yaml:"query"`
	ContentType string            `yaml:"contentType"`
	Template    string            `yaml:"template"`
}

// Cfg - yaml file
type Cfg struct {
	Mappings []Mapping
}

// ReadConfig read configuration file
func ReadConfig(cfgFilePath string) Cfg {

	filename, err := filepath.Abs(cfgFilePath)
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	cfg := Cfg{}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalf("yamlFile.Get unmarshal   #%v ", err)
	}

	log.Printf("yamlFile.struct   #%v ", cfg)

	return cfg
}

func (r *Mapping) MatchHeaders(rHds *http.Header) bool {
	for k, v := range r.Headers {
		if rHds.Get(k) != v {
			return false
		}
	}
	return true
}

func (r *Mapping) MatchParams(params url.Values) bool {
	for k, v := range r.QueryParams {
		if params.Get(k) != v {
			return false
		}
	}
	return true
}
