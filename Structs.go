package BluesoundAPI


// BluesoundStatus is Status for player
type BluesoundStatus struct {
	Album		string
	Artist		string
	Name 		string
	Quality		string
	Length		int16
	Second		int16
	Artwork		string
	Filename 	string
	Repeat		string
	Volume 		int16
	Service 	string
	State		string
	Shuffle		string
}

// BluesoundSyncStatus is SyncStatus for player
type BluesoundSyncStatus struct {
	Icon		string	`xml:"icon,attr"`
	Volume 		int		`xml:"volume,attr"`
	ModelName	string	`xml:"modelName,attr"`
	Name		string	`xml:"name,attr"`
	Model		string	`xml:"model,attr"`
	Brand		string	`xml:"brand,attr"`
	Localetag	int		`xml:"etag,attr"`
	SchemaVersion	int	`xml:"schemaVersion,attr"`
	SyncStat 	int		`xml:"syncStat,attr"`
	ID			string	`xml:"id,attr"`
	MAC			string	`xml:"mac,attr"`
}

// BluesoundController is the main controller object
type BluesoundController struct {
	Status     BluesoundStatus
	SyncStatus BluesoundSyncStatus
	Name       string
	baseURL    string
}

const(
	BluesoundHTTPURI_Status		string 	= "Status"
	Bluesound_ShuffleOn 		int16 	= 1
	Bluesound_ShuffleOff 		int16 	= 0
	Bluesound_RepeatAll			int16	= 0
	Bluesound_RepeatTrack		int16	= 1
	Bluesound_RepeatOff			int16	= 2
)