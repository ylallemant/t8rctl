package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
)

var (
	DefaultCacheClusters = "clusters.yaml"
	DefaultFolder        = ".t8rctl"
	CacheFolder          = "cache"
	DefaultTTL           = "24h"
	HomeFolder           string
	_                    api.Cache = &cache{}
)

func init() {
	var err error
	HomeFolder, err = os.UserHomeDir()
	if err != nil {
		panic(errors.Wrapf(err, "could not find the home directory"))
	}
}

func New(path, ttl string) (*cache, error) {
	entity := new(cache)

	basepath := filepath.Dir(path)
	os.MkdirAll(basepath, os.ModePerm)

	entity.path = path

	duration, err := time.ParseDuration(ttl)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse the ttl duration: %s", ttl)
	}

	entity.ttl = duration

	CurrentManager.register(entity)

	return entity, nil
}

func BasePath() string {
	return filepath.Join(HomeFolder, DefaultFolder, CacheFolder)
}

type cache struct {
	path string
	age  time.Duration
	size int64
	ttl  time.Duration
}

func (c *cache) Path() string {
	return c.path
}

func (c *cache) Age() time.Duration {
	return c.age
}

func (c *cache) Size() int64 {
	return c.size
}

func (c *cache) TTL() time.Duration {
	return c.ttl
}

func (c *cache) Expires() time.Time {
	return time.Now().Add(c.ttl - c.age)
}

func (c *cache) Exists() bool {
	if stats, err := os.Stat(c.path); err == nil {
		c.age = time.Since(stats.ModTime())
		c.size = stats.Size()

		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}

func (c *cache) Valid() bool {
	if !c.Exists() {
		return false
	}

	if c.size == 0 {
		return false
	}

	return c.age < c.ttl
}

func (c *cache) Read() ([]byte, error) {
	content, err := ioutil.ReadFile(c.path)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "could not read cache file")
	}

	return content, nil
}

func (c *cache) Write(content []byte) error {
	handler, err := os.OpenFile(c.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		errors.Wrapf(err, "could not open cache file handler")
	}
	defer handler.Close()

	_, err = handler.Write(content)
	if err != nil {
		errors.Wrapf(err, "could not write to cache file")
	}

	return nil
}

func (c *cache) Purge() error {
	if c.Exists() {
		err := os.Remove(c.path)
		if err != nil {
			return errors.Wrapf(err, "failed to delete cache file %s", c.path)
		}

		fmt.Println("purged cache file", c.path)
	}
	return nil
}
