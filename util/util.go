// Author: coolliu
// Date: 2021/7/27

package util

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func FileToYaml(file string, data interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("readfile {err:%s}", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		return fmt.Errorf("Unmarshal yaml {err:%s}", err.Error())
	}
	return nil
}

func FileToJson(file string, data interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("readfile {err:%s}", err.Error())
	}
	err = json.Unmarshal(yamlFile, data)
	if err != nil {
		return fmt.Errorf("Unmarshal json {err:%s}", err.Error())
	}
	return nil
}

func ToJson(data interface{}) string {
	d, _ := json.Marshal(data)
	return string(d)
}
