package gomoduleexample

func Greeting(name string) string {
	return "Hello, " + name
}

func GetCharFromString(str string, idx int) string {
	temp := str[idx]
	return string(temp)
}