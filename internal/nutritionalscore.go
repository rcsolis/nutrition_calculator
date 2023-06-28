package internal

import "fmt"

/**
 * Get the nutritional score
 */
func getNutritionalScore(nd NutritionalData, st ScoreType) NutritionalScore {
	var value int = 0
	var positive int = 0
	var negative int = 0

	if st != Water {
		fruitPoints := nd.Fruits.GetPoints(st)
		fibrePoints := nd.Fibre.GetPoints(st)
		negative = nd.Energy.GetPoints(st) + nd.Sugars.GetPoints(st) + nd.SaturatedFattyAcids.GetPoints(st) + nd.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + nd.Protein.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative > 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

/**
 * Run the application
 */
func Run() string {
	ns := getNutritionalScore(
		NutritionalData{
			Energy:              getEnergyFromKcal(0),
			Sugars:              SugarGram(10),
			SaturatedFattyAcids: SaturatedFattyAcidGram(2),
			Sodium:              getSodiumFromSalt(5),
			Fruits:              FruitGram(60),
			Fibre:               FibreGram(4),
			Protein:             ProteinGram(2),
			IsWater:             false,
		},
		Food)
	letterScore := ns.GetNutriScore()
	return fmt.Sprintf("Value:%d \nPositive: %d \nNegative: %d\nScore: %s", ns.Value, ns.Positive, ns.Negative, letterScore)
}
