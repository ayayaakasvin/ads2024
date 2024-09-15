package lab1

import (
	"fmt"
)

type Deque struct {
	Elements []int
	Size int
}

func (q *Deque) Front () (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("empty deque")
	}

	return q.Elements[0], nil
}

func (q *Deque) Back () (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("empty deque")
	}

	return q.Elements[q.Size - 1], nil
}

func (q *Deque) AddFront (elem int) {
	q.Elements = append([]int{elem}, q.Elements...)
	q.Size++
}

func (q *Deque) AddBack (elem int) {
	q.Elements = append(q.Elements, elem)
	q.Size++
}

func (q *Deque) Pop () (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("empty deque")
	}

	if q.Size == 1 {
		front, err := q.Front()
		if err != nil {
			return -1, err
		}
		
		q.Elements = []int{}
		q.Size--

		return front * 2, nil
	}

	front, err := q.Front()
	if err != nil {
		return -1, err
	}

	back, err := q.Back()
	if err != nil {
		return -1, err
	}

	q.Elements = q.Elements[1 : q.Size-1]
	q.Size -= 2
	return front + back, nil
}

func (q *Deque) IsEmpty () bool {
	return q.Size == 0
}

func (q *Deque) DisplayInfo () {
	fmt.Println("Displaying deque :")
	for _, v := range q.Elements {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func ConstructDeque() *Deque {
	q := &Deque{
		Elements:     []int{},
		Size:         0,
	}

	return q
}

func JStage(q *Deque) []any {
	var result []any

	for {
		var operation string

		fmt.Scanf("%s", &operation)

		if operation == "!" {
			break
		}

		switch operation {
		case "*":
			output, err := q.Pop()
			if err != nil {
				result = append(result, "error")
				break
			}

			result = append(result, output)
		case "+":
			var number int
			fmt.Scanf("%d", &number)

			q.AddFront(number)
		case "-":
			var number int
			fmt.Scanf("%d", &number)

			q.AddBack(number)
		default:
			fmt.Printf("Unexpected operation %s", operation)
		}
	}

	return result
}