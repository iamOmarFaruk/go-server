package handlers

import (
	"html/template"
	"net/http"

	"go-server/internal/services"
)

type PageHandler struct {
	pageService *services.PageService
	templates   map[string]*template.Template
}

func NewPageHandler(templatePath string) (*PageHandler, error) {
	// Create a template map for different pages
	templates := make(map[string]*template.Template)
	
	// Parse home template
	homeTmpl, err := template.ParseFiles(templatePath + "/home.html")
	if err != nil {
		return nil, err
	}
	templates["home"] = homeTmpl
	
	// Parse about template
	aboutTmpl, err := template.ParseFiles(templatePath + "/about.html")
	if err != nil {
		return nil, err
	}
	templates["about"] = aboutTmpl
	
	// Parse testimonials template
	testimonialsTmpl, err := template.ParseFiles(templatePath + "/testimonials.html")
	if err != nil {
		return nil, err
	}
	templates["testimonials"] = testimonialsTmpl
	
	// Parse 404 template
	notFoundTmpl, err := template.ParseFiles(templatePath + "/404.html")
	if err != nil {
		return nil, err
	}
	templates["404"] = notFoundTmpl

	return &PageHandler{
		pageService: services.NewPageService(),
		templates:   templates,
	}, nil
}

func (h *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetHomeData()
	err := h.templates["home"].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) About(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetAboutData()
	err := h.templates["about"].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) Testimonials(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetTestimonialData()
	err := h.templates["testimonials"].Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	err := h.templates["404"].Execute(w, nil)
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
}