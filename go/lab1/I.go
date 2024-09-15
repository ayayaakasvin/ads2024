package lab1

import (
	"fmt"
)

type Student struct {
	Team byte
}

func StudentConstructor (team byte) *Student {
	return &Student{
		Team: team,
	}
}

type Stack []*Student

func (s *Stack) Top() *Student {
    if len(*s) == 0 {
        return nil
    }
    return (*s)[len(*s)-1]
}

func (s *Stack) Push(v *Student) {
    *s = append((*s), v)
}

func (s *Stack) Pop() {
    if len(*s) == 0 {
		fmt.Printf("underflow\n")
		return
	}

	*s = (*s)[len(*s) - 1:]
}

func (s *Stack) Empty() bool {
	return len((*s)) == 0
}

func NewStack () *Stack {
	return &Stack{}
}

func teamDominance(input []byte) bool {
	dominant := input[0]

	for _, v := range input {
		if v != dominant {
			return false
		}
	}

	return true
}

func ISolution(voteOrder string) string {
	var (
		KStack *Stack = NewStack()
		SStack *Stack = NewStack()
		mutableVoteOrder []byte = []byte(voteOrder)
	)

	for !teamDominance(mutableVoteOrder) {
		for i := 0; i < len(mutableVoteOrder); {
			if mutableVoteOrder[i] == 'K' {
				if !SStack.Empty() {
					mutableVoteOrder = append(mutableVoteOrder[:i], mutableVoteOrder[i+1:]...)
					SStack.Pop()
					continue
				}

				KStack.Push(StudentConstructor(mutableVoteOrder[i]))
			} else {
				if !KStack.Empty() {
					mutableVoteOrder = append(mutableVoteOrder[:i], mutableVoteOrder[i+1:]...)
					KStack.Pop()
					continue
				}

				SStack.Push((StudentConstructor(mutableVoteOrder[i])))
			}

			i++
		}
	}

	if mutableVoteOrder[0] == 'K' {
		return "KATSURAGI"
	} else {
		return "SAKAYANAGI"
	}
}