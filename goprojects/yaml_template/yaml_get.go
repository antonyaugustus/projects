package main

import (
	"gopkg.in/yaml.v2"
	//"errors"
	"fmt"
	"io/ioutil"
)

// Test Kitchen instance configuration
type Config struct {
	Index    string
	NumShard string
	Xmx      string
	Shard1   []string
}

func main() {
	var config Config
	source, err := ioutil.ReadFile("projects/goprojects/yaml_template/kmds.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Shard1[0])
}
