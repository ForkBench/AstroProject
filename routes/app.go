package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-faker/faker/v4"

	"astroproject/astro/services"
	"astroproject/components"
	"astroproject/pages"
)

// NewChiRouter creates a new chi router.
func NewChiRouter(session *services.Session) *chi.Mux {

	r := chi.NewRouter()

	logger := httplog.NewLogger("astro-logger", httplog.Options{
		// All log
		LogLevel: slog.LevelInfo,
		Concise:  true,
	})

	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	// ULTRA IMPORTANT : This middleware is used to prevent caching of the pages.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
		})
	})

	// Serve static files.
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		HXRender(w, r, pages.HomePage(), session)
		w.WriteHeader(http.StatusOK)
	})

	// Competitions
	r.Route("/competition", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			competition := session.AddCompetition(faker.FirstName(), "U17", "Foil")

			HXRender(w, r, components.CompetitionElement(competition), session)
			// 303 See Other.
			w.WriteHeader(http.StatusSeeOther)
		})

		r.Route("/{competitionID}", func(r chi.Router) {

			// Get competition by ID.
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				string_id := chi.URLParam(r, "competitionID")
				id, err := strconv.Atoi(string_id)
				if err != nil {
					// Status Not Found.
					w.WriteHeader(http.StatusNotFound)
					return
				}
				HXRender(w, r, pages.CompetitionPage(session.GetCompetition(uint8(id))), session)
				w.WriteHeader(http.StatusOK)
			})

			// Remove competition by ID.
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				string_id := chi.URLParam(r, "competitionID")
				id, err := strconv.Atoi(string_id)
				if err != nil {
					// Status Not Found.
					w.WriteHeader(http.StatusNotFound)
					return
				}

				if session.RemoveCompetition(uint8(id)) {
					// Status OK.
					w.WriteHeader(http.StatusOK)
				} else {
					// Status Not Found.
					w.WriteHeader(http.StatusNotFound)
				}

			})

			r.Route("/stage/{stageID}", func(r chi.Router) {

				// Get stage by ID.
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					stage_id_string := chi.URLParam(r, "stageID")
					competition_id_string := chi.URLParam(r, "competitionID")
					stage_id, err := strconv.Atoi(stage_id_string)
					if err != nil {
						// Status Not Found.
						w.WriteHeader(http.StatusNotFound)
						return
					}

					competition_id, err := strconv.Atoi(competition_id_string)
					if err != nil {
						// Status Not Found.
						w.WriteHeader(http.StatusNotFound)
						return
					}

					competition := session.GetCompetition(uint8(competition_id))

					if competition == nil {
						// Status Not Found.
						w.WriteHeader(http.StatusNotFound)
						return
					}

					stage := competition.GetStage(uint8(stage_id))

					if stage == nil {
						// Status Not Found.
						w.WriteHeader(http.StatusNotFound)
						return
					}

					stageKind := (*stage).GetKind()

					switch stageKind {
					case services.REGISTRATIONS:
						HXRender(w, r, components.DisplaySeedings(competition, (*stage).(*services.SeedingStage)), session)
						// 303 See Other.
						w.WriteHeader(http.StatusSeeOther)
					default:
						// Status Not Found.
						w.WriteHeader(http.StatusNotFound)
					}
				})

				r.Route("/player", func(r chi.Router) {

					// Add player to stage.
					r.Post("/", func(w http.ResponseWriter, r *http.Request) {
						stage_id_string := chi.URLParam(r, "stageID")
						competition_id_string := chi.URLParam(r, "competitionID")
						stage_id, err := strconv.Atoi(stage_id_string)
						if err != nil {
							// Status Not Found.
							w.WriteHeader(http.StatusNotFound)
							return
						}

						competition_id, err := strconv.Atoi(competition_id_string)
						if err != nil {
							// Status Not Found.
							w.WriteHeader(http.StatusNotFound)
							return
						}

						competition := session.GetCompetition(uint8(competition_id))
						stage := competition.GetStage(uint8(stage_id))

						if (*stage).GetKind() != services.REGISTRATIONS {
							// Status Not Found.
							w.WriteHeader(http.StatusNotFound)
							return
						}

						seedingStage := (*stage).(*services.SeedingStage)

						player := services.GenerateRandomPlayer()

						player.PlayerInitialRank = competition.CompetitionPlayerNumber + 1

						// TODO: Remove competition.AddPlayer
						if competition.AddPlayer(player) && (seedingStage).AddPlayer(player) {
							HXRender(w, r, components.PlayerElement(seedingStage.SeedingSeedings[player.PlayerID]), session)

							// 303 See Other.
							w.WriteHeader(http.StatusSeeOther)
						} else {
							// Status Not Found.
							w.WriteHeader(http.StatusNotFound)
						}
					})

				})

			})
		})
	})

	return r
}
