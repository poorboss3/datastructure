package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func alertName(keyName []string, keyTime []string) []string {
	nameMap := make(map[string][]string)
	if len(keyName) != len(keyTime) {
		panic("key name len not cap the time")
	}
	for i, name := range keyName {
		if _, ok := nameMap[name]; !ok {
			nameMap[name] = make([]string, 0)
		}
		nameMap[name] = append(nameMap[name], keyTime[i])
	}
	exist := make([]string, 0)
	for name, times := range nameMap {
		if len(times) < 3 {
			continue
		}
		sort.Strings(times)
		for i := 0; i < len(times)-2; i++ {
			if minuteSub(times[i], times[i+2]) > -1 && minuteSub(times[i], times[i+2]) <= 60 {
				exist = append(exist, name)
				break
			}
		}
	}
	sort.Strings(exist)
	return exist
}

func minuteSub(start string, end string) int {
	smin := strings.Split(start, ":")
	emin := strings.Split(end, ":")
	if cap(smin) != 2 || cap(emin) != 2 {
		panic("time format error")
	}
	if strTOint(emin[0]) >= strTOint(smin[0]) {
		return (strTOint(emin[0])-strTOint(smin[0]))*60 + strTOint(emin[1]) - strTOint(smin[1])
	} else {
		return -1
	}
}

func strTOint(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return i
	}
}
