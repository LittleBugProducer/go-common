package main

import (
	"time"
)

func main() {

	dimissionContext := &DimissionContext{
		Name:           "张三",
		WorkingYears:   3,
		DepartmentName: "技术部",
		LastDay:        time.Now().Add(time.Hour * 24 * 30),
	}

	handleNode := NewDimissionHandlerChain()
	err := handleNode.handle(dimissionContext)
	if err != nil {
		panic(err)
	}

}
