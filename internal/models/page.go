package models

type PageData struct {
	Title       string
	Description string
	CurrentPage string
	CurrentYear int
	ContentType  string  // home, about, testimonials, 404
}

type HomePageData struct {
	PageData
	HeroTitle    string
	HeroSubtitle string
	Features     []Feature
}

type AboutPageData struct {
	PageData
	CompanyName    string
	CompanyStory   string
	TeamMembers    []TeamMember
	FoundedYear    int
	Mission        string
}

type TestimonialPageData struct {
	PageData
	Testimonials []Testimonial
	Stats        []Stat
}

type Feature struct {
	Title       string
	Description string
	Icon        string
}

type TeamMember struct {
	Name      string
	Role      string
	Bio       string
	ImageURL  string
}

type Testimonial struct {
	Name       string
	Role       string
	Company    string
	Content    string
	Rating     int
	ImageURL   string
}

type Stat struct {
	Number string
	Label  string
}

type Course struct {
	Title       string
	Description string
	Instructor  string
	Duration    string
	Level       string
	Price       string
	ImageURL    string
	Category    string
	Students    int
}

type CoursesPageData struct {
	PageData
	Courses    []Course
	Categories []string
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Touched: 2025-12-21
 * └─ go-server ───┘
 */