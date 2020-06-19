package yamlreader

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func NewYamlReader(path string) *yamlReader {
	return &yamlReader{
		path: path,
	}
}

type yamlReader struct {
	path string
}

func (reader *yamlReader) GetIntValue(key string, defaultValue int) (int, bool) {
	value, ok := reader.load()[key]
	if ok == false {
		return defaultValue, false
	}

	intValue, ok := value.(int)
	if ok == false {
		intValue = defaultValue
	}
	return intValue, ok
}

func (reader *yamlReader) GetBoolValue(key string, defaultValue bool) (bool, bool) {
	value, ok := reader.load()[key]
	if ok == false {
		return defaultValue, false
	}

	boolValue, ok := value.(bool)
	if ok == false {
		boolValue = defaultValue
	}
	return boolValue, ok
}

func (reader *yamlReader) load() map[string]interface{} {
	var result = map[string]interface{}{}

	content, err := ioutil.ReadFile(reader.path)

	if err != nil {
		log.Printf("Error reading %s: %v\n", reader.path, err)
	} else {
		err = yaml.Unmarshal(content, result)

		if err != nil {
			log.Printf("Error parsing yaml %s: %v\n", reader.path, err)
		}
	}

	return result
}
