package database

import (
	"github.com/raphaelc484/go-email.git/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}
