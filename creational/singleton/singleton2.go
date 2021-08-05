package singleton

type single2 struct{}

var singleInstance2 *single2

func init() {
	singleInstance2 = &single2{}
}

func GetInstance2() *single2 {
	return singleInstance2
}
