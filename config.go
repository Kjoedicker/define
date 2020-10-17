package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Website []struct {
		API struct {
			Website string `yaml:"website"`
			Link    string `yaml:"link"`
			Apikey  string `yaml:"apikey"`
		} `yaml:"api"`
	} `yaml:"website"`
	Dictionary string `yaml:"dictionary"`
}

func getConfig() (*config, string) {

	defPath, ok := os.LookupEnv("DEFINE_PATH")
	if !ok {
		panic("$DEFINE_PATH - enviromental variable not set")
	}

	buf, err := ioutil.ReadFile(defPath + "/conf.yaml")
	if err != nil {
		panic("conf.yaml - not in path")
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		panic("conf.yaml - invalid configuration")
	}

	return conf, defPath
}

func getDictConf(conf *config) string {
	return conf.Dictionary
}

func parseConfig(conf *config, idx int) (string, string, string) {

	return conf.Website[idx].API.Website, conf.Website[idx].API.Link, conf.Website[idx].API.Apikey
}