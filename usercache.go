package minecraft

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// UserCache represents users found in the cache
type UserCache []User

// GetByID will return a cached user with the specified id and if found
func (uc UserCache) GetByID(id string) (User, bool) {
	for _, user := range uc {
		if strings.EqualFold(user.ID, id) {
			return user, true
		}
	}

	return User{}, false
}

// Get will return a cached user with the specified name and if found
func (uc UserCache) Get(name string) (User, bool) {
	for _, user := range uc {
		if strings.EqualFold(user.Name, name) {
			return user, true
		}
	}

	return User{}, false
}

// User represents a user in the usercache.json
type User struct {
	Name    string `json:"name"`
	ID      string `json:"uuid"`
	Expires Time   `json:"expiresOn"`
}

// ReadCache will decode the user cache from the provided io.Reader
func ReadCache(r io.Reader) (UserCache, error) {
	cache := UserCache{}

	err := json.NewDecoder(r).Decode(&cache)

	return cache, errors.Wrap(err, "failed to read user cache")
}

// OpenCache will open and read the UserCache from the specified path (e.g /opt/minecraft/usercache.json)
func OpenCache(path string) (UserCache, error) {
	cache, err := os.Open(path)
	if err != nil {
		return UserCache{}, errors.Wrapf(err, "failed to open cache %q", path)
	}

	defer cache.Close()

	return ReadCache(cache)
}
