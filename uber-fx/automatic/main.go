package main

import (
	"fmt"
	"go.uber.org/fx"
)

// Service - Definindo a estrutura do serviço
type Service struct{}

// DoSomething - Ação do serviço
func (s *Service) DoSomething() {
	fmt.Println("Service: Fazendo algo importante!")
}

// NewService - Construtor do serviço para injeção de dependência pelo fx
func NewService() *Service {
	return &Service{}
}

// Controller - Estrutura do controlador que depende do Service
type Controller struct {
	service *Service
}

// NewController - Construtor do controlador para injeção de dependência pelo fx
func NewController(s *Service) *Controller {
	return &Controller{service: s}
}

// ExecuteAction - Ação controlador chamando serviço
func (c *Controller) ExecuteAction() {
	fmt.Println("Controller: Preparando para executar uma ação")
	c.service.DoSomething()
}

// StartApp - Função envolvida por fx.Invoke para iniciar a aplicação
func StartApp(c *Controller) {
	fmt.Println("Aplicação iniciada!")
	c.ExecuteAction()
}

func main() {
	// Construindo e iniciando a aplicação usando uber fx
	fx.New(
		// Provedores de dependência (Construtores que fx vai chamar para injetar dependências)
		fx.Provide(
			NewService,
			NewController,
		),

		// Inicializador da aplicação, fx.Invoke pode ser usado para chamar funções que iniciam a aplicação
		fx.Invoke(StartApp),
	).Run()
}