package useCase

type TemplateStruct struct {
	Name        string
	Surname     string
	Age         int
	CallbackUrl string
}

func NewTemplateStruct(name, surname string, age int, callbackUrl string) TemplateStruct {
	return TemplateStruct{
		Name:        name,
		Surname:     surname,
		Age:         age,
		CallbackUrl: callbackUrl,
	}
}
