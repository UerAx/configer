/*
 * @Author: ww
 * @Date: 2022-07-03 17:35:13
 * @Description:
 * @FilePath: /goconf/conf.go
 */
package goconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/uerax/goconf/category"
)

type CfgFile struct {
	//mx 		sync.Mutex
	Path	string
	File	string
	Data	map[string]interface{}
	tmp		map[string]interface{}
}

func (t *CfgFile) New() {
	t = &CfgFile{Data: make(map[string]interface{}, 0), tmp: make(map[string]interface{}, 0)}
}

/*
* read all configuration file on this path
* @param path string 文件目录
*/
func (t *CfgFile) ReadFiles(path string) error {
	files, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("获取文件信息出错 : %v", err)
	}

	if !files.IsDir() {
		return t.ReadConfig(path)
	}

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("文件目录打开出错 : %v", err)
	}
	
	for _, file := range dir {
		p := path + "/" + file.Name()
		t.ReadFiles(p)
	}
	t.Data = t.tmp
	t.Path = path
	return nil
}

/*
* if you just have one configuration file, that you can use this method
* @param file string 配置文件绝对路径或者相对路径
*/
func (t *CfgFile) ReadConfig(file string) error {
	
	f, err := os.Stat(file)
	
	if err != nil || f.IsDir() {
		return fmt.Errorf("配置文件或者文件夹不存在 : %v", err)
	}
	
	ext := strings.ToLower(path.Ext(file))
	
	switch ext {
	case ".yaml":
		data, err := category.ReadYaml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.tmp[f.Name()] = data

	case ".yml":
		data, err := category.ReadYaml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.tmp[f.Name()] = data
	case ".json":
		data, err := category.ReadJson(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.tmp[f.Name()] = data
	case ".toml":
		data, err := category.ReadToml(file, struct{}{})
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.tmp[f.Name()] = data
	case ".ini":
		data, err := category.ReadIni(file)
		if err != nil {
			return fmt.Errorf("配置文件读取出错 %v", err)
		}
		t.tmp[f.Name()] = data
	default:
		return fmt.Errorf("配置文件类型[%s]不支持,目前仅支持yaml,yml,json,toml格式", f.Name())
	}

	t.File = file
	return nil
}

/**
* @param section string "section:如果没有可以不填"
* @param key	string	"key:必填的参数key"
* @return value interface{} "value: 自行转换成 string|slice|map 等类型"
*/
func (t *CfgFile) GetValue(in... string) interface{} {
	if len(in) < 1 {
		return nil
	}
	tmp := t.Data
	// 首层为文件名
	for _, data := range tmp {
		// 空文件跳过
		if _, ok := data.(map[string]interface{}); !ok {
			continue
		}
		tmp = data.(map[string]interface{})
		// 开始遍历该文件是否有key对应的值
		for i := 0; i < len(in); i++ {
			if v, ok := tmp[in[i]]; !ok {
				break
			} else {
				// 无法转换证明已经没有下一层
				if _, ok := v.(map[string]interface{}); !ok {
					// 非最后一个key但是已经没有下一层证明不在该文件里
					if i + 1 != len(in) {
						break
					}
					return v
				}
				tmp = v.(map[string]interface{})
			}
		}
	}
	

	return nil
}

/*
* @param filename string 文件名
* @param obj interface{} 转换的结构体
*/
func (t *CfgFile) Unmarshal4Name(name string, obj interface{}) error {
	if v, ok := t.Data[name]; ok {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%s 文件不存在", name) 
}

func (t *CfgFile) Reload() error {
	if len(t.Path) != 0 {
		t.ReadFiles(t.Path)
	} else if len(t.File) != 0 {
		t.ReadConfig(t.File)
	} else {
		return fmt.Errorf("未读取过配置文件无法重载")
	}
	return nil
}
