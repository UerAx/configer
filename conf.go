/*
 * @Author: ww
 * @Date: 2022-07-03 17:35:13
 * @Description:
 * @FilePath: /goconf/conf.go
 */
package goconf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/uerax/goconf/category"
)

type CfgFile struct {
	Files	[]string
	// fixme 由于是map[string]interface{}格式 struct暂无法支持
	Data	map[string]interface{}
	//fixme 后期后改为双类型的data struct和map分开
	isStruct bool
}

func NewCfgFile() *CfgFile {
	return &CfgFile{make([]string, 0), make(map[string]interface{}, 0), false}
}

func (t *CfgFile) New() *CfgFile {
	return &CfgFile{make([]string, 0), make(map[string]interface{}, 0), false}
}

func (t *CfgFile) ReadAll(path string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("文件目录打开出错 : %v", err)
	}
	
	for _, file := range dir {
		p := path + "/" + file.Name()
		if file.IsDir() {
			t.ReadAll(p)
		} else {
			t.ReadConfig(p)
		}
		
	}
}

func (t *CfgFile) ReadConfig(file string, obj... interface{}) {
	
	f, err := os.Stat(file)
	
	if err != nil || f.IsDir() {
		log.Printf("配置文件或者文件夹不存在 : %v", err)
	}
	
	ext := strings.ToLower(path.Ext(file))

	if len(obj) == 0 {
		obj = append(obj, make(map[string]interface{})) 
	} else {
		t.isStruct = true
	}
	
	switch ext {
	case ".yaml":
		data, err := category.ReadYaml(file, obj[0])
		if err != nil {
			log.Printf("配置文件读取出错 %v", err)
		}
		t.Data[f.Name()] = data

	case ".yml":
		data, err := category.ReadYaml(file, obj[0])
		if err != nil {
			log.Printf("配置文件读取出错 %v", err)
		}
		t.Data[f.Name()] = data
	case ".json":
		data, err := category.ReadJson(file, obj[0])
		if err != nil {
			log.Printf("配置文件读取出错 %v", err)
		}
		t.Data[f.Name()] = data
	case ".toml":
		data, err := category.ReadToml(file, obj[0])
		if err != nil {
			log.Printf("配置文件读取出错 %v", err)
		}
		t.Data[f.Name()] = data
	default:
		log.Printf("配置文件类型[%s]不支持,目前仅支持yaml,yml,json,toml格式", ext)
	}
	
}

/**
* @param section string "section:如果没有可以不填"
* @param key	string	"key:必填的参数key"
* @return value interface{} "value: 自行转换成 string|slice|map 等类型"
*/
func (t *CfgFile) GetValue(in... string) (interface{}, error) {
	if t.isStruct {
		return nil, fmt.Errorf("获取结构体请使用 GetStruct 方法")
	}
	if len(in) < 1 || len(in) > 2 {
		return nil, fmt.Errorf("参数有误,只能是 (section, key) 或则 (key) ")
	}
	if len(in) == 1 {
		for _, d := range t.Data {
			if v, ok := d.(map[string]interface{})[in[0]]; ok {
				return v, nil
			}
		}
	}
	if len(in) == 2 {
		for _, d := range t.Data {
			if v, ok := d.(map[string]interface{})[in[0]]; ok {
				if v1, ok := v.(map[string]interface{})[in[1]]; ok {
					return v1, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("没有找到对应的配置")
}

func (t *CfgFile) GetStruct() []interface{} {
	rsl := make([]interface{}, 0, len(t.Data))
	for _, v := range t.Data {
		rsl = append(rsl, v)
	}
	return rsl
}