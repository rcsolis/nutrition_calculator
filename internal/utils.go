package internal

func getPointsFromRange(value float64, steps []float64) int {
	lenSteps := len(steps)
	for i, v := range steps {
		if value > v {
			return lenSteps - i
		}
	}
	return 0
}

func getEnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func getSodiumFromSalt(salt float64) SodiumMilligram {
	return SodiumMilligram(salt / 2.54)
}
