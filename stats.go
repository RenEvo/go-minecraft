package minecraft

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

// StatsGroups represents the stats collection
type StatsGroups map[string]Stats

// Stats represents a collection of stats
type Stats map[string]int

// Statistics represents a collection of versioned stats
type Statistics struct {
	Groups  StatsGroups `json:"stats"`
	Version int         `json:"DataVersion"`
}

// UserStatistics includes the user and the statistics
type UserStatistics struct {
	Statistics
	User
}

// ReadStatistics will decode the statistics from the provided io.Reader
func ReadStatistics(r io.Reader) (Statistics, error) {
	statistics := Statistics{}

	err := json.NewDecoder(r).Decode(&statistics)

	return statistics, errors.Wrap(err, "failed to decode statistics")
}

// OpenStatistics will open and read the Statitics from the specified path (e.g /opt/minecraft/world/stats/playerid.json)
func OpenStatistics(path string) (Statistics, error) {
	stats, err := os.Open(path)
	if err != nil {
		return Statistics{}, errors.Wrapf(err, "failed to open stats %q", path)
	}

	defer stats.Close()

	return ReadStatistics(stats)
}
