package campaign

import (
	"errors"
	"testing"

	"github.com/raphaelc484/go-email.git/internal/contract"
	internalerrors "github.com/raphaelc484/go-email.git/internal/internalErrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Bodys",
		Emails:  []string{"teste1@test.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidationDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignDTO{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidationRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
