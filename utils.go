package gomoduleexample

func greeting(name string) string {
	return "Hello, " + name
}

func getCharFromString(str string, idx int) string {
	temp := str[idx]
	return string(temp)
}