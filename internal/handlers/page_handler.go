package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func (h *PageHandler) AgeCheck(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetAgeCheckData()
	err := h.templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PageHandler) VerifyAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/age-check", http.StatusSeeOther)
		return
	}

	// Sanitize and validate input
	ageStr := strings.TrimSpace(r.FormValue("age"))
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		// Invalid input, redirect back to age check
		http.Redirect(w, r, "/age-check", http.StatusSeeOther)
		return
	}

	if age >= 20 {
		// Set cookie
		cookie := &http.Cookie{
			Name:     "age_verified",
			Value:    "true",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/courses", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/access-denied", http.StatusSeeOther)
	}
}

func (h *PageHandler) AccessDenied(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetAccessDeniedData()
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