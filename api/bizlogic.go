package api

import (
	"math/rand"
	"time"
)

func Random(leng int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(leng)
	return randomNumber
}

func Quicksort(slc []int) []int {
	if len(slc) < 2 {
		return slc
	}
	i := Random(len(slc))
	pivot := slc[i]
	slc[i], slc[len(slc)-1] = slc[len(slc)-1], slc[i]
	var hslc []int
	var lslc []int
	for _, num := range slc[:len(slc)-1] {
		if num >= pivot {
			hslc = append(hslc, num)
		} else {
			lslc = append(lslc, num)
		}
	}
	high := Quicksort(hslc)
	low := Quicksort(lslc)

	//final1:=append(low, slc[0])
	return append(append(low, pivot), high...)
}

func Repcount(str string) (map[string]int, string, int) {
	m := make(map[string]int)
	for _, v := range str {
		s := string(v)
		m[s] = m[s] + 1
	}
	highest := 0
	var t string
	for key, val := range m {
		if val >= highest {
			highest = val
			t = key
		}
	}
	return m, t, highest
}
