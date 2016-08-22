package generator

import (
	"testing"
)

////////////////
// Benchmarks //
////////////////

func factorial(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

var result [][]int

// AA = (s, s2 []int) ; No Closure
// AB = (s, s2 []int) ; Closure
// BA = (s []int, length, depth int) ; No Closure
// BB = (s []int, length, depth int) ; Closure

// 1AA

func perm1AA(s, s2 []int) {
	if len(s2) == 0 {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return
	}
	for i := range s2 {
		s2[0], s2[i] = s2[i], s2[0]
		perm1AA(s, s2[1:])
		s2[0], s2[i] = s2[i], s2[0]
	}
}

func Perm1AA(s []int) {
	result = make([][]int, 0, factorial(len(s)))
	perm1AA(s, s)
	result = nil
}

// 1AB

func perm1AB(consumer func([]int), s, s2 []int) {
	if len(s2) == 0 {
		consumer(s)
		return
	}
	for i := range s2 {
		s2[0], s2[i] = s2[i], s2[0]
		perm1AB(consumer, s, s2[1:])
		s2[0], s2[i] = s2[i], s2[0]
	}
}

func Perm1AB(s []int) {
	result = make([][]int, 0, factorial(len(s)))
	consumer := func(s []int) {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
	}
	perm1AB(consumer, s, s)
	result = nil
}

// 1BA

func perm1BA(s []int, length, depth int) {
	if depth == length {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return
	}
	for i := depth; i < length; i++ {
		s[depth], s[i] = s[i], s[depth]
		perm1BA(s, length, depth+1)
		s[depth], s[i] = s[i], s[depth]
	}
}

func Perm1BA(s []int) {
	result = make([][]int, 0, factorial(len(s)))
	perm1BA(s, len(s), 0)
	result = nil
}

// 1BB

func perm1BB(consumer func([]int), s []int, length, depth int) {
	if depth == length {
		consumer(s)
		return
	}
	for i := depth; i < length; i++ {
		s[depth], s[i] = s[i], s[depth]
		perm1BB(consumer, s, length, depth+1)
		s[depth], s[i] = s[i], s[depth]
	}
}

func Perm1BB(s []int) {
	result = make([][]int, 0, factorial(len(s)))
	consumer := func(s []int) {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
	}
	perm1BB(consumer, s, len(s), 0)
	result = nil
}

// 2AA

func perm2AA(s, s2 []int) bool {
	if len(s2) == 0 {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return true
	}
	for i := range s2 {
		s2[0], s2[i] = s2[i], s2[0]
		if !perm2AA(s, s2[1:]) {
			return false
		}
		s2[0], s2[i] = s2[i], s2[0]
	}
	return true
}

func Perm2AA(s []int) bool {
	result = make([][]int, 0, factorial(len(s)))
	retval := perm2AA(s, s)
	result = nil
	return retval
}

// 2AB

func perm2AB(predicate func([]int) bool, s, s2 []int) bool {
	if len(s2) == 0 {
		return predicate(s)
	}
	for i := range s2 {
		s2[0], s2[i] = s2[i], s2[0]
		if !perm2AB(predicate, s, s2[1:]) {
			return false
		}
		s2[0], s2[i] = s2[i], s2[0]
	}
	return true
}

func Perm2AB(s []int) bool {
	result = make([][]int, 0, factorial(len(s)))
	predicate := func(s []int) bool {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return true
	}
	retval := perm2AB(predicate, s, s)
	result = nil
	return retval
}

// 2BA

func perm2BA(s []int, length, depth int) bool {
	if depth == length {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return true
	}
	for i := depth; i < length; i++ {
		s[depth], s[i] = s[i], s[depth]
		if !perm2BA(s, length, depth+1) {
			return false
		}
		s[depth], s[i] = s[i], s[depth]
	}
	return true
}

func Perm2BA(s []int) bool {
	result = make([][]int, 0, factorial(len(s)))
	retval := perm2BA(s, len(s), 0)
	result = nil
	return retval
}

// 2BB

func perm2BB(predicate func([]int) bool, s []int, length, depth int) bool {
	if depth == length {
		return predicate(s)
	}
	for i := depth; i < length; i++ {
		s[depth], s[i] = s[i], s[depth]
		if !perm2BB(predicate, s, length, depth+1) {
			return false
		}
		s[depth], s[i] = s[i], s[depth]
	}
	return true
}

func Perm2BB(s []int) bool {
	result = make([][]int, 0, factorial(len(s)))
	predicate := func(s []int) bool {
		sCopy := make([]int, len(s))
		copy(sCopy, s)
		result = append(result, sCopy)
		return true
	}
	perm2BB(predicate, s, len(s), 0)
	result = nil
	return true
}

func benchmarkPerm1(b *testing.B, perm1Func func([]int)) {
	for i := 0; i < b.N; i++ {
		perm1Func([]int{1, 2, 3, 4, 5, 6})
	}
}

func BenchmarkPerm1AA(b *testing.B) {
	benchmarkPerm1(b, Perm1AA)
}

func BenchmarkPerm1AB(b *testing.B) {
	benchmarkPerm1(b, Perm1AB)
}

func BenchmarkPerm1BA(b *testing.B) {
	benchmarkPerm1(b, Perm1BA)
}

func BenchmarkPerm1BB(b *testing.B) {
	benchmarkPerm1(b, Perm1BB)
}

func benchmarkPerm2(b *testing.B, perm2Func func([]int) bool) {
	for i := 0; i < b.N; i++ {
		perm2Func([]int{1, 2, 3, 4, 5, 6})
	}
}

func BenchmarkPerm2AA(b *testing.B) {
	benchmarkPerm2(b, Perm2AA)
}

func BenchmarkPerm2AB(b *testing.B) {
	benchmarkPerm2(b, Perm2AB)
}

func BenchmarkPerm2BA(b *testing.B) {
	benchmarkPerm2(b, Perm2BA)
}

func BenchmarkPerm2BB(b *testing.B) {
	benchmarkPerm2(b, Perm2BB)
}
