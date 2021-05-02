package tools

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	fmt.Println(GenerateToken(12345))
}
