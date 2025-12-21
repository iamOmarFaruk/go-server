package handlers

import (
	"html/template"
	"net/http"

	"go-server/internal/services"
)

type PageHandler struct {
	pageService *services.PageService
	templates   *template.Template
}

func NewPageHandler(templatePath string) (*PageHandler, error) {
	// Parse all templates - layout contains nav and footer includes
	templates, err := template.ParseGlob(templatePath + "/*.html")
	if err != nil {
		return nil, err
	}

	return &PageHandler{
		pageService: services.NewPageService(),
		templates:   templates,
	}, nil
}

func (h *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	// Execute layout template with home page data
	data := h.pageService.GetHomeData()
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) About(w http.ResponseWriter, r *http.Request) {
	// Execute layout template with about page data
	data := h.pageService.GetAboutData()
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) Testimonials(w http.ResponseWriter, r *http.Request) {
	// Execute layout template with testimonials page data
	data := h.pageService.GetTestimonialData()
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) Courses(w http.ResponseWriter, r *http.Request) {
	// Execute layout template with courses page data
	data := h.pageService.GetCoursesData()
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	
	// Execute layout template with 404 page data
	data := h.pageService.GetHomeData()
	data.Title = "Page Not Found - 404 - Go Web App"
	data.Description = "Oops! The page you're looking for doesn't exist."
	data.CurrentPage = "404"
	data.ContentType = "404"
	
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Touched: 2025-12-21
 * └─ go-server ───┘
 */