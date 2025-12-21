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
			Title:       "Home - Go Web App",
			Description: "Welcome to our amazing Go web application",
			CurrentPage: "home",
			CurrentYear: time.Now().Year(),
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
			Title:       "About Us - Go Web App",
			Description: "Learn more about our company and team",
			CurrentPage: "about",
			CurrentYear: time.Now().Year(),
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
				ImageURL: "/static/images/team/alex.jpg",
			},
			{
				Name:     "Sarah Chen",
				Role:     "Lead Developer",
				Bio:      "Full-stack developer specializing in Go and cloud architecture",
				ImageURL: "/static/images/team/sarah.jpg",
			},
			{
				Name:     "Mike Rodriguez",
				Role:     "DevOps Engineer",
				Bio:      "Containerization expert and CI/CD pipeline specialist",
				ImageURL: "/static/images/team/mike.jpg",
			},
		},
	}
}

func (s *PageService) GetTestimonialData() *models.TestimonialPageData {
	return &models.TestimonialPageData{
		PageData: models.PageData{
			Title:       "Testimonials - Go Web App",
			Description: "What our clients say about our Go web development services",
			CurrentPage: "testimonials",
			CurrentYear: time.Now().Year(),
		},
		Testimonials: []models.Testimonial{
			{
				Name:    "David Kim",
				Role:    "CTO",
				Company: "TechStart Inc.",
				Content: "The performance improvement we saw after migrating to Go was incredible. Page loads dropped by 60% and server costs went down significantly.",
				Rating:  5,
				ImageURL: "/static/images/testimonials/david.jpg",
			},
			{
				Name:    "Emma Thompson",
				Role:    "Engineering Manager",
				Company: "Digital Solutions Ltd.",
				Content: "The team's expertise in Go and clean architecture patterns helped us build a scalable system that handles millions of requests daily.",
				Rating:  5,
				ImageURL: "/static/images/testimonials/emma.jpg",
			},
			{
				Name:    "Carlos Martinez",
				Role:    "Founder",
				Company: "StartupHub",
				Content: "From concept to production in record time. The Go-based solution they built is both fast and maintainable.",
				Rating:  5,
				ImageURL: "/static/images/testimonials/carlos.jpg",
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