package main

import (
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/sviut/photo-locker/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4"
	"github.com/sviut/photo-locker/controllers"
	"github.com/sviut/photo-locker/templates"
	"github.com/sviut/photo-locker/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml",
	))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{DB: db}
	sessionService := models.SessionService{DB: db}

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	csrfKey := "UVbMZYGb26WI9OD3n2myJwstbR+Dsk1xusq5qB96Uk4="
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: fix before deploy
		csrf.Secure(false),
	)

	fmt.Println("Listening on port :3333...")
	_ = http.ListenAndServe(":3333", csrfMw(r))
}
