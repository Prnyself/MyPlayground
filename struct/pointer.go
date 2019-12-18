package _struct

import "fmt"

type Task struct {
	Name string
}

func Call() {
	t := Task{Name: "original name"}
	t.ChangeName("first")
	fmt.Println(t)

	t.ChangeNameByPointer("second")
	fmt.Println(t)

	tp := &Task{Name: "pointer name"}
	tp.ChangeName("first")
	fmt.Println(tp)

	tp.ChangeNameByPointer("second")
	fmt.Println(tp)
}

func (t Task) ChangeName(name string) {
	t.Name = name + " changed in ChangeName"
	fmt.Println("t in ChangeName", t)
}

func (t *Task) ChangeNameByPointer(name string) {
	t.Name = name + " changed in ChangeNameByPointer"
	fmt.Println("t in ChangeNameByPointer", t)
}
