package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var SysConfs conf

type conf struct {
	ServerHost  string `yaml:"serverHost"`
	SqlAddr     string `yaml:"sqlAddr"`
	SqlUser     string `yaml:"sqlUser"`
	SqlDatabase string `yaml:"sqlDatabase"`
}

func InitConfig() error {
	//ymlFile, err := ioutil.ReadFile("config.yml")
	ymlFile, err := ioutil.ReadFile("/Users/lijunbo/go/src/github.com/JunboLi-Ray/go-demo/config/config.yml")
	if err != nil {
		log.Printf("ymlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymlFile, &SysConfs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}
