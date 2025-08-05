package types

type Cfg struct {
	LoggedIn    bool  `json:"loggedIn"`
	CurrentUser *User `json:"currentUser"`
}
