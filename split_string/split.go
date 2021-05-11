package split_string

import "strings"

func Split(str string, sep string) []string {
	var ret []string
	idx := strings.Index(str, sep)
	for idx > -1 {
		ret = append(ret, str[:idx])
		str = str[idx+len(sep):]
		idx = strings.Index(str, sep)
	}

	ret = append(ret, str)
	return ret
}
