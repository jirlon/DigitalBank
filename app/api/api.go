package api

import (
	"context"
	"fmt"
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
	retorno, _ := accountRepo.FindAll(context.TODO())
	fmt.Println(retorno)
	createAccountUC := usecase.NewCreateAccountUseCase(accountRepo)
	listAccountUC := usecase.NewListAccountUseCase(accountRepo)
	getBalanceUC := usecase.NewGetBalanceUseCase(accountRepo)

	createAccountHandler := handler.NewCreateAccountHandler(createAccountUC)
	listAccountHandler := handler.NewListAccountHandler(listAccountUC)
	getBalanceHandler := handler.NewGetBalanceHandler(getBalanceUC)

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)

	accountsRouter := chi.NewRouter()
	accountsRouter.Route("/accounts", func(r chi.Router) {
		r.Post("/", rest.Handle(createAccountHandler.CreateAccount))
		r.Get("/", rest.Handle(listAccountHandler.ListAccounts))
		r.Get("/{account_id}/balance", rest.Handle(getBalanceHandler.GetBalance))
	})

	mainRouter.Mount("/digitalbank/v1", accountsRouter)

	return mainRouter
}
