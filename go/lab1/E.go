package lab1

import (
	"fmt"
	"reflect"
)

type Queue struct {
	Elements []any
	Size int
}

type GameEnd struct {
	IsDraw bool
	Winner string
	Rounds int
}

func (q *Queue) Top () any {
	if q.IsEmpty() {
		return nil
	}

	return q.Elements[0]
}

func (q *Queue) Add (elems ...any) {
	if len(elems) == 0 {
		return
	}

	for _, elem := range elems {
        v := reflect.ValueOf(elem)
        if v.Kind() == reflect.Slice {
            for i := 0; i < v.Len(); i++ {
                q.Elements = append(q.Elements, v.Index(i).Interface())
                q.Size++
            }
        } else {
            q.Elements = append(q.Elements, elem)
            q.Size++
        }
    }
}

func (q *Queue) Pop () any {
	if q.IsEmpty() {
		return nil
	}

	var temp any = q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Size--
	return temp
}

func (q *Queue) IsEmpty () bool {
	return q.Size == 0
}

func (q *Queue) DisplayInfo () {
	fmt.Println("Displaying queue :")
	var counter int
	for _, v := range q.Elements {
		fmt.Printf("%v ", v)
		counter++
	}
	fmt.Println()
}

func ConstructQueue() *Queue {
	return &Queue{}
}

func ConstructQueueWithElems(elems ...any) *Queue {
	result := ConstructQueue()
	result.Add(elems...)
	return result
}

func GameStage (Q1, Q2 *Queue) GameEnd {
	var RoundCount int

	for !Q1.IsEmpty() && !Q2.IsEmpty() && RoundCount < 1000000 {
		var (
			firstElem = Q1.Pop()
			secondElem = Q2.Pop()
		)

		if firstElem == nil || secondElem == nil {
			break
		}

		var (
			FirstQueueTop int = firstElem.(int)
			SecondQueueTop int = secondElem.(int)
		)

		if FirstQueueTop == 0 && SecondQueueTop == 9 {
			Q1.Add(FirstQueueTop, SecondQueueTop)
		} else if FirstQueueTop == 9 && SecondQueueTop == 0 {
			Q2.Add(SecondQueueTop, FirstQueueTop)		
		} else if FirstQueueTop > SecondQueueTop {
			Q1.Add(FirstQueueTop, SecondQueueTop)
		} else {
			Q2.Add(SecondQueueTop, FirstQueueTop)
		}

		RoundCount++
	}

	if RoundCount >= 1000000 {
		return GameEnd{
			IsDraw: true,
		}
	} else if Q1.IsEmpty() {
		return GameEnd{
			IsDraw: false,
			Winner: "Nursik",
			Rounds: RoundCount,
		}
	} else if Q2.IsEmpty() {
		return GameEnd{
			IsDraw: false,
			Winner: "Boris",
			Rounds: RoundCount,
		}
	} else {
		return GameEnd{
			IsDraw: true,
		}
	}
}

func ESolution(nursik, boris []int) string {
	NursiksQ := ConstructQueueWithElems(nursik)
	BorissQ := ConstructQueueWithElems(boris)

	Result := GameStage(NursiksQ, BorissQ)

	if Result.IsDraw {
		return "Draw"
	} else {
		return fmt.Sprintf("%s %d", Result.Winner, Result.Rounds)
	}
}