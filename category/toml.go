/*
 * @Author: ww
 * @Date: 2022-07-04 01:20:48
 * @Description:
 * @FilePath: /goconf/category/toml.go
 */
package category

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func ReadToml(file string, obj interface{}) (interface{}, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}