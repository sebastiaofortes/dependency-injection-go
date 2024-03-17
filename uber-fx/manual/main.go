package main

import "fmt"

// Service - Um exemplo de serviço que nossa aplicação irá utilizar
type Service struct {
}

// DoSomething - Método exemplo de Service
func (s *Service) DoSomething() {
	fmt.Println("Service: Fazendo algo importante!")
}

// NewService - Construtor do serviço
func NewService() *Service {
	return &Service{}
}

// Controller - Controlador que depende de Service
type Controller struct {
	service *Service
}

// NewController - Construtor para Controller que necessita de uma referência de Service
func NewController(s *Service) *Controller {
	return &Controller{service: s}
}

// ExecuteAction - Método exemplo de Controller
func (c *Controller) ExecuteAction() {
	fmt.Println("Controller: Preparando para executar uma ação")
	c.service.DoSomething()
}

// StartApp - Função que inicializa a aplicação
func StartApp(c *Controller) {
	fmt.Println("Aplicação iniciada!")
	c.ExecuteAction()
}

func main() {
	myService := NewService()
	myController := NewController(myService)

	StartApp(myController)
}