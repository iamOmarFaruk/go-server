package handlers

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go-server/internal/models"
	"go-server/internal/services"
)

type PageHandler struct {
	pageService *services.PageService
	templates   *template.Template
	isDev       bool
}

func NewPageHandler(templatePath string) (*PageHandler, error) {
	// Parse layout and partials
	templates, err := template.ParseGlob(templatePath + "/layout.html")
	if err != nil {
		return nil, err
	}
	// Parse partials
	_, err = templates.ParseGlob(templatePath + "/partials/*.html")
	if err != nil {
		return nil, err
	}

	isDev := os.Getenv("DEV_MODE") == "true"

	return &PageHandler{
		pageService: services.NewPageService(),
		templates:   templates,
		isDev:       isDev,
	}, nil
}

func (h *PageHandler) render(w http.ResponseWriter, pageFile string, data interface{}) {
	var tmpl *template.Template
	var err error

	if h.isDev {
		// Re-parse everything in dev mode
		tmpl, err = template.ParseFiles("templates/layout.html")
		if err != nil {
			http.Error(w, "Layout Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = tmpl.ParseGlob("templates/partials/*.html")
		if err != nil {
			http.Error(w, "Partials Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Clone the base templates (layout + partials)
		tmpl, err = h.templates.Clone()
		if err != nil {
			http.Error(w, "Template Clone Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Parse the specific page
	_, err = tmpl.ParseFiles("templates/pages/" + pageFile)
	if err != nil {
		http.Error(w, "Page Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Inject IsDev into data if it's a pointer to struct that has it, 
    // or wrap it. For simplicity, we assume models have IsDev or we handle it.
    // The previous code did `data.IsDev = h.isDev`. To keep that:
    if pageData, ok := data.(*models.PageData); ok {
        pageData.IsDev = h.isDev
    } else if homeData, ok := data.(*models.HomePageData); ok {
        homeData.IsDev = h.isDev
    } else if coursesData, ok := data.(*models.CoursesPageData); ok {
        coursesData.IsDev = h.isDev
    } else if aboutData, ok := data.(*models.AboutPageData); ok {
        aboutData.IsDev = h.isDev
    } else if testData, ok := data.(*models.TestimonialPageData); ok {
        testData.IsDev = h.isDev
    }

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *PageHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetHomeData()
	h.render(w, "home.html", &data)
}

func (h *PageHandler) About(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetAboutData()
	h.render(w, "about.html", &data)
}

func (h *PageHandler) Testimonials(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetTestimonialData()
	h.render(w, "testimonials.html", &data)
}

func (h *PageHandler) Courses(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	data := h.pageService.GetCoursesData(category)
	h.render(w, "courses.html", &data)
}

func (h *PageHandler) AgeCheck(w http.ResponseWriter, r *http.Request) {
	data := h.pageService.GetAgeCheckData()
	h.render(w, "age-check.html", data)
}

func (h *PageHandler) VerifyAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/age-check", http.StatusSeeOther)
		return
	}

	ageStr := strings.TrimSpace(r.FormValue("age"))
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Redirect(w, r, "/age-check", http.StatusSeeOther)
		return
	}

	if age >= 20 {
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
	h.render(w, "access-denied.html", data)
}

func (h *PageHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	
	data := h.pageService.GetHomeData()
	data.Title = "Page Not Found - 404 - GoNexus"
	data.Description = "Oops! The page you're looking for doesn't exist."
	data.CurrentPage = "404"
	data.ContentType = "404"
	
	h.render(w, "404.html", &data.PageData)
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Touched: 2025-12-21
 * │ Updated: 2025-12-21
 * └─ go-server ───┘
 */