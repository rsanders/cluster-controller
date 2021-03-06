package client

const (
	DownwardAPIVolumeSourceType             = "downwardAPIVolumeSource"
	DownwardAPIVolumeSourceFieldDefaultMode = "defaultMode"
	DownwardAPIVolumeSourceFieldItems       = "items"
)

type DownwardAPIVolumeSource struct {
	DefaultMode *int64                  `json:"defaultMode,omitempty"`
	Items       []DownwardAPIVolumeFile `json:"items,omitempty"`
}
