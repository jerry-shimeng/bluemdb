package bmmodel


type ReturnModel struct  {
	Status int
	Message string
	Data interface{}
}

func (this *ReturnModel)ToJson() string {
	v := interface{}(*this)
	s:= JsonConvertToString(&v)
	return s
}


func ErrorResult(s string)*ReturnModel{
	m := ReturnModel{1,s,nil}
	return &m
}


func SuccessResult(s string)*ReturnModel{
	m := ReturnModel{0,"ok",s}
	return &m
}