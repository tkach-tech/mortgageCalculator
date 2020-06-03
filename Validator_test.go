package main

import (
	"testing"
	"testing/quick"
)

func TestValidate(t *testing.T)  {
	t.Parallel()

	var err error = &quick.CheckError{}

	validate(err)
}
