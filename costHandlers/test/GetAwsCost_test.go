package test

import (
	costhandlers "awscostapi/costHandlers"
	"fmt"
	"testing"
)

func TestGetAwsCosts(t *testing.T) {

	result := costhandlers.GetAwsCost()

	fmt.Println(result[17][0])
	
}
