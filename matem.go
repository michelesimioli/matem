package matem

// Doppio ritorna il doppio di un intero
func Doppio(a int) int {
	return 2 * a
}

// Triplo ritorna il triplo di un intero
func Triplo(a int) int {
	return 3 * a
}

// Quadruplo ritorna il triplo di un intero
func Quadruplo(a int) int {
	return 4 * a
}

// Massimo : il numero massimo trattato
const Massimo int = 100

// Meta ritorna la metà di un intero come float32
// purchè l'intero sia minore del Massimo
func Meta(a int) float32 {
	if check(a) {
		return float32(a) / 2.0
	}
	return float32(0.0)
}

func check(x int) bool {
	if x <= Massimo {
		return true
	}
	return false
}
