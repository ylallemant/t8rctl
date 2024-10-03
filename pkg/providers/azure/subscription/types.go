package subscription

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/ylallemant/t8rctl/pkg/account"
	"github.com/ylallemant/t8rctl/pkg/api"
)

func NewSubscription(subscription *armsubscriptions.Subscription) Subscription {
	return Subscription{
		Name:     *subscription.DisplayName,
		Id:       *subscription.SubscriptionID,
		TenantId: *subscription.TenantID,
	}
}

type Subscription struct {
	Name     string
	Id       string
	TenantId string
}

func (i *Subscription) convert() (api.Account, error) {
	properties := map[string]string{
		"tenantId": i.TenantId,
	}

	return account.New(
		api.Azure,
		i.Id,
		i.Name,
		properties,
	)
}
