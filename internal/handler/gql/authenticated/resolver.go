package authenticated

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/service"

	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service service.ServiceRegistry
	logger  *zap.Logger
}

// NewSchema creates an ExecutableSchema instance.
func NewSchema(client *ent.Client, validator *validator.Validate, validationTranslator ut.Translator, logger *zap.Logger) graphql.ExecutableSchema {
	service := service.New(client, logger)

	config := Config{
		Resolvers: &Resolver{service: service, logger: logger},
	}
	config.Directives.Validation = validationResolver(validator, validationTranslator, logger)
	return NewExecutableSchema(config)
}
