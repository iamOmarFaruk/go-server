package middleware

import (
	"net/http"
)

// AgeGate middleware checks for an age_verified cookie.
// If not found, it redirects the user to the age check page.
func AgeGate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("age_verified")
		if err != nil || cookie.Value != "true" {
			http.Redirect(w, r, "/age-check", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*
 * ┌── o m a r ──┐
 * │ @iamOmarFaruk
 * │ omarfaruk.dev
 * │ Created: 2025-12-21
 * │ Updated: 2025-12-21
 * └─ go-server ───┘
 */
