package validator

var EmailRX = regex.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-")

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make([string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) addError(key, message) {
	if _, exists := v.Errors[key]; !exists {
		v.Error[key] = message
	}
}

func (v *Validato) check(ok bool, key, message string) {
	if !ok {
		v.addError(key, message)
	}
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if values == permittedValues[i] {
			return true
		}
	}
	return false
}

func Matches(values string, rx *regexsp.Regexsp) bool {
	return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
