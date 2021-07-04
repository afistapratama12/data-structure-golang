package stackqueue

import "fmt"

type Data interface {
	Add(val interface{})
	Remove()
}

type Stack struct {
	data []interface{}
}

type Queue struct {
	data []interface{}
}

func (q *Queue) Add(val interface{}) {
	q.data = append(q.data, val)
}

func (q *Queue) Remove() {
	q.data = q.data[1:]
}

func (s *Stack) Add(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Remove() {
	s.data = s.data[:len(s.data)-2]
}

func SandQ() {
	init := Queue{}

	data := []string{"keren", "asik", "hidup"}
	data2 := []int{24, 26, 30}

	for i := 0; i < len(data); i++ {
		init.Add(data[i])
	}

	for i := 0; i < len(data2); i++ {
		init.Add(data2[i])
	}

	fmt.Println(init.data...)
	init.Remove()
	init.Remove()
	init.Remove()

	fmt.Println(init.data...)
}
