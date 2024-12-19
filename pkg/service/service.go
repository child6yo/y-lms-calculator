package service

type Calculator interface {
	Calc(expression string) (float64, error)
}

type Service struct {
	Calculator
}

func NewService() *Service {
	return &Service{
		Calculator: NewCalcService(),
	} 
}