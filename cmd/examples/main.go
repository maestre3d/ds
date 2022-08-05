package main

import (
	"log"
	"strings"

	"github.com/maestre3d/ds"
)

func main() {
	slA := []int{1, 2, 3, 4, 5}
	slB := []int{0, 2, 4, 6, 8, 9}

	log.Print(detectIntersection(slA, slB))

	log.Print(detectDupe([]string{"a", "b", "c", "d", "c", "e", "f"}))
	log.Print(detectMissingAlphabetLetter("the quick brown box jumps over the lazy dog"))
	log.Print(checkLint("{ var x = { y: [1, 2, 3] } )"))
	log.Print(reverseString("tenet"))
	countToN(10, 1)

	traverseKNaryStruct([]any{
		1,
		2,
		3,
		[]any{4, 5, 6},
		7,
		[]any{
			8,
			[]any{
				9, 10, 11,
				[]any{12, 13, 14},
			},
		},
		[]any{
			15, 16, 17, 18, 19,
			[]any{
				20, 21, 22,
				[]any{
					23, 24, 25,
					[]any{
						26, 27, 28,
					}, 29, 30, 31,
				}, 32,
			}, 33,
		},
	}, 0)
}

func detectIntersection(slA, slB []int) []int {
	occurSet := make(ds.HashSet[int], len(slA))
	occurSet.Append(slA...)

	intersections := make([]int, 0)
	for _, item := range slB {
		if occurSet.Exists(item) {
			intersections = append(intersections, item)
		}
	}
	return intersections
}

func detectDupe(sl []string) string {
	set := make(ds.HashSet[string], len(sl))

	for _, item := range sl {
		if set.Exists(item) {
			return item
		}
		set.Append(item)
	}
	return ""
}

func detectMissingAlphabetLetter(sentence string) string {
	const (
		startAlphabetIndex = 97
		endAlphabetIndex   = 122
	)

	set := make(ds.HashSet[byte], endAlphabetIndex-startAlphabetIndex)
	for i := range sentence {
		set.Append(sentence[i])
	}

	for i := startAlphabetIndex; i <= endAlphabetIndex; i++ {
		c := byte(i)
		if !set.Exists(c) {
			return string(c)
		}
	}
	return ""
}

func checkLint(ln string) bool {
	closingMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := ds.NewSliceStack[byte](0)
	for i := range ln {
		c := ln[i]
		isClosing := c == '}' || c == ']' || c == ')'
		if c == '{' || c == '[' || c == '(' {
			stack.Push(c)
		} else if isClosing && stack.Len() > 0 {
			out := stack.Peek()
			if closingMap[c] == out {
				_ = stack.Pop()
			}
		} else if isClosing && stack.Len() == 0 {
			return false
		}
	}

	return stack.Len() == 0
}

func reverseString(v string) string {
	if len(v) == 0 {
		return ""
	}

	stack := ds.NewSliceStack[byte](len(v))
	for i := range v {
		c := v[i]
		stack.Push(c)
	}

	buf := strings.Builder{}
	for i := 0; i < len(v); i++ {
		buf.WriteByte(stack.Pop())
	}
	return buf.String()
}

func countToN(stop, c int) {
	if c > stop {
		return
	}
	log.Print(c)
	c++
	countToN(stop, c)
}

func traverseKNaryStruct(in []any, index int) {
	if len(in) == index {
		return
	}

	switch in[index].(type) {
	case []any:
		childNode := in[index].([]any)
		traverseKNaryStruct(childNode, 0)
	default:
		log.Print(in[index])
	}

	index++
	traverseKNaryStruct(in, index)
}
