package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/raphaelc484/go-email.git/internal/contract"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err
}
