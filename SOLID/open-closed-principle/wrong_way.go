package openclosedprinciple

import "reflect"

type Funcionarios struct{}
type Estagiario struct{}

func FolhaDePagamentos[TipoFuncionario Funcionarios | Estagiario](funcionario TipoFuncionario) float64 {

	if reflect.TypeOf(Funcionarios{}) == reflect.TypeOf(funcionario) {
		return 0.0
	} else {
		return 0.0
	}
}
