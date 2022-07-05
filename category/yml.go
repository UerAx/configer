/*
 * @Author: ww
 * @Date: 2022-07-03 18:06:03
 * @Description:
 * @FilePath: /goconf/category/yml.go
 */
package category

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadYaml(file string, obj interface{}) (interface{}, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}