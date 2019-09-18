package golib

import "sort"

type Pair struct {
	Key   string
	Value string
}
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func SortMapByValue(array map[string]string, sortingType string) PairList {
	p := make(PairList, len(array))
	i := 0
	for k, v := range array {
		p[i] = Pair{k, v}
		i++
	}
	if sortingType == "DESC" {
		sort.Reverse(p)
	} else {
		sort.Sort(p)
	}
	return p
}
