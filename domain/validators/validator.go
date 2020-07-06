package validators

import (
	"github.com/klassmann/cpfcnpj"
)

// IsValidCPF check CPF is valid
func IsValidCPF(cpf string) bool {
	return cpfcnpj.ValidateCPF(cpf)
}

// IsValidCNPJ check CNPJ is valid
func IsValidCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}

// IsValidBrazilianID check CNPJ/CPF is valid
func IsValidBrazilianID(id string) bool {
	return cpfcnpj.ValidateCNPJ(id) || cpfcnpj.ValidateCPF(id)
}
