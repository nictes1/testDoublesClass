package calculadoraDos

type Config interface {
	SumaEnabled(cliente string) bool
}

// Función que recibe dos enteros y retorna la suma resultante
func SumarRestricted(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2
}
