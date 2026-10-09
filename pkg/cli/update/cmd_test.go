package update

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/google/go-github/v92/github"
)

func TestMatchingReleaseAssets(t *testing.T) {
	name := fmt.Sprintf("t8rctl_1.0.0_%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)
	binary := &github.ReleaseAsset{Name: github.Ptr(name)}
	checksum := &github.ReleaseAsset{Name: github.Ptr(name + ".md5")}
	unrelated := &github.ReleaseAsset{Name: github.Ptr("README.md")}
	release := &github.RepositoryRelease{
		Assets: []*github.ReleaseAsset{unrelated, checksum, binary},
	}

	gotBinary, found := matchingBinary(release)
	if !found || gotBinary != binary {
		t.Fatalf("matchingBinary() = %v, %v; want selected binary asset", gotBinary, found)
	}
	gotChecksum, found := matchingChecksum(release)
	if !found || gotChecksum != checksum {
		t.Fatalf("matchingChecksum() = %v, %v; want selected checksum asset", gotChecksum, found)
	}

	release.Assets = []*github.ReleaseAsset{unrelated}
	if asset, found := matchingBinary(release); found || asset != nil {
		t.Fatalf("matchingBinary() = %v, %v; want nil, false", asset, found)
	}
	if asset, found := matchingChecksum(release); found || asset != nil {
		t.Fatalf("matchingChecksum() = %v, %v; want nil, false", asset, found)
	}
}
