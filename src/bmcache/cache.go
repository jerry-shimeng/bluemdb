package bmcache

import "fmt"

//type BmCache struct {
//
//}
//
//func (this *BmCache)Run(){
//
//}

var data map[string]string

func CacheInit() {
	data = make(map[string]string, 0)

	CommandInit()
	ChannelInit()
	fmt.Println("cache server is running")
}
//获取数据对象的引用
func GetDataMap()*map[string]string{
	return &data
}


