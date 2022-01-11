package calculadoraDos

import (
	"testing"

	"github.com/go-playground/assert"
)

// se crea un un struct mockConfig
type mockConfig struct {
	clienteUsado string
}

// El mock debe implementar el método necesario y comprobar que SumaEnabled sea llamado y que se haga exactamente con el mismo cliente que recibió SumarRestricted
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}
func TestSumarRestricted(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "Nico Crack"
	resultadoEsperado := 8
	// Se genera el objeto mock a usar para satisfacer la necesidad de la función Sumar
	myMock := &mockConfig{}
	// Se ejecuta el test y se valida el resultado y que el mock haya registrado la información correcta
	resultado := SumarRestricted(num1, num2, myMock, cliente)
	assert.Equal(t, resultadoEsperado, resultado)
	assert.Equal(t, cliente, myMock.clienteUsado)
}
