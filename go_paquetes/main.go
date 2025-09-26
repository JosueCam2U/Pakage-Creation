package main

import (
	"convertidor/converter" // Importa el paquete personalizado converter
	"fmt"
)

func main() {

	usd := 0.0
	fmt.Print("Ingrese el valor que quiera convertitr desde dolares: ")
	fmt.Scan(&usd)

	// Iterar sobre cada moneda disponible para realizar la conversión
	// _: ignora el índice del slice (no se necesita)
	// currency: cada código de moneda en el slice
	for _, currency := range converter.GetCurrencies() {
		result := converter.Convert(usd, currency)
		fmt.Println("usted tiene: ", usd, result, currency)
	}

	txt := ""
	fmt.Print("Ingrese un texto")
	fmt.Scan(&txt)

	// Analizar el texto para contar la frecuencia de cada vocal individual
	// CountVowels() retorna un mapa donde:
	// - Clave: vocal ("a", "e", "i", "o", "u")
	// - Valor: cantidad de veces que aparece esa vocal
	vowels := converter.CountVowels(txt)
	total := converter.TotalVowels(txt)

	fmt.Println("Texto :", txt)

	// Iterar sobre el mapa de vocales para mostrar el conteo individual de cada una
	// vowel: cada vocal (clave del mapa)
	// count: cantidad de veces que aparece esa vocal (valor del mapa)
	for vowel, count := range vowels {
		fmt.Println("Vocales :", vowel, count)
	}
	fmt.Println("Total de vocales: ", total)
}
