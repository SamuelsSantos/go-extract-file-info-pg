package validators

import (
	"github.com/klassmann/cpfcnpj"
)

// Validator struct handle validations
type Validator struct{}

// Validation interface handle validations
type Validation interface {
	IsValidCPF(cpf string) bool
	IsValidCNPJ(cnpj string) bool
	IsValidBrazilianID(cnpj string) bool
}

// IsValidCPF check CPF is valid
func (v *Validator) IsValidCPF(cpf string) bool {
	return cpfcnpj.ValidateCPF(cpf)
}

// IsValidCNPJ check CNPJ is valid
func (v *Validator) IsValidCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}

// IsValidBrazilianID check CNPJ/CPF is valid
func (v *Validator) IsValidBrazilianID(id string) bool {
	return cpfcnpj.ValidateCNPJ(id) || cpfcnpj.ValidateCPF(id)
}
