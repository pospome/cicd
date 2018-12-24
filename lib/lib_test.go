package lib_test

import (
	"github.com/pospome/cicd/lib"
	"testing"
)

func TestNewMyStruct(t *testing.T) {
	s := lib.NewMyStruct("test")
	if s.Name != "test" {
		t.Errorf("error")
	}
}
