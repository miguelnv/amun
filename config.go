package main

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
	err = yaml.Unmarshal(yamlFile, &y)
	log.Printf("yamlFile.struct   #%v ", y)

	return y
}

func (response *Response) ContainsHeaders(rHeaders *http.Header) bool {

	for k, v := range response.Headers {
		if rHeaders.Get(k) != v {
			return false
		}
	}
	return true
}

func (response *Response) ContainsParams(params url.Values) bool {

	for k, v := range response.QueryParams {
		if params.Get(k) != v {
			return false
		}
	}
	return true
}
