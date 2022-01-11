package calculadora

import "testing"

// se crea un un struct dummyLogger
type dummyLogger struct{}

//  Se escriben las funciones necesarios para que dummyLogger cumpla con la interfaz que va a reemplazar (Logger)
func (d *dummyLogger) Log(string) error {
	return nil
}

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la función Sumar
	myDummy := &dummyLogger{}
	// Se ejecuta el test
	resultado := Sumar(num1, num2, myDummy)
	// Se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
