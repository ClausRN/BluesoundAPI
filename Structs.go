package BluesoundAPI

import (
	"encoding/xml"
	"os"
	"sync"
)

// BluesoundStatus is Status for player
type BluesoundStatus struct {
	XMLName  xml.Name `xml:"status"`
	Album    string   `xml:"album"`
	Artist   string   `xml:"artist"`
	Name     string   `xml:"name"`
	Quality  string   `xml:"quality"`
	Format   string   `xml:"streamformat"`
	Length   int16    `xml:"totlen"`
	Second   int16    `xml:"secs"`
	Artwork  string   `xml:"image"`
	Filename string   `xml:"fn"`
	Repeat   string   `xml:"repeat"`
	Volume   int16    `xml:"volume"`
	Service  string   `xml:"service"`
	State    string   `xml:"state"`
	Shuffle  string   `xml:"shuffle"`
	SongNo   int      `xml:"song"`
}

// BluesoundSyncStatus is SyncStatus for player
type BluesoundSyncStatus struct {
	XMLName       xml.Name `xml:"SyncStatus"`
	Icon          string   `xml:"icon,attr"`
	Volume        int      `xml:"volume,attr"`
	ModelName     string   `xml:"modelName,attr"`
	Name          string   `xml:"name,attr"`
	Model         string   `xml:"model,attr"`
	Brand         string   `xml:"brand,attr"`
	SchemaVersion int      `xml:"schemaVersion,attr"`
	SyncStat      int      `xml:"syncStat,attr"`
	ID            string   `xml:"id,attr"`
	MAC           string   `xml:"mac,attr"`
}

// BluesoundCommandState Returned by player command request
type BluesoundCommandState struct {
	XMLName xml.Name `xml:"state"`
	State   string   `xml:",chardata"`
}

// BluesoundCommandStateSkipBack Returned by player command request
type BluesoundCommandStateSkipBack struct {
	XMLName xml.Name `xml:"id"`
	TrackID int32    `xml:",chardata"`
}

// BluesoundVolume Returned by player set volume request
type BluesoundVolume struct {
	XMLName xml.Name `xml:"volume"`
	Volume  int      `xml:",chardata"`
}

// BluesoundPlayerVersion Returns player software version
type BluesoundPlayerVersion struct {
	XMLName xml.Name `xml:"version"`
	Version string   `xml:",chardata"`
}

// BluesoundPlaylistNames Returned by player command request
type BluesoundPlaylistNames struct {
	//XMLName xml.Name `xml:"name"`
	Artwork string `xml:"image,attr"`
	Name    string `xml:",chardata"`
}

// BluesoundPlaylists Returned by player command request
type BluesoundPlaylists struct {
	XMLName   xml.Name                 `xml:"playlists"`
	Service   string                   `xml:"service,attr"`
	Playlists []BluesoundPlaylistNames `xml:"name"`
}

// BluesoundTrackName Returned by player command request
type BluesoundTrackName struct {
	//XMLName xml.Name `xml:"name"`
	AlbumID  int32  `xml:"albumid,attr"`
	ArtistID int32  `xml:"artistid,attr"`
	SongID   string `xml:"songid,attr"`
	Service  string `xml:"service,attr"`
	Artist   string `xml:"art"`
	Album    string `xml:"alb"`
	Title    string `xml:"title"`
	Filename string `xml:"fn"`
	Length   int32  `xml:"time"`
	TrackNo  int32  `xml:"track"`
	DiscNo   int32  `xml:"discno"`
	Quality  string `xml:"quality"`
}

// BluesoundPlaylist Returned by player command request
type BluesoundPlaylist struct {
	XMLName xml.Name             `xml:"songs"`
	Service string               `xml:"service,attr"`
	ID      int32                `xml:"id,attr"`
	Tracks  []BluesoundTrackName `xml:"song"`
}

// BluesoundAddSong Returned by player command request
type BluesoundAddSong struct {
	XMLName xml.Name `xml:"addsong"`
	Added   int32    `xml:"count,attr"`
	Total   int32    `xml:"length,attr"`
	ID      int32    `xml:"id,attr"`
}

// BluesoundPlayQueue Returned by player command request
type BluesoundPlayQueue struct {
	XMLName  xml.Name             `xml:"playlist"`
	Name     string               `xml:"name,attr"`
	Modified int32                `xml:"modified,attr"`
	Length   int32                `xml:"length,attr"`
	ID       int32                `xml:"id,attr"`
	Shuffle  int                  `xml:"shuffle,attr"`
	Repeat   int                  `xml:"repeat,attr"`
	Tracks   []BluesoundTrackName `xml:"song"`
}

// BluesoundController is the main controller object
type BluesoundController struct {
	status     BluesoundStatus
	syncStatus BluesoundSyncStatus
	stopSignal chan os.Signal
	baseURL    string
}

var (
	statusUpdate sync.RWMutex
)

const (
	bluesoundHTTPURIStatus     string = "Status"
	bluesoundHTTPURISyncStatus string = "SyncStatus"
	bluesoundHTTPURIPlay       string = "Play"
	bluesoundHTTPURIPause      string = "Pause"
	bluesoundHTTPURIBack       string = "Back"
	bluesoundHTTPURISkip       string = "Skip"
	// ShuffleOn Shuffle is on
	ShuffleOn int16 = 1
	// ShuffleOff Shuffle is off
	ShuffleOff int16 = 0
	// RepeatAll Repeat all tracks
	RepeatAll int16 = 0
	// RepeatTrack Repeat single track
	RepeatTrack int16 = 1
	// RepeatOff Repeat is off
	RepeatOff int16 = 2
)
