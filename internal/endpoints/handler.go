package endpoints

import "github.com/raphaelc484/go-email.git/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
