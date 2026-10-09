package ps

import (
	"testing"

	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/cluster/options"
	"github.com/ylallemant/t8rctl/pkg/global"
)

func TestRunAppliesDisableCache(t *testing.T) {
	previousOptions := *globalOptions.Current
	previousGlobal := *global.Current
	t.Cleanup(func() {
		*globalOptions.Current = previousOptions
		*global.Current = previousGlobal
	})

	globalOptions.Current.Provider = "nonexistent-test-provider"
	for _, disabled := range []bool{true, false} {
		globalOptions.Current.DisableCache = disabled
		if err := rootCmd.RunE(rootCmd, nil); err == nil {
			t.Fatal("expected missing provider error")
		}
		if global.Current.DisableCache != disabled {
			t.Fatalf("DisableCache = %v, want %v", global.Current.DisableCache, disabled)
		}
	}
}
