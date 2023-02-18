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
	
	hash := make(map[string]map[string]string)
	hash[""] = make(map[string]string)
	key := ""
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
					hash[key] = make(map[string]string)
				}
			}
		} else {
			s := strings.Split(f.Text(), "=")
			if len(s) > 1 {
				hash[key][strings.Trim(strings.TrimSpace(s[0]), "\"")] = strings.Trim(strings.TrimSpace(s[1]), "\"")
			}
		}

	}

	return hash, nil
}