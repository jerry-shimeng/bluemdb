package bmcache
import "fmt"


func Get(key, value string) string {
	r := data[key]
//	fmt.Println(data)
	return r
}

func Set(key, value string) string {
	data[key] = value
	return "1"
}

func Del(key, value string) string {
	delete(data, key)
	return "1"
}

func Add(key, value string) string {
	Set(key, value)
	return "1"
}

func Append(key ,value string)string{
	s := data[key]
	s = fmt.Sprint(s,value)
	return Set(key,s)
}

//设置并获取更改的值
func GetSet(key ,value string)string{
	s := data[key]
	data[key] = value
	return s
}