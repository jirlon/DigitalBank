package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jirlon/digitalbank/app/http/handler"
	"github.com/jirlon/digitalbank/app/http/handler/rest"
	"github.com/jirlon/digitalbank/app/repositories"
	"github.com/jirlon/digitalbank/app/usecase"
)

func NewRouter(dbpool *pgxpool.Pool) http.Handler {
	accountRepo := repositories.New(dbpool)
	createAccountUC := usecase.NewCreateAccountUseCase(accountRepo)
	accountHandler := handler.NewAccountHandler(createAccountUC)

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)

	accountsRouter := chi.NewRouter()

	accountsRouter.Post("/accounts", rest.Handle(accountHandler.CreateAccount))

	mainRouter.Mount("/digitalbank/v1", accountsRouter)

	return mainRouter
}
