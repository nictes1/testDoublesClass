package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// se crea un un struct spy compuesto por un booleano que nos informará si ocurre el llamado a Log
type spyLogger struct {
	spyCalled bool
}

//  Para espiar creamos un loggerSpy que setea en true spyCalled si entra al método
func (s *spyLogger) Log(string) error {
	s.spyCalled = true
	return nil
}
func TestSumarConSpy(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto spy a usar
	mySpy := &spyLogger{}
	// Se ejecuta el test y se validan el resultado y que spyCalled sea true para dar el test por válido
	resultado := Sumar(num1, num2, mySpy)
	assert.Equal(t, resultadoEsperado, resultado)
	assert.True(t, mySpy.spyCalled)
}
