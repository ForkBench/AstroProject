package routes

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"astroproject/astro/services"
	"astroproject/components"
	"astroproject/pages"
)

// NewChiRouter creates a new chi router.
func NewChiRouter(session *services.Session) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		HXRender(w, r, pages.HomePage(), session)
	})

	r.Post("/add-competition", func(w http.ResponseWriter, r *http.Request) {
		if session.AddCompetition("Competition", "U20", "Foil") {
			templ.Handler(components.CompetitionElement(session.GetLastCompetition())).ServeHTTP(w, r)
		}
	})

	r.Get("/competition/{competitionID}", func(w http.ResponseWriter, r *http.Request) {
		competitionIDStr := chi.URLParam(r, "competitionID")
		competitionID, err := strconv.ParseUint(competitionIDStr, 10, 8)
		if err != nil {
			// Return a 404.
			http.NotFound(w, r)
		}
		HXRender(w, r, pages.CompetitionPage(session.GetCompetition(uint8(competitionID))), session)
	})

	return r
}
