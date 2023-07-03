package mix

import (
	"errors"
	"math"
	"strings"
	"unicode"
)

func Difference(a, b float64, x, y int) float64 {
	diff := a*float64(x) - b*float64(y)
	return math.Abs(diff)
}

func Distance(v1, v2, t float64) float64 {
	return t * (v1 + v2)
}

func DistanceBoat(v1, v2, t1, t2 float64) (float64, error) {
	if v2 > v1 {
		return 0.0, errors.New("oqim juda kuchli, suzma")
	}
	return (v1+v2)*t1 + (v1-v2)*t2, nil
}

func SaleKonfet(kg, narx float64) float64 {
	// res := (kg-10)*narx/2
	if kg > 10 {
		return 10*narx + (kg-10)*narx/2
	}
	return narx * kg
}

func NumberOfVovels(text string) (number int) {

	allowed := "aeiouAEIOU"

	for i := 0; i < len(text); i++ {
		if strings.Contains(allowed, string(text[i])) {
			number++
		}
	}

	return number
}

func OnlyVowels(text string) (vowels string) {
	allowed := "aeiouAEIOU"

	for i := 0; i < len(text); i++ {
		if !strings.Contains(allowed, string(text[i])) && unicode.IsLetter(rune(text[i])) {
			vowels += string(text[i])
		}
	}

	return vowels
}
