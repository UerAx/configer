/*
 * @Author: ww
 * @Date: 2022-07-04 01:18:38
 * @Description:
 * @FilePath: /goconf/category/json.go
 */
package category

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJson(file string, obj interface{}) (interface{}, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}