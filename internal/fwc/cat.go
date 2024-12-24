package fwc

var persian Rule = Rule{
	Conditions: []Fact{Gentle, LaidBack, RoundEyes, SmallEars, BushyTail, LargeSize, RoundFace, HighMaintenance, Vocal},
	Conclusion: "Persian",
}

var siamese Rule = Rule{
	Conditions: []Fact{AlmondEyes, ColorPoints, LargeEars, LongTail, SleekBody, Vocal, Active, ModerateMaintenance, Vocal},
	Conclusion: "Siamese",
}

var maineCoon Rule = Rule{
	Conditions: []Fact{TuftedEars, LongShaggyCoat, SmallEars, Friendly, LongTail, MuscularBody, Gentle, HighMaintenance, Intelligence},
	Conclusion: "Maine Coon",
}

var bengal Rule = Rule{
	Conditions: []Fact{SpottedCoat, MarbledCoat, AthleticBuild, HighlyEnergetic, Curious, LargeEars, LongTail, SleekBody, LowMaintenance},
	Conclusion: "Bengal",
}

var sphynx Rule = Rule{
	Conditions: []Fact{Hairless, WrinkledSkin, LargeEars, MuscularBody, AlmondEyes, Vocal, Social, HighMaintenance, Quiet},
	Conclusion: "Sphynx",
}

var ragdoll Rule = Rule{
	Conditions: []Fact{SemiLongHair, Docile, LovesBeingHeld, BlueEyes, LongTail, LargeSize, Friendly, ModerateMaintenance, Quiet},
	Conclusion: "Ragdoll",
}

var scottishFold Rule = Rule{
	Conditions: []Fact{FoldedEars, RoundFace, Adaptable, RoundEyes, ShortTail, SleekBody, Gentle, ModerateMaintenance, Quiet},
	Conclusion: "Scottish Fold",
}

var abyssinian Rule = Rule{
	Conditions: []Fact{TickedTabbyCoat, Active, AlmondEyes, LargeEars, LongTail, SleekBody, Curious, LowMaintenance, Quiet},
	Conclusion: "Abyssinian",
}

var burmese Rule = Rule{
	Conditions: []Fact{ShinyCoat, PeopleOriented, RoundEyes, LargeEars, ShortTail, SleekBody, Affectionate, LowMaintenance, Quiet},
	Conclusion: "Burmese",
}

var lion Rule = Rule{
	Conditions: []Fact{Mane, LivesInPrides, RoundEyes, LargeEars, LongTail, MuscularBody, Gentle, Vocal, Intelligence},
	Conclusion: "Lion",
}

var tiger Rule = Rule{
	Conditions: []Fact{StripedCoat, LargeSize, AlmondEyes, LargeEars, LongTail, MuscularBody, Solitary, HighMaintenance, Intelligence},
	Conclusion: "Tiger",
}

var leopard Rule = Rule{
	Conditions: []Fact{SpottedCoat, Solitary, ExcellentClimber, RoundEyes, LargeEars, LongTail, SleekBody, ModerateMaintenance, Quiet},
	Conclusion: "Leopard",
}

var cheetah Rule = Rule{
	Conditions: []Fact{SlimBuild, TearMarks, FastestLandAnimal, AlmondEyes, LargeEars, LongTail, Active, LowMaintenance, Quiet},
	Conclusion: "Cheetah",
}

var snowLeopard Rule = Rule{
	Conditions: []Fact{ThickFur, LongTail, ColdAdapted, AlmondEyes, LargeEars, MuscularBody, Solitary, ModerateMaintenance, Intelligence},
	Conclusion: "Snow Leopard",
}

var bobcat Rule = Rule{
	Conditions: []Fact{ShortCoat, ShortTail, TuftedEars, SpottedCoat, AlmondEyes, MuscularBody, Active, LowMaintenance, Vocal},
	Conclusion: "Bobcat",
}

var serval Rule = Rule{
	Conditions: []Fact{LongLegs, SpottedCoat, LargeEars, AlmondEyes, LongTail, SleekBody, Curious, ModerateMaintenance, Intelligence},
	Conclusion: "Serval",
}

var caracal Rule = Rule{
	Conditions: []Fact{ReddishCoat, TuftedEars, AthleticBuild, AlmondEyes, LargeEars, LongTail, Solitary, LowMaintenance, Vocal},
	Conclusion: "Caracal",
}

var lynx Rule = Rule{
	Conditions: []Fact{ShortCoat, ShortTail, TuftedEars, ThickFur, RoundEyes, MuscularBody, Solitary, ModerateMaintenance, Vocal},
	Conclusion: "Lynx",
}

var Cat []Rule = []Rule{
	persian,
	siamese,
	maineCoon,
	bengal,
	sphynx,
	ragdoll,
	scottishFold,
	abyssinian,
	burmese,
	lion,
	tiger,
	leopard,
	cheetah,
	snowLeopard,
	bobcat,
	serval,
	caracal,
	lynx,
}
