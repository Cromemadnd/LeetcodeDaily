package main

import (
	"strconv"
)

type Spreadsheet struct {
	Data [26]map[int]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{}
}

func (ss *Spreadsheet) SetCell(cell string, value int) {
	col := cell[0] - 'A'
	row, _ := strconv.Atoi(cell[1:])
	if ss.Data[col] == nil {
		ss.Data[col] = make(map[int]int)
	}
	ss.Data[col][row] = value
}

func (ss *Spreadsheet) ResetCell(cell string) {
	ss.SetCell(cell, 0)
}

func (ss *Spreadsheet) GetCell(cell string) int {
	col := cell[0] - 'A'
	row, _ := strconv.Atoi(cell[1:])
	if ss.Data[col] == nil {
		return 0
	}
	return ss.Data[col][row]
}

func isReference(s string) bool {
	return len(s) > 1 && s[0] >= 'A' && s[0] <= 'Z'
}

func (ss *Spreadsheet) GetValue(formula string) int {
	index := 1
	for formula[index] != '+' {
		index++
	}
	a, b := formula[1:index], formula[index+1:]
	aval, bval := 0, 0

	if isReference(a) {
		aval = ss.GetCell(a)
	} else {
		aval, _ = strconv.Atoi(a)
	}

	if isReference(b) {
		bval = ss.GetCell(b)
	} else {
		bval, _ = strconv.Atoi(b)
	}

	return aval + bval
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
