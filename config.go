package main

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Website    string `yaml:"website"`
	Link       string `yaml:"link"`
	APIKey     string `yaml:"apikey"`
	Dictionary string `yaml:"dictionary"`
}

func getConfig() (string, string, string, string, string, error) {

	defPath, ok := os.LookupEnv("DEFINE_PATH")
	if !ok {
		return "", "", "", "", "", errors.New("$DEFINE_PATH - enviromental variable not set")
	}

	buf, err := ioutil.ReadFile(defPath + "/conf.yaml")
	if err != nil {
		return "", "", "", "", "", errors.New("conf.yaml - not in path")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return "", "", "", "", "", errors.New("conf.yaml - invalid configuration")
	}
	return conf.Website, conf.Link, conf.APIKey, conf.Dictionary, defPath, nil
}