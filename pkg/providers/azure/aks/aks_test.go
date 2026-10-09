package aks

import (
	"errors"
	"testing"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/global"
)

type failingAccounts struct {
	api.AccountManager
	err error
}

func (a failingAccounts) List() ([]api.Account, error) {
	return nil, a.err
}

func TestListBypassesMemoryCache(t *testing.T) {
	previous := global.Current.DisableCache
	t.Cleanup(func() { global.Current.DisableCache = previous })

	client := &AksClient{cache: []api.Cluster{nil}}
	wantErr := errors.New("account fetch attempted")
	accounts := failingAccounts{err: wantErr}

	global.Current.DisableCache = false
	if clusters, err := client.List(accounts); err != nil || len(clusters) != 1 {
		t.Fatalf("cached List() = %v, %v", clusters, err)
	}

	global.Current.DisableCache = true
	if _, err := client.List(accounts); !errors.Is(err, wantErr) {
		t.Fatalf("uncached List() error = %v, want %v", err, wantErr)
	}
}
