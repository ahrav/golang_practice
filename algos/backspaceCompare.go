package main

func backspaceCompare2(str1, str2 string) bool {
	str1 = stripBackspaces(str1)
	str2 = stripBackspaces(str2)
	return str1 == str2
}

func stripBackspaces(str string) string {
	var res string
	c := 0
	for i:=len(str)-1; i>=0; i-- {
		if string(str[i]) != "#" {
			if c != 0 {
				c--
				continue
			}
			res=string(str[i])+res
		} else {
			c++
		}
	}
	return res
}

func backspaceCompare(str1, str2 string) bool {
	idx1, idx2 := len(str1)-1, len(str2)-1
	for idx1 >= 0 || idx2 >= 0 {
		i1 := nextValidChar(str1, idx1)
		i2 := nextValidChar(str2, idx2)
		if i1 < 0 && i2 < 0 {
			return true
		}
		if i1 < 0 || i2 < 0 {
			return false
		}
		if string(str1[i1]) != string(str2[i2]) {
			return false
		}
		idx1 = i1 - 1
		idx2 = i2 - 1
	}
	return true
}

func nextValidChar(str string, idx int) int {
	c := 0
	for idx >= 0 {
		if string(str[idx]) == "#" {
			c++
		} else if c > 0 {
			c--
		} else {
			break
		}
		idx--
	}
	return idx
}
