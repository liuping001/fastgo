// Copyright 2021, Tencent Inc.
// Author: coolliu
// Date: 2021/5/24

package util

import (
	"fmt"
	"testing"
)

func errorB() {
	err := Errorf("time out")
	fmt.Printf("%s\n", err.Error())
}

func errorA() {
	errorB()
}

func TestError(t *testing.T) {
	errorA()
}
