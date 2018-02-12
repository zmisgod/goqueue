package gostruct

import (
	"fmt"
	"testing"
)

func TestStack (t *testing.T) {
	stackSize := 2
	mystack := CreateStack(stackSize)
	var stackone Stack
	stackone.Name = "zmisgo"
	stackone.Job = "php"
	stackone.Age = 12
	mystack.Push(stackone)
	var stacktwo Stack
	stacktwo.Name = "zmisgod"
	stacktwo.Job = "golang"
	stacktwo.Age = 13
	mystack.Push(stacktwo)

	var popStack Stack
	mystack.Pop(&popStack)
	fmt.Println(popStack)
	mystack.StackTraverse()
}