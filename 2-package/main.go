package main

import (
	"learn-go/calc" // moduleName/packageName
	"learn-go/calc/data"
)

// https://github.com/kubernetes/kubernetes/tree/master/pkg

func main() {
	calc.Add()
	//calc.Sum = 100
	calc.PrintSomething()

	data.DBConn = "mysql"
	data.GetData()

}
