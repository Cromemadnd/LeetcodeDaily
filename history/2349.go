package main

import (
	"github.com/emirpasic/gods/sets/treeset"
)

type NumberContainers struct {
	IndexToValue map[int]int
	ValueToIndex map[int]*treeset.Set
}

func Constructor() NumberContainers {
	return NumberContainers{
		IndexToValue: make(map[int]int),
		ValueToIndex: make(map[int]*treeset.Set),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if _, found := this.IndexToValue[index]; found {
		value := this.IndexToValue[index]
		if set, ok := this.ValueToIndex[value]; ok {
			set.Remove(index)
		}
	}

	this.IndexToValue[index] = number
	if this.ValueToIndex[number] == nil {
		this.ValueToIndex[number] = treeset.NewWithIntComparator()
	}
	this.ValueToIndex[number].Add(index)
}

func (this *NumberContainers) Find(number int) int {
	if set, ok := this.ValueToIndex[number]; ok {
		it := set.Iterator()
		if it.First() {
			return it.Value().(int)
		}
	}
	return -1
}

func main() {
	commands := []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"}
	params := [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10}, {10}, {1, 20}, {10}}

	var obj NumberContainers
	for i, cmd := range commands {
		switch cmd {
		case "NumberContainers":
			obj = Constructor()
		case "change":
			index, number := params[i][0], params[i][1]
			obj.Change(index, number)
		case "find":
			number := params[i][0]
			result := obj.Find(number)
			println(result)
		}
	}
}
