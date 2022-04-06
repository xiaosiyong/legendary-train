package main

/*
将string中空格换成%20
*/
func replaceString(str string) string {
	if len(str) <= 1 {
		return str
	}
	//求出空格个数
	var result string
	for _, v := range str {
		if string(v) == " " {
			result += "%20"
		} else {
			result += string(v)
		}
	}
	return result

}
