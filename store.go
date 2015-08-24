package mpower

// The Store holds the store information and ised to define the store data for mpower transaction
type Store struct {
	Name          string `json:"name,omitempty"`
	Tagline       string `json:"tagline,omitempty"`
	PhoneNumber   string `json:"phone,omitempty"`
	PostalAddress string `json:"postal_address,omitempty"`
	LogoURL       string `json:"logo_url,omitempty"`
}

// NewStore - returns a new store object
//
// Example.
//    mpowerStore := NewStore("Awesome Store")
func NewStore(name string) *Store {
	store := &Store{}
	store.Name = name
	return store
}
