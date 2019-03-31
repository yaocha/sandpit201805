package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type hoge struct {
	Hoge []string `yaml:"hoge"`
}

func main() {
	configYAML, err := ioutil.ReadFile("hoge.yaml")
	if err != nil {
		log.Printf("err: %v", err)
	}

	h := &hoge{}
	err = yaml.Unmarshal(configYAML, &h)
	if err != nil {
		log.Printf("err: %v", err)
	}

	log.Print(h.Hoge)

	ioutil.WriteFile("test.txt", configYAML, os.ModePerm)
}
