package types

// Device data model
type Device struct {
	ID          string `json:"id,omitempty"`
	DeviceModel string `json:"deviceModel,omitempty"`
	Name        string `json:"name,omitempty"`
	Note        string `json:"note,omitempty"`
	Serial      string `json:"serial,omitempty"`
}
