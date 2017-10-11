package BluesoundAPI

// BluesoundStatus is Status for player
type BluesoundStatus struct {
	SongNo         int32
	CurrentTrackNo int32
	Songs          []string
}

// BluesoundSyncStatus is SyncStatus for player
type BluesoundSyncStatus struct {
	SongNo         int32
	CurrentTrackNo int32
	Songs          []string
}

// BluesoundController is the main controller object
type BluesoundController struct {
	Status     BluesoundStatus
	SyncStatus BluesoundSyncStatus
	Name       string
}
