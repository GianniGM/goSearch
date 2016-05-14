package user

import "googleSearch/gSearch"

type UserIFace interface {
	Set(name string)
	GetName() string
	Search(tts string) string
	GetNext() string
	GetPrev() string
}

type User struct {
	Name          string
	searchResults []string
	index         int
}

// func (u *User) Set(name string) {
// 	u.name = name
// }

func (u User) GetName() string {
	return u.Name
}

func (u *User) Search(tts string) string {
	u.index = 0
	u.searchResults = gSearch.SendGoogleSearchRequest(tts)
	if u.searchResults == nil {
		return "Sorry! Not Found"
	} else {
		return u.searchResults[0]
	}
}

func (u *User) GetNext() string {
	if u.index < len(u.searchResults) {
		u.index++
		return u.searchResults[u.index]
	} else {
		return "no more results! :)"
	}

}

func (u *User) GetPrev() string {
	if u.index > 0 {
		u.index--
		return u.searchResults[u.index]
	} else {
		return u.searchResults[0]
	}
}
