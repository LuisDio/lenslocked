package main

import (
	"fmt"
	"net/http"

	"github.com/LuisDio/lenslocked/controllers"
	"github.com/LuisDio/lenslocked/templates"
	"github.com/LuisDio/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
