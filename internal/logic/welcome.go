package logic

type Appeal struct {
	// сюда прийдет уровень (интерфейсный) БД
}

func (a Appeal) Welcome() string {
	return "Welcome to the ATM!"
}

func New() *Appeal {
	return &Appeal{}
}
