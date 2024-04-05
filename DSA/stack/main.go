package main

import (
	"fmt"
)

type Stack struct {
	top int
	data []int
}

func main() {
	var myStack Stack = Stack{0, []int {1, 2,3,4} };

	push(&myStack, 1);

	printStack(&myStack);

}

func push (s *Stack, val int) {
	if (s.top >= len(s.data)) {
		fmt.Println("Stack overflow");
		return;
	} 

	s.data = append(s.data, val)
	s.top++;
}

func pop (s *Stack) int {
	if (s.top < 0) {
		fmt.Println("Stack underflow");
		return -1;
	}

	x := s.data[s.top];
	s.top--;
	return x;
}

/*
	returns the top element of the stack without removing it
	* returns true if the stack is empty
*/
func peak (s *Stack) int {return 0;}

func printStack (s *Stack) {
	var i int;
	for i = 0 ; i < s.top; i++ {
		fmt.Println(s.data[i]);
	};
}
