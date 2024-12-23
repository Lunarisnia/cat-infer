package main

import (
	"fmt"

	"github.com/Lunarisnia/cat-infer/internal/fwc"
)

func main() {
	knownFacts := make([]fwc.Fact, 0)
	knownFacts = append(knownFacts, fwc.Fact("BlackFooted"))
	knownFacts = append(knownFacts, fwc.Fact("Small"))
	knownFacts = append(knownFacts, fwc.Fact("RoundEyes"))
	knownFacts = append(knownFacts, fwc.Fact("Cat"))

	rules := make([]fwc.Rule, 0)
	rules = append(rules, fwc.Rule{
		Conditions: []fwc.Fact{"BlackFooted"},
		Conclusion: "Cat",
	})
	rules = append(rules, fwc.Rule{
		Conditions: []fwc.Fact{"Small", "SquareEyes"},
		Conclusion: "MinecraftCat",
	})
	rules = append(rules, fwc.Rule{
		Conditions: []fwc.Fact{"Small", "RoundEyes"},
		Conclusion: "NaturallySmallCat",
	})
	rules = append(rules, fwc.Rule{
		Conditions: []fwc.Fact{"NaturallySmallCat"},
		Conclusion: "Cute",
	})

	newFacts := fwc.ForwardChaining(knownFacts, rules)

	inferenceRule := make([]fwc.Rule, 0)
	inferenceRule = append(inferenceRule, fwc.Rule{
		Conditions: []fwc.Fact{"Cat", "Cute", "NaturallySmallCat", "RoundEyes"},
		Conclusion: "Munchkin",
	})

	inferenceResult := fwc.ForwardChainingInference(newFacts, inferenceRule)
	if inferenceResult == nil {
		fmt.Println("Unindentifiable")
	}
	fmt.Println("Inference Result: ", *inferenceResult)
}
