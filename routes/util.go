package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/mavolin/go-htmx"

	"astroproject/astro/services"
	"astroproject/components"
)

func HXRender(w http.ResponseWriter, r *http.Request, component templ.Component, session *services.Session) {
	hxRequest := htmx.Request(r)

	if hxRequest == nil {
		component = components.Page(component, session)
	}
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
