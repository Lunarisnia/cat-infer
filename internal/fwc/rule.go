package fwc

type Rule struct {
	Conditions []Fact
	Conclusion Fact
}
