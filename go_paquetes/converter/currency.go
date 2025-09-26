// Package converter proporciona funcionalidades para conversión de monedas
package converter

func Convert(amount float64, curency string) float64 {
	// Mapa que contiene las tasas de cambio respecto al dólar USD
	// Clave: código de moneda (string)
	// Valor: tasa de cambio (float64)
	worth := map[string]float64{
		"EUR": 0.85, "GBP": 0.73, "KRW": 1100.0, "BTC": 0.000023,
	}

	// Obtener la tasa de cambio para la moneda especificada
	// worths: contiene la tasa de cambio si existe
	// exist: booleano que indica si la moneda existe en el mapa

	worths, exist := worth[curency]
	// Verificar si la moneda solicitada existe en nuestras conversiones
	if !exist {
		// Si la moneda no existe, retornar 0 como indicador de error
		return 0
	}
	// Calcular y retornar el valor convertido
	// Fórmula: cantidad en USD * tasa de cambio
	return amount * worths
}

// Retornar un slice con los códigos de todas las monedas disponibles
// El orden en el slice define el orden en que se mostrarán las monedas
func GetCurrencies() []string {
	return []string{"EUR", "GBP", "KRW", "BTC"}
}
