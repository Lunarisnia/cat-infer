package fwc

import (
	"fmt"
	"testing"
)

func Test_ForwardChaining(t *testing.T) {
	t.Run("Basic Rule", func(t *testing.T) {
		knownFacts := make([]Fact, 0)
		knownFacts = append(knownFacts, Gentle)
		knownFacts = append(knownFacts, LaidBack)
		knownFacts = append(knownFacts, RoundEyes)
		knownFacts = append(knownFacts, SmallEars)
		knownFacts = append(knownFacts, BushyTail)
		knownFacts = append(knownFacts, LargeSize)
		knownFacts = append(knownFacts, RoundFace)
		knownFacts = append(knownFacts, HighMaintenance)
		knownFacts = append(knownFacts, Vocal)

		inferenceRule := Cat

		inferenceResult := ForwardChainingInference(knownFacts, inferenceRule)
		if inferenceResult == nil {
			fmt.Println("Unindentifiable")
			return
		}
		fmt.Println("Inference Result: ", *inferenceResult)
	})
}
