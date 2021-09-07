package rhelper

import "math"

type DiscType string

const (
	DiscTypeNominal    DiscType = "nominal"
	DiscTypePercentage DiscType = "percentage"
)

func CalculateDiscount(nom, discValue float64, types DiscType) (disc, grandtotal float64) {
	if types == DiscTypePercentage {
		disc = CalculateDiscountPercentage(nom, discValue)
		grandtotal = nom - disc
		return
	} else if types == DiscTypeNominal {
		disc = math.Min(discValue, nom)
		grandtotal = nom - disc
		return
	}

	grandtotal = nom
	return
}

// CalculatePercentage will return nominal from nom * percentage / 100
// when disc < 0 what happend?
// when nominal or discount is inf -1 then return 0
// when nominal or discount is nan then return 0
func CalculateDiscountPercentage(nom, ptg float64) (res float64) {
	ptg = math.Max(ptg, 0)
	ptg = math.Min(ptg, 100)

	nom = math.Max(nom, 0)

	return CalculatePercentage(nom, ptg)
}

func CalculatePercentage(nom, ptg float64) (res float64) {
	if math.IsInf(nom, -1) || math.IsInf(ptg, -1) {
		return
	} else if math.IsNaN(nom) || math.IsNaN(ptg) {
		return
	}

	res = nom * ptg / 100
	return
}
