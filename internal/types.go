package internal

// Define constants for the nutritional data points
var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyBeverageLevels = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarBeverageLevels = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

// Define types for the nutritional data as aliases of float64
type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcidGram float64
type SodiumMilligram float64
type FruitGram float64
type FibreGram float64
type ProteinGram float64

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(e), energyBeverageLevels)
	}
	return getPointsFromRange(float64(e), energyLevels)
}

func (e SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(e), sugarBeverageLevels)
	}
	return getPointsFromRange(float64(e), sugarLevels)
}

func (e SaturatedFattyAcidGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(e), saturatedFattyAcidLevels)
}

func (e SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(e), sodiumLevels)
}

func (e FruitGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		if e > 80 {
			return 10
		} else if e > 60 {
			return 4
		} else if e > 40 {
			return 2
		}
		return 0
	}

	if e > 80 {
		return 5
	} else if e > 60 {
		return 2
	} else if e > 40 {
		return 1
	}
	return 0
}

func (e FibreGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(e), fiberLevels)
}

func (e ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(e), proteinLevels)
}

// Define a type for the nutritional score as an alias of int
type ScoreType int

// Define constants (enum) for the nutritional score
const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

// Define the nutritional score type
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 19, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}

	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

// Define the nutritional data type
type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidGram
	Sodium              SodiumMilligram
	Fruits              FruitGram
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}
