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
			Title:          "Wanderlust | Premium Travel Experiences & Web Solutions",
			Description:    "Discover the world's most breathtaking destinations. Your nexus for adventure and technology.",
			Keywords:       "travel, wanderlust, go, golang, web development, courses, adventures, tours",
			SocialImageURL: "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=1200&q=80",
			CurrentPage:    "home",
			CurrentYear:    time.Now().Year(),
			ContentType:    "home",
		},
		HeroSlides: []models.HeroSlide{
			{
				Title:    "Discover The Unseen World",
				Subtitle: "Create unforgettable memories across the globe.",
				CTA1Text: "Explore Destinations",
				CTA1Link: "/stories",
				CTA2Text: "View Packages",
				CTA2Link: "/courses",
				ImageURL: "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=1920&q=80",
				Gradient: "bg-black/40",
			},
			{
				Title:    "Embrace Ancient History",
				Subtitle: "Walk through the echoes of timeless civilizations.",
				CTA1Text: "Start Exploring",
				CTA1Link: "/stories",
				CTA2Text: "Learn More",
				CTA2Link: "/about",
				ImageURL: "https://images.unsplash.com/photo-1548013146-72479768bada?w=1920&q=80",
				Gradient: "bg-black/40",
			},
			{
				Title:    "Your Paradise Awaits",
				Subtitle: "Experience the ultimate luxury of pure relaxation.",
				CTA1Text: "Find Paradise",
				CTA1Link: "/stories",
				CTA2Text: "View Courses",
				CTA2Link: "/courses",
				ImageURL: "https://images.unsplash.com/photo-1506929562872-bb421503ef21?w=1920&q=80",
				Gradient: "bg-black/40",
			},
		},
	Features: []models.Feature{
			{
				Title:       "Pack Like a Pro",
				Description: "Learn the art of minimalist packing. Roll your clothes, use packing cubes, and never pay for baggage fees again.",
				Category:    "Travel Hack",
				Location:    "3 min read",
				Author:      "Sarah Jenkins",
				ImageURL:    "https://images.unsplash.com/photo-1553531384-cc64ac80f931?w=800&q=80",
			},
			{
				Title:       "Hidden Flight Deals",
				Description: "Discover the secrets to finding the cheapest flights. Use incognito mode, flexible dates, and lesser-known search engines.",
				Category:    "Budget",
				Location:    "5 min read",
				Author:      "Mike Chen",
				ImageURL:    "https://images.unsplash.com/photo-1436491865332-7a61a109cc05?w=800&q=80",
			},
			{
				Title:       "Solo Travel Safety",
				Description: "Essential safety tips for solo travelers. From sharing your itinerary to blending in with the locals.",
				Category:    "Safety",
				Location:    "4 min read",
				Author:      "Elena Rodriguez",
				ImageURL:    "https://images.unsplash.com/photo-1488646953014-85cb44e25828?w=800&q=80",
			},
		},
		AboutSection: models.AboutSection{
			Title:    "Wanderlust: More Than Just Travel",
			Subtitle: "Our Mission",
			Description: "We believe that travel is the best way to learn about the world and yourself. Started in 2023, Wanderlust has grown from a small blog to a global community of explorers. Our mission is to make travel accessible, sustainable, and unforgettable for everyone.",
			ImageURL: "https://images.unsplash.com/photo-1522071823991-b9671f99118f?w=1000&q=80",
			CTAText:  "Read Our Story",
			CTALink:  "/about",
		},
		PopularGuides: []models.GuideCard{
			{
				Title:       "The Ultimate Japan Itinerary",
				Description: "Two weeks in the land of the rising sun. From Tokyo's neon lights to Kyoto's ancient temples.",
				ImageURL:    "https://images.unsplash.com/photo-1493976040374-85c8e12f0c0e?w=800&q=80",
				Category:    "Asia",
				Author:      "Yuki Tanaka",
				ReadTime:    "12 min read",
				Link:        "/stories",
			},
			{
				Title:       "Backpacking Europe on a Budget",
				Description: "How to see Paris, Rome, and Berlin without breaking the bank. Tips on hostels, rail passes, and cheap eats.",
				ImageURL:    "https://images.unsplash.com/photo-1471341971476-ae15ff5dd4ea?w=800&q=80",
				Category:    "Europe",
				Author:      "Alex Johnson",
				ReadTime:    "15 min read",
				Link:        "/stories",
			},
			{
				Title:       "Safari 101: A Beginner's Guide",
				Description: "Everything you need to know before booking your first African safari. When to go, what to pack, and where to stay.",
				ImageURL:    "https://images.unsplash.com/photo-1516426122078-c23e76319801?w=800&q=80",
				Category:    "Africa",
				Author:      "Sarah Jenkins",
				ReadTime:    "10 min read",
				Link:        "/stories",
			},
		},
		Testimonials: []models.TestimonialCard{
			{
				Name:     "Emily Clarke",
				Role:     "Digital Nomad",
				Company:  "Freelance",
				Content:  "Wanderlust's guides have been my bible for the past year. I've found hidden gems I never would have discovered on my own!",
				ImageURL: "https://i.pravatar.cc/150?u=emily",
				Rating:   5,
			},
			{
				Name:     "James Peterson",
				Role:     "Photographer",
				Company:  "LensCrafters",
				Content:  "The community here is amazing. I've met fellow travelers in three different countries thanks to the meetup forums.",
				ImageURL: "https://i.pravatar.cc/150?u=james",
				Rating:   5,
			},
			{
				Name:     "Sophia Miller",
				Role:     "Travel Blogger",
				Company:  "Soph Travels",
				Content:  "Detailed, accurate, and inspiring. I recommend Wanderlust to all my followers who want authentic travel experiences.",
				ImageURL: "https://i.pravatar.cc/150?u=sophia",
				Rating:   5,
			},
		},
		FAQItems: []models.FAQItem{
			{Question: "How do I book a trip?", Answer: "We partner with trusted travel agencies. You can find booking links directly in our destination guides."},
			{Question: "Is Wanderlust free to use?", Answer: "Yes! All our guides and community features are completely free. We also offer premium courses for aspiring travel bloggers."},
			{Question: "Can I contribute a story?", Answer: "Absolutely! We love hearing from our community. Visit the 'Write for Us' page to submit your travel stories."},
			{Question: "Do you offer group tours?", Answer: "We occasionally organize group expeditions. Subscribe to our newsletter to be the first to know about upcoming trips."},
		},
	}
}

func (s *PageService) GetAboutData() *models.AboutPageData {
	return &models.AboutPageData{
		PageData: models.PageData{
			Title:          "About Wanderlust | Innovating Web Development",
			Description:    "Meet the team behind Wanderlust. We are passionate about building high-performance Go web applications and crafting unforgettable travel experiences.",
			Keywords:       "about us, team, wanderlust, go developers, company mission, tech startup",
			SocialImageURL: "/static/images/default-avatar.svg", // Ideally a team photo
			CurrentPage:    "about",
			CurrentYear:    time.Now().Year(),
			ContentType:    "about",
		},
		CompanyName:  "Wanderlust Solutions",
		CompanyStory: "We are passionate about building high-performance web applications using Go. Our team combines expertise in modern web development with deep knowledge of Go's ecosystem to deliver exceptional solutions.",
		FoundedYear:  2023,
		Mission:      "To inspire and enable everyone to explore the beauty of our planet.",
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
			Title:          "Client Success Stories | Wanderlust Reviews",
			Description:    "See how Wanderlust has transformed businesses with high-performance Go solutions. Read reviews from our satisfied clients.",
			Keywords:       "testimonials, reviews, client stories, wanderlust, go development case studies",
			SocialImageURL: "/static/images/default-avatar.svg",
			CurrentPage:    "testimonials",
			CurrentYear:    time.Now().Year(),
			ContentType:    "testimonials",
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
			Title:          "Master Go Programming | Wanderlust Courses",
			Description:    "Level up your skills with our expert-led Go programming courses. From fundamentals to advanced microservices and cloud-native development.",
			Keywords:       "go courses, golang tutorials, learn go, backend development, devops training, wanderlust",
			SocialImageURL: "https://images.unsplash.com/photo-1515879218367-8466d910aaa4?w=1200&q=80",
			CurrentPage:    "courses",
			CurrentYear:    time.Now().Year(),
			ContentType:    "courses",
		},
		Categories:       []string{"All", "Go", "Web Development", "DevOps", "Backend"},
		Courses:          filteredCourses,
		SelectedCategory: category,
	}
}

func (s *PageService) GetAgeCheckData() *models.PageData {
	return &models.PageData{
		Title:          "Age Verification Required | Wanderlust",
		Description:    "Please verify your age to access restricted content.",
		Keywords:       "age check, verification",
		CurrentPage:    "age-check",
		CurrentYear:    time.Now().Year(),
		ContentType:    "age-check",
	}
}

func (s *PageService) GetAccessDeniedData() *models.PageData {
	return &models.PageData{
		Title:          "Access Restricted | Wanderlust",
		Description:    "You must be of legal age to view this content.",
		Keywords:       "access denied, restricted",
		CurrentPage:    "access-denied",
		CurrentYear:    time.Now().Year(),
		ContentType:    "access-denied",
	}
}
func (s *PageService) GetStoriesData() *models.StoriesPageData {
	return &models.StoriesPageData{
		PageData: models.PageData{
			Title:          "Travel Stories & Adventures | Wanderlust Blog",
			Description:    "Immerse yourself in inspiring travel stories from around the globe. Discover hidden gems, culture, and adventure.",
			Keywords:       "travel blog, travel stories, wanderlust, adventure, culture, food",
			SocialImageURL: "https://images.unsplash.com/photo-1493976040374-85c8e12f0c0e?w=1200&q=80",
			CurrentPage:    "stories",
			CurrentYear:    time.Now().Year(),
			ContentType:    "stories",
		},
		Stories: []models.Story{
			{
				Title:    "Hidden Gems of Kyoto",
				Author:   "Yuki Tanaka",
				Location: "Kyoto, Japan",
				Category: "Culture",
				Content:  "Wandering through the bamboo groves of Arashiyama early in the morning, exploring centuries-old temples, and discovering the quiet beauty of Kyoto's traditional tea houses...",
				ImageURL: "https://images.unsplash.com/photo-1493976040374-85c8e12f0c0e?w=800&q=80",
			},
			{
				Title:    "Santorini Sunsets",
				Author:   "Elena Costa",
				Location: "Santorini, Greece",
				Category: "Relaxation",
				Content:  "The white-washed buildings clinging to the cliffside, the deep blue of the Aegean Sea, and the most magical sunset I have ever witnessed from Oia...",
				ImageURL: "https://images.unsplash.com/photo-1570077188670-e3a8d69ac5ff?w=800&q=80",
			},
			{
				Title:    "Trekking Patagonia",
				Author:   "Mark Stevens",
				Location: "Patagonia, Chile",
				Category: "Adventure",
				Content:  "Battling the winds at the base of the Torres del Paine, seeing glaciers calve into turquoise lakes, and feeling the raw power of nature at the end of the world...",
				ImageURL: "https://images.unsplash.com/photo-1518182170546-07fb61429c2f?w=800&q=80",
			},
			{
				Title:    "Culinary Journey in Rome",
				Author:   "Sofia Romano",
				Location: "Rome, Italy",
				Category: "Food",
				Content:  "From the perfect carbonara in Trastevere to the best gelato near the Pantheon. A journey through the eternal city, one bite at a time...",
				ImageURL: "https://images.unsplash.com/photo-1529260830199-42c42dda5f3d?w=800&q=80",
			},
			{
				Title:    "Northern Lights",
				Author:   "Erik Larson",
				Location: "Tromsø, Norway",
				Category: "Adventure",
				Content:  "Chasing the Aurora Borealis through the frozen landscapes of the Arctic Circle, dog sledding under the stars, and experiencing the magic of polar nights...",
				ImageURL: "https://images.unsplash.com/photo-1531366936337-7c912a4589a7?w=800&q=80",
			},
			{
				Title:    "Temples of Angkor",
				Author:   "Linh Nguyen",
				Location: "Siem Reap, Cambodia",
				Category: "Culture",
				Content:  "Watching the sunrise over Angkor Wat, exploring the jungle-claimed ruins of Ta Prohm, and feeling the ancient history of the Khmer Empire...",
				ImageURL: "https://images.unsplash.com/photo-1569668723429-9fb4119debe9?w=800&q=80",
			},
		},
	}
}
/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Updated: 27-12-25
 * └─ go-server ───┘
 */