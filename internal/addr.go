package internal

import (
	"fmt"
	"regexp"
)

func GetAddressesList(r string) (result []string, err error) {
	addrRegex := regexp.MustCompile(`([\d]{1,3}).([\d]{1,3}).([\d]{1,3}).([\d]{1,3})/([\d]{1,2})`)
	strings := addrRegex.FindStringSubmatch(r)

	if strings[5] == "24" { // 192.168.x.YYY/24
		result = make([]string, 256)

		for index := range result {
			result[index] = fmt.Sprintf("%s.%s.%s.%d", strings[1], strings[2], strings[3], index)
		}

		return result, nil
	}

	if strings[5] == "16" { // 172.x.y.ZZZ/16
		result = make([]string, 256*256)

		for index := range result {
			result[index] = fmt.Sprintf("%s.%s.%d.%d", strings[1], strings[2], index/(256), index%256)
		}

		return result, err
	}

	err = fmt.Errorf("unexpected address type: %s", strings[5])

	return result, err
}
