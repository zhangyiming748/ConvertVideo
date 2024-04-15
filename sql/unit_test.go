package sql

import (
	"testing"
)

// go test v -cgo -run TestInit
func init() {
	SetEngine()
}
func TestSetOne(t *testing.T) {
	c := new(Conv)
	c.SrcName = "1"
	c.SetOne()
}
