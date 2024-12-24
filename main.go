package main

import (
	"fmt"

	"github.com/Lunarisnia/cat-infer/internal/fwc"
	"github.com/Lunarisnia/cat-infer/internal/userio"
)

const furAndCoat = "Fur & Coat"

const eyes = "Eyes"

const ears = "Ears"

const tail = "Tail"

const build = "Build"

const facialFeatures = "Facial Features"

const behaviourAndTemprament = "Behaviour & Temprament"

const uniqueTraits = "Unique Traits"

const careNeeds = "Care Needs"

const intelligence = "Intelligence & Communication"

var categories = []string{
	furAndCoat,
	eyes,
	ears,
	tail,
	build,
	facialFeatures,
	behaviourAndTemprament,
	uniqueTraits,
	careNeeds,
	intelligence,
}

var factList map[string][]fwc.Fact = map[string][]fwc.Fact{
	// Fur and Coat
	furAndCoat: {
		fwc.LongFur,
		fwc.SemiLongHair,
		fwc.ShortCoat,
		fwc.LongShaggyCoat,
		fwc.SpottedCoat,
		fwc.MarbledCoat,
		fwc.StripedCoat,
		fwc.ShinyCoat,
		fwc.TickedTabbyCoat,
		fwc.Hairless,
		fwc.WrinkledSkin,
		fwc.ReddishCoat,
		fwc.ThickFur,
	},

	// Eyes
	eyes: {
		fwc.RoundEyes,
		fwc.LargeEyes,
		fwc.BlueEyes,
		fwc.AlmondEyes,
	},

	// Ears
	ears: {
		fwc.SmallEars,
		fwc.TuftedEars,
		fwc.FoldedEars,
		fwc.LargeEars,
	},

	// Tail
	tail: {
		fwc.BushyTail,
		fwc.ShortTail,
		fwc.LongTail,
	},

	// Build
	build: {
		fwc.LargeSize,
		fwc.SleekBody,
		fwc.AthleticBuild,
		fwc.MuscularBody,
		fwc.LongLegs,
		fwc.SlimBuild,
	},

	// Facial Features
	facialFeatures: {
		fwc.FlatFace,
		fwc.RoundFace,
		fwc.TearMarks,
	},

	// Behavior and Temperament
	behaviourAndTemprament: {
		fwc.Gentle,
		fwc.LaidBack,
		fwc.Friendly,
		fwc.Social,
		fwc.Affectionate,
		fwc.Docile,
		fwc.Playful,
		fwc.Curious,
		fwc.Active,
		fwc.HighlyEnergetic,
		fwc.Adaptable,
		fwc.PeopleOriented,
		fwc.LovesBeingHeld,
		fwc.Solitary,
	},

	// Unique Traits
	uniqueTraits: {
		fwc.ColorPoints,
		fwc.Mane,
		fwc.LivesInPrides,
		fwc.ColdAdapted,
		fwc.ExcellentClimber,
	},

	// Care Needs
	careNeeds: {
		fwc.HighMaintenance,
		fwc.ModerateMaintenance,
		fwc.LowMaintenance,
	},

	// Intelligence and Communication
	intelligence: {
		fwc.Vocal,
		fwc.Intelligence,
		fwc.Quiet,
	},
}

func main() {
	knownFacts := make([]fwc.Fact, 0)
	j := 0
	for j < len(categories) {
		cat := categories[j]
		fmt.Println("Describe its", cat)
		possibleChoice := make(map[string]struct {
			Valid bool
			Fact  fwc.Fact
		})
		for i, fact := range factList[cat] {
			possibleChoice[fmt.Sprint(i+1)] = struct {
				Valid bool
				Fact  fwc.Fact
			}{
				Valid: true,
				Fact:  fact,
			}
			fmt.Printf("\u2003%v. %v\n", i+1, fact)
		}
		fmt.Print("Input your choice (ex: 1): ")
		choice, err := userio.ReadInput()
		if err != nil {
			panic(err)
		}
		if possibleChoice[choice].Valid {
			j++
			knownFacts = append(knownFacts, possibleChoice[choice].Fact)
		} else {
			fmt.Println("Invalid option, Try again.")
		}
	}
	mostConfident := fwc.InferenceScore{}
	catSpecies := fwc.ForwardChainingProbabilist(knownFacts, fwc.Cat)
	for _, prob := range catSpecies {
		if mostConfident.Score < prob.Score {
			mostConfident = prob
		}
		fmt.Printf("%s got Score: %v%% \n", prob.Conclusion, int((prob.Score/9.0)*100.0))
	}
	fmt.Println("It is likely that the cat you mean is:", mostConfident.Conclusion)
}
