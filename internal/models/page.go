package models

type PageData struct {
	Title          string
	Description    string
	Keywords       string
	SocialImageURL string
	CurrentPage    string
	CurrentYear    int
	ContentType    string // home, about, testimonials, 404
	IsDev          bool
}

type HomePageData struct {
	PageData
	HeroSlides   []HeroSlide
	Features     []Feature
}

type HeroSlide struct {
	Title       string
	Subtitle    string
	CTA1Text    string
	CTA1Link    string
	CTA2Text    string
	CTA2Link    string
	ImageURL    string
	Gradient    string
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
	ImageURL    string
	Category    string
	Location    string
	Author      string
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

type Story struct {
	Title    string
	Author   string
	Location string
	Category string
	Content  string
	ImageURL string
}

type StoriesPageData struct {
	PageData
	Stories []Story
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
	Courses          []Course
	Categories       []string
	SelectedCategory string
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Touched: 2025-12-21
 * └─ go-server ───┘
 */