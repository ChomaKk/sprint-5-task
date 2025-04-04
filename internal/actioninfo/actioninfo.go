package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {

	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		res, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
