package user

import (
	"errors"
	"googleSearch/gSearchV2"
	"strconv"
)

type UserIFace interface {
	Set(Name string)
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

// func (u *User) Set(Name string) {
// 	u.Name = Name
// }

// func (u *User) SetName(Name string) {
// 	u.Name = Name
// 	u.index = 0
// }

func (u User) GetName() string {
	return u.Name
}

func (u *User) Search(tts string) (string, error) {
	u.index = 0
	u.searchResults = *gSearch.SendGoogleSearchRequest(tts)
	if u.searchResults == nil {
		return "", errors.New("NOT_FOUND")
	} else {
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[0] + "\n", nil
	}
}

func (u *User) GetNext() (string, string) {
	if u.index+1 < len(u.searchResults) {
		u.index++
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[u.index] + "\n", ""
	} else {
		return "", "no more results! :)\n"
	}

}

func (u *User) GetPrev() (string, string) {
	if u.index > 0 {
		u.index--
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[u.index] + "\n", ""
	} else {
		return "", strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[0] + "\n"
	}
}
