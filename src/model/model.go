package model

import (
	t1 "github.com/gefracto/kostrika-go-tasks/src/task1"
	t4 "github.com/gefracto/kostrika-go-tasks/src/task4"
	t5 "github.com/gefracto/kostrika-go-tasks/src/task5"
	t6 "github.com/gefracto/kostrika-go-tasks/src/task6"
	t7 "github.com/gefracto/kostrika-go-tasks/src/task7"
)

type T1 t1.Data

type T2 struct {
	Width  int
	Height int
}

//type Triangle t3.Triangle
//type T3 t3.T3
type Triangle struct {
	Name string
	A    float64
	B    float64
	C    float64
}

type T3 struct {
	SliceOfTriangles []Triangle
}

type T4 t4.T4

type T5 t5.T5

type T6 t6.T6

type T7 t7.T7
