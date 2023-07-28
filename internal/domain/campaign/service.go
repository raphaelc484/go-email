package campaign

import (
	"github.com/raphaelc484/go-email.git/internal/contract"
	internalerrors "github.com/raphaelc484/go-email.git/internal/internalErrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaignDTO) (string, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}
