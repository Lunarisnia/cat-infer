package fwc

type Fact string

func containFact(facts []Fact, target Fact) bool {
	for _, f := range facts {
		if f == target {
			return true
		}
	}
	return false
}
