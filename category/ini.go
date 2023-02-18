package category

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ReadIni(file string) (interface{}, error) {

	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	f := bufio.NewScanner(readFile)
	f.Split(bufio.ScanLines)
	
	hash := make(map[string]interface{})
	hash[""] = make(map[string]interface{})
	key := ""
	tmpHash := make(map[string]interface{})
	for f.Scan() {
		tmp := strings.TrimSpace(f.Text())
		if  len(tmp) > 0 && tmp[0] == ';' {
			continue;
		}
		judge := regexp.MustCompile(`^\[.*\]$`)
		extract := regexp.MustCompile(`\[(.*)\]`)
		if judge.MatchString(f.Text()) {
			ext := extract.FindStringSubmatch(f.Text())
			if len(ext) == 2 {
				key = ext[1]
				if _, ok := hash[key]; !ok {
					tmpHash = make(map[string]interface{})
					hash[key] = tmpHash
				}
			}
		} else {
			s := strings.Split(f.Text(), "=")
			if len(s) > 1 {
				tmpHash[strings.Trim(strings.TrimSpace(s[0]), "\"")] = strings.Trim(strings.TrimSpace(s[1]), "\"")
			}
		}

	}

	return hash, nil
}