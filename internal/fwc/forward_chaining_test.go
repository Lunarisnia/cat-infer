package fwc

import (
	"fmt"
	"testing"
)

func Test_ForwardChaining(t *testing.T) {
	t.Run("Basic Rule", func(t *testing.T) {
		knownFacts := make([]Fact, 0)
		knownFacts = append(knownFacts, Fact("BlackFooted"))
		knownFacts = append(knownFacts, Fact("Small"))
		knownFacts = append(knownFacts, Fact("RoundEyes"))
		knownFacts = append(knownFacts, Fact("Cat"))

		rules := make([]Rule, 0)
		rules = append(rules, Rule{
			Conditions: []Fact{"BlackFooted"},
			Conclusion: "Cat",
		})
		rules = append(rules, Rule{
			Conditions: []Fact{"Small", "SquareEyes"},
			Conclusion: "MinecraftCat",
		})
		rules = append(rules, Rule{
			Conditions: []Fact{"Small", "RoundEyes"},
			Conclusion: "NaturallySmallCat",
		})
		rules = append(rules, Rule{
			Conditions: []Fact{"NaturallySmallCat"},
			Conclusion: "Cute",
		})

		newFacts := ForwardChaining(knownFacts, rules)
		fmt.Println("New Facts: ", newFacts)
	})
}
