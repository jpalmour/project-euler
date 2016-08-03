package main

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	for _, test := range []struct {
		nums      []int
		length    int
		expCombos [][]int
	}{
		{[]int{}, 0, [][]int{{}}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1, 2}, 1, [][]int{{1}, {2}}},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {1, 3}, {2, 3}}},
	} {
		combos := combinations(test.nums, test.length)
		if !reflect.DeepEqual(test.expCombos, combos) {
			t.Errorf("combinations(%v, %v)=%v, expected %v", test.nums, test.length, combos, test.expCombos)
		}
	}
}

func TestPermutations(t *testing.T) {
	for _, test := range []struct {
		nums     []int
		expPerms [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	} {
		perms := permutations(test.nums)
		if !reflect.DeepEqual(test.expPerms, perms) {
			t.Errorf("permutations(%v)=%v, expected %v", test.nums, perms, test.expPerms)
		}
	}
}

func TestIntPermutations(t *testing.T) {
	for _, test := range []struct {
		nums     []int
		expPerms []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{12, 21}},
		{[]int{1, 2, 3}, []int{123, 132, 213, 231, 312, 321}},
	} {
		perms := intPermutations(test.nums)
		if !reflect.DeepEqual(test.expPerms, perms) {
			t.Errorf("intPermutations(%v)=%v, expected %v", test.nums, perms, test.expPerms)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, test := range []struct {
		a       []int
		b       []int
		expDiff []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3, 4}, []int{3, 4}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{}, []int{1, 2, 3, 4}},
		{[]int{}, []int{1, 2, 3}, []int{}},
	} {
		diff := difference(test.a, test.b)
		if !reflect.DeepEqual(diff, test.expDiff) {
			t.Errorf("diff(%v, %v)=%v, expected %v", test.a, test.b, diff, test.expDiff)
		}
	}
}

func TestDigitMatch(t *testing.T) {
	for _, test := range []struct {
		num      int
		digits   []int
		expMatch bool
	}{
		{1234, []int{1, 2, 3, 4}, true},
		{1111, []int{1}, false},
		{84572, []int{7, 4, 5, 2, 8}, true},
		{83572, []int{7, 4, 5, 2, 8}, false},
		{845722, []int{7, 4, 5, 2, 8}, false},
	} {
		match := digitMatch(test.num, test.digits)
		if match != test.expMatch {
			t.Errorf("digitMatch(%v, %v)=%v, expected %v", test.num, test.digits, match, test.expMatch)
		}
	}
}
