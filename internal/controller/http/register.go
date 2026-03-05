package httpController

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, qc *QuestionController, oc *OptionsController) {
	r.Route("/questions", func(r chi.Router) {
		r.Get("/find-all", qc.FindAll)
		r.Post("/add-question", qc.AddQuestion)
	})

	r.Route("/options", func(r chi.Router) {
		r.Get("/find-all", oc.FindAll)
		r.Get("/{questionId}", oc.FindByQuestionID)
		r.Post("/add-option", oc.AddOption)
	})
}
