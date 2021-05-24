// Author: coolliu
// Date: 2021/5/24

package util

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type ErrorStack struct {
	error
}
type StackInfo struct {
	name string
	line string
}

func (s *StackInfo) Name(name string) {
	name = strings.Trim(name, " ")
	strs := strings.Split(name, "(")
	if len(strs) <= 0 {
		return
	}
	name = strs[0]
	strs = strings.Split(name, ".")
	if len(strs) <= 0 {
		return
	}
	s.name = strs[len(strs)-1]
}
func (s *StackInfo) Line(line string) {
	line = strings.Trim(line, " ")
	strs := strings.Split(line, " ")
	if len(strs) <= 0 {
		return
	}
	line = strs[0]
	strs = strings.Split(line, "/")
	if len(strs) <= 0 {
		return
	}
	s.line = strs[len(strs)-1]
}
func Errorf(info string) error {
	stack := string(debug.Stack())
	stacks := strings.Split(stack, "\n")

	shortStacks := []StackInfo{}
	for i := 5; i+2 < len(stacks); i = i + 2 {
		var stackInfo StackInfo
		stackInfo.Name(stacks[i])
		stackInfo.Line(stacks[i+1])
		shortStacks = append(shortStacks, stackInfo)
	}

	msg := ""
	for i := len(shortStacks) - 1; i >= 0; i-- {
		if msg != "" {
			msg += "-> "
		}
		msg += fmt.Sprintf("%s(%s) ", shortStacks[i].name, shortStacks[i].line)
	}
	err := fmt.Errorf("%s: %s", msg, info)
	return &ErrorStack{error: err}
}
