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
//    s.mpowerStore = NewStore("Awesome Store", "Easy shopping", "0272271893", "P.0. Box MP555, Accra", "http://www.awesomestore.com.gh/logo.png")
func NewStore(name, tagline, phoneNumber, postalAddress, logoURL string) *Store {
	store := &Store{
		Name:          name,
		Tagline:       tagline,
		PhoneNumber:   phoneNumber,
		PostalAddress: postalAddress,
		LogoURL:       logoURL,
	}

	return store
}
