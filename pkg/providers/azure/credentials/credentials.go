package credentials

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
)

var Current *azidentity.DefaultAzureCredential

func init() {
	cred, err := newCredential()
	if err != nil {
		panic(errors.Wrapf(err, "failed to initialize credentials for \"%s\"", api.Azure))
	}

	Current = cred
}

func newCredential() (*azidentity.DefaultAzureCredential, error) {
	// This CLI defaults to developer credentials; deployment credentials remain opt-in.
	if _, configured := os.LookupEnv("AZURE_TOKEN_CREDENTIALS"); !configured {
		if err := os.Setenv("AZURE_TOKEN_CREDENTIALS", "dev"); err != nil {
			return nil, errors.Wrap(err, "could not configure default Azure credentials")
		}
	}

	return azidentity.NewDefaultAzureCredential(nil)
}
