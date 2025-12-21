package services

import (
	"time"
	"go-server/internal/models"
)

type PageService struct{}

func NewPageService() *PageService {
	return &PageService{}
}

func (s *PageService) GetHomeData() *models.HomePageData {
	return &models.HomePageData{
		PageData: models.PageData{
			Title:       "Home - GoNexus",
			Description: "Welcome to our amazing Go web application",
			CurrentPage: "home",
			CurrentYear: time.Now().Year(),
			ContentType:  "home",
		},
		HeroTitle:    "Welcome to Go Web Development",
		HeroSubtitle: "Building fast, scalable web applications with Go",
		Features: []models.Feature{
			{
				Title:       "Lightning Fast",
				Description: "Go's performance and concurrency make your apps incredibly fast",
				Icon:        "⚡",
			},
			{
				Title:       "Simple & Clean",
				Description: "Clean syntax and minimal dependencies for maintainable code",
				Icon:        "✨",
			},
			{
				Title:       "Production Ready",
				Description: "Built-in testing, logging, and deployment tools",
				Icon:        "🚀",
			},
		},
	}
}

func (s *PageService) GetAboutData() *models.AboutPageData {
	return &models.AboutPageData{
		PageData: models.PageData{
			Title:       "About Us - GoNexus",
			Description: "Learn more about our company and team",
			CurrentPage: "about",
			CurrentYear: time.Now().Year(),
			ContentType:  "about",
		},
		CompanyName:  "Go Web Solutions",
		CompanyStory: "We are passionate about building high-performance web applications using Go. Our team combines expertise in modern web development with deep knowledge of Go's ecosystem to deliver exceptional solutions.",
		FoundedYear:  2023,
		Mission:      "To make Go web development accessible and enjoyable for developers worldwide.",
		TeamMembers: []models.TeamMember{
			{
				Name:     "Alex Johnson",
				Role:     "CEO & Founder",
				Bio:      "12+ years in web development, Go enthusiast since 2018",
				ImageURL: "/static/images/default-avatar.svg",
			},
			{
				Name:     "Sarah Chen",
				Role:     "Lead Developer",
				Bio:      "Full-stack developer specializing in Go and cloud architecture",
				ImageURL: "/static/images/default-avatar.svg",
			},
			{
				Name:     "Mike Rodriguez",
				Role:     "DevOps Engineer",
				Bio:      "Containerization expert and CI/CD pipeline specialist",
				ImageURL: "/static/images/default-avatar.svg",
			},
		},
	}
}

func (s *PageService) GetTestimonialData() *models.TestimonialPageData {
	return &models.TestimonialPageData{
		PageData: models.PageData{
			Title:       "Testimonials - GoNexus",
			Description: "What our clients say about our Go web development services",
			CurrentPage: "testimonials",
			CurrentYear: time.Now().Year(),
			ContentType:  "testimonials",
		},
		Testimonials: []models.Testimonial{
			{
				Name:    "David Kim",
				Role:    "CTO",
				Company: "TechStart Inc.",
				Content: "The performance improvement we saw after migrating to Go was incredible. Page loads dropped by 60% and server costs went down significantly.",
				Rating:  5,
				ImageURL: "/static/images/default-avatar.svg",
			},
			{
				Name:    "Emma Thompson",
				Role:    "Engineering Manager",
				Company: "Digital Solutions Ltd.",
				Content: "The team's expertise in Go and clean architecture patterns helped us build a scalable system that handles millions of requests daily.",
				Rating:  5,
				ImageURL: "/static/images/default-avatar.svg",
			},
			{
				Name:    "Carlos Martinez",
				Role:    "Founder",
				Company: "StartupHub",
				Content: "From concept to production in record time. The Go-based solution they built is both fast and maintainable.",
				Rating:  5,
				ImageURL: "/static/images/default-avatar.svg",
			},
		},
		Stats: []models.Stat{
			{Number: "50+", Label: "Projects Completed"},
			{Number: "99.9%", Label: "Uptime"},
			{Number: "60%", Label: "Avg Performance Improvement"},
			{Number: "24/7", Label: "Support Available"},
		},
	}
}

func (s *PageService) GetCoursesData(category string) *models.CoursesPageData {
	allCourses := []models.Course{
		{
			Title:       "Go Fundamentals",
			Description: "Master the basics of Go programming language from scratch. Learn syntax, data types, functions, and more.",
			Instructor:  "Alex Johnson",
			Duration:    "12 hours",
			Level:       "Beginner",
			Price:       "$49",
			ImageURL:    "https://images.unsplash.com/photo-1515879218367-8466d910aaa4?w=400&h=250&fit=crop",
			Category:    "Go",
			Students:    1245,
		},
		{
			Title:       "Advanced Go Patterns",
			Description: "Deep dive into advanced Go patterns, concurrency, channels, and best practices for production code.",
			Instructor:  "Sarah Chen",
			Duration:    "18 hours",
			Level:       "Advanced",
			Price:       "$89",
			ImageURL:    "https://images.unsplash.com/photo-1516116216624-53e697fedbea?w=400&h=250&fit=crop",
			Category:    "Go",
			Students:    876,
		},
		{
			Title:       "Building REST APIs with Go",
			Description: "Learn to build production-ready RESTful APIs using Go, Chi router, and PostgreSQL.",
			Instructor:  "Mike Rodriguez",
			Duration:    "15 hours",
			Level:       "Intermediate",
			Price:       "$69",
			ImageURL:    "https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=400&h=250&fit=crop",
			Category:    "Backend",
			Students:    1532,
		},
		{
			Title:       "Docker & Kubernetes for Go Devs",
			Description: "Containerize your Go applications and deploy them to Kubernetes clusters with confidence.",
			Instructor:  "Alex Johnson",
			Duration:    "20 hours",
			Level:       "Intermediate",
			Price:       "$79",
			ImageURL:    "https://images.unsplash.com/photo-1667372393119-3d4c48d07fc9?w=400&h=250&fit=crop",
			Category:    "DevOps",
			Students:    943,
		},
		{
			Title:       "Full-Stack Go with HTMX",
			Description: "Build modern, interactive web applications using Go templates and HTMX without JavaScript.",
			Instructor:  "Sarah Chen",
			Duration:    "14 hours",
			Level:       "Intermediate",
			Price:       "$59",
			ImageURL:    "https://images.unsplash.com/photo-1547658719-da2b51169166?w=400&h=250&fit=crop",
			Category:    "Web Development",
			Students:    687,
		},
		{
			Title:       "Microservices with Go",
			Description: "Design and implement microservices architecture using Go, gRPC, and message queues.",
			Instructor:  "Mike Rodriguez",
			Duration:    "22 hours",
			Level:       "Advanced",
			Price:       "$99",
			ImageURL:    "https://images.unsplash.com/photo-1558494949-ef010cbdcc31?w=400&h=250&fit=crop",
			Category:    "Backend",
			Students:    1087,
		},
	}

	var filteredCourses []models.Course
	if category == "" || category == "All" {
		filteredCourses = allCourses
		category = "All"
	} else {
		for _, course := range allCourses {
			if course.Category == category {
				filteredCourses = append(filteredCourses, course)
			}
		}
	}

	return &models.CoursesPageData{
		PageData: models.PageData{
			Title:       "Courses - GoNexus",
			Description: "Explore our comprehensive programming courses to boost your skills",
			CurrentPage: "courses",
			CurrentYear: time.Now().Year(),
			ContentType: "courses",
		},
		Categories:       []string{"All", "Go", "Web Development", "DevOps", "Backend"},
		Courses:          filteredCourses,
		SelectedCategory: category,
	}
}

func (s *PageService) GetAgeCheckData() *models.PageData {
	return &models.PageData{
		Title:       "Age Verification - GoNexus",
		Description: "Please verify your age to access this content.",
		CurrentPage: "age-check",
		CurrentYear: time.Now().Year(),
		ContentType: "age-check",
	}
}

func (s *PageService) GetAccessDeniedData() *models.PageData {
	return &models.PageData{
		Title:       "Access Denied - GoNexus",
		Description: "You must be at least 20 years old to access this content.",
		CurrentPage: "access-denied",
		CurrentYear: time.Now().Year(),
		ContentType: "access-denied",
	}
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Updated: 2025-12-21
 * └─ go-server ───┘
 */