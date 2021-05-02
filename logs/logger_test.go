package logs

import (
	"fmt"
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	Logger.Info("I'm about to do something!")
	Logger.Errorf("Error running doSomething: %v", "err")
}

func TestName(t *testing.T) {
	try := make(map[string]interface{})
	try["sub"] = int64(1)
	switch try["sub"].(type) {
	case float64:
		fmt.Println("float64")
	case int64:
		fmt.Println("int64")
	default:
		fmt.Println("other")
	}
	a := try["sub"].(int64)
	fmt.Println(a)
}
