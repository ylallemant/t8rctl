package credentials

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
)

var Current *azidentity.DefaultAzureCredential

func init() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(errors.Wrapf(err, "failed to obtain a credentials for \"%s\"\nplease run following command : az login --scope https://management.core.windows.net//.default", api.Azure))
	}

	Current = cred
}
