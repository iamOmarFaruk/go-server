package models

type PageData struct {
	Title       string
	Description string
	CurrentPage string
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