package configurations

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Dictionary struct {
	Dict         []enToFa `yaml:"dict"`
	PhoneNumbers []string `yaml:"phone_numbers"`
}

type DictionaryResult struct {
	Maps         map[string]string
	PhoneNumbers []string
}

func (c *Dictionary) GetConf() *DictionaryResult {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	resultmap := make(map[string]string)
	for _, element := range c.Dict {
		resultmap[element.En] = element.Fa
	}
	return &DictionaryResult{
		Maps:         resultmap,
		PhoneNumbers: c.PhoneNumbers,
	}
}

type enToFa struct {
	En string `yaml:"en"`
	Fa string `yaml:"fa"`
}
