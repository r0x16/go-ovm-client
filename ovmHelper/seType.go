package ovmHelper

// storageElement element
type Se struct {
	Id          *Id    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Locked      bool   `json:"locked,omitempty"`   //false,
	ReadOnly    bool   `json:"readOnly,omitempty"` //: false,
	Generation  int    `json:"generation,omitempty"`
	Page83Id    string `json:"page83Id,omitempty"`
	Shareable   bool   `json:"shareable,omitempty"`
}
