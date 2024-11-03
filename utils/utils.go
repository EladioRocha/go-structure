package utils

import (
	"math"
	"math/big"
)

// roundTo redondea un *big.Float a la cantidad especificada de decimales
func roundTo(value *big.Float, decimals uint32) *big.Float {
	// Crear el multiplicador 10^decimals
	multiplier := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))

	// Multiplicar el valor por el multiplicador
	temp := new(big.Float).Mul(value, multiplier)

	// Redondear usando SetPrec y dividir de nuevo por el multiplicador
	rounded := new(big.Float).SetPrec(64).Copy(temp).SetMode(big.ToNearestEven)
	rounded.Quo(rounded, multiplier)

	return rounded
}