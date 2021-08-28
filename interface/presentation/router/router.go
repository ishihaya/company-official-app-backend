package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ishihaya/company-official-app-backend/di"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/middleware"
	_ "github.com/ishihaya/company-official-app-backend/pkg/swagger" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/ishihaya/company-official-app-backend/pkg/factory"
)

type Router struct {
	chi.Router
}

func New() *Router {
	r := chi.NewRouter()
	return &Router{r}
}

func (r *Router) HealthCheck() {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		factory.JSON(w, http.StatusOK, "ok")
	})
}

func (r *Router) Swagger() {
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}

func (r *Router) Routes() {
	authMiddleware := di.InitAuth()
	userHandler := di.InitUser()

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.AuthAPI)
		r.Use(middleware.CurrentTime)
		r.Route("/user", func(r chi.Router) {
			r.Get("/", userHandler.Get)
			r.Post("/", userHandler.Create)
		})
	})

}

func (r *Router) RunServer(port int) {
	log.Printf("Listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}
