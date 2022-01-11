package calculadoraDos

import (
	"testing"

	"github.com/go-playground/assert"
)

// se crea un un struct fakeConfig que implemente una lógica en la que sólo habilita la suma al cliente "John Doe"
type fakeConfig struct{}

func (f *fakeConfig) SumaEnabled(cliente string) bool {
	return cliente == "John 33"
}
func TestSumarRestrictedFake(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John 33"
	//cliente_dos := "Mister Pmosh"
	resultadoEsperado := 8
	//resultadoEsperadoError := -99999
	// Se genera el objeto fake a usar
	myFake := &fakeConfig{}
	// Se ejecuta el test y Se valida que para el cliente autorizado devuelva el resultado correcto de la suma y que para el cliente no autorizado devuelva el número -99999
	resultado := SumarRestricted(num1, num2, myFake, cliente)
	assert.Equal(t, resultadoEsperado, resultado)
	//resultado2 := SumarRestricted(num1, num2, myFake, cliente_dos)
	//assert.Equal(t, resultadoEsperadoError, resultado2)
}
