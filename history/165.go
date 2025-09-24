package main

import "strconv"

func getNextSegment(version string, pos *int) int {
	if *pos >= len(version) {
		return 0
	}

	segment := ""
	for *pos < len(version) && version[*pos] == '0' {
		(*pos)++
	}

	for *pos < len(version) && version[*pos] != '.' {
		segment += string(version[*pos])
		(*pos)++
	}
	(*pos)++ // skip the dot

	result, _ := strconv.Atoi(segment)
	return result
}

func compareVersion(version1 string, version2 string) int {
	pos1, pos2 := 0, 0
	for pos1 < len(version1) || pos2 < len(version2) {
		seg1 := getNextSegment(version1, &pos1)
		seg2 := getNextSegment(version2, &pos2)
		if seg1 > seg2 {
			return 1
		}
		if seg1 < seg2 {
			return -1
		}
	}
	return 0
}
