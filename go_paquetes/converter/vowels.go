// Package converter proporciona utilidades para análisis de texto, incluyendo conteo de vocal
package converter

import "strings"

// CountVowels analiza un texto y cuenta la frecuencia de cada vocal individual
//map[string]int: mapa donde las claves son las vocales
//y los valores son la cantidad de veces que aparece cada vocal
func CountVowels(text string) map[string]int {
	vowels := map[string]int{
		"a": 0, "e": 0, "i": 0, "o": 0, "u": 0,
	}
	// Mapa inicializado con todas las vocales y contadores en cero
	// Esta estructura garantiza que siempre se retornen todas las vocales,
	// incluso si alguna no aparece en el texto

	// Convertir todo el texto a minúsculas para hacer el análisis
	lowerTxt := strings.ToLower(text)

	// Iterar sobre cada carácter del texto convertido a minúsculas
	// _: ignora el índice (no lo necesitamos)
	// char: cada carácter individual del texto como rune (permite manejar Unicode)

	for _, char := range lowerTxt {
		charStr := string(char)

		// Verificar si el carácter actual es una vocal
		// exists: booleano que indica si el carácter está en el mapa de vocales
		if _, exists := vowels[charStr]; exists {
			vowels[charStr]++
		}
	}
	return vowels
}

// TotalVowels cuenta el número total de vocales en un texto, sin distinguir entre tipos de vocal
func TotalVowels(text string) int {
	count := 0
	vowels := "aeiou"
	lowerTxt := strings.ToLower(text)

	// Iterar sobre cada carácter del texto
	for _, char := range lowerTxt {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
