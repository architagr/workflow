package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct{
	Ip string `json:"Ip"`
	Port string `jsom:"Port"`
	DataBaseConnection string `json:"DataBaseConnection"`
}

func (e *Configuration) Init(environment string) (Configuration, error) {
	var err1 error
	fileName := "config."+environment+".json"

	var data Configuration
	file, err := ioutil.ReadFile(fileName)
	if err != nil{
		err1 = fmt.Errorf("error reading config - %s", fileName)
	}
	
	unmarshalError := json.Unmarshal(file, &data)
	if unmarshalError != nil {
		err1 = fmt.Errorf("unable to decode json file - %s", fileName)
	}
	return data, err1
}