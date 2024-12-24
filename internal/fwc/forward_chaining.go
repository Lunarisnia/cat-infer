package fwc

// 1. We define rules with array of facts as conditions to be fulfilled
// 2. We take the defined knownFacts and then loop over it to see if it match
// 3. If it is then add the new conclusion to the knownFacts and go over to the next rule
// 4. Rinse and repeat until final conclusion is met
func ForwardChaining(knownFacts []Fact, rules []Rule) []Fact {
	knownFactChecklist := make(map[Fact]bool)
	for _, f := range knownFacts {
		knownFactChecklist[f] = true
	}

	for _, r := range rules {
		allFulfilled := true
		for _, cond := range r.Conditions {
			if !containFact(knownFacts, cond) {
				allFulfilled = false
			}

			if knownFactChecklist[r.Conclusion] {
				allFulfilled = false
			}
		}

		if allFulfilled {
			knownFacts = append(knownFacts, r.Conclusion)
		}
	}

	return knownFacts
}

func ForwardChainingInference(knownFacts []Fact, inferenceRules []Rule) *Fact {
	for _, r := range inferenceRules {
		found := true
		for _, cond := range r.Conditions {
			if !containFact(knownFacts, cond) {
				found = false
				break
			}
		}
		if found {
			return &r.Conclusion
		}
	}
	return nil
}
