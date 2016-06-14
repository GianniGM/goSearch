package gSearch

import (
	"errors"
	"strconv"
)

type MySearchIFace interface {
	Search(tts string) string
	GetNext() string
	GetPrev() string
}

type MySearch struct {
	ID            int64
	searchResults []string
	index         int
}

// func (u *MySearch) Set(Name string) {
// 	u.Name = Name
// }

// func (u *MySearch) SetName(Name string) {
// 	u.Name = Name
// 	u.index = 0
// }

func (u MySearch) GetID() int64 {
	return u.ID
}

func (u *MySearch) Search(tts string) (string, error) {
	u.index = 0
	var err error

	if u.searchResults, err = SendGoogleSearchRequest(tts); err != nil {
		return "", errors.New("NOT_FOUND")
	} else {
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[0] + "\n", nil
	}
}

func (u *MySearch) GetNext() (string, string) {
	if u.index+1 < len(u.searchResults) {
		u.index++
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[u.index] + "\n", ""
	} else {
		return "", "no more results! :)\n"
	}

}

func (u *MySearch) GetPrev() (string, string) {
	if u.index > 0 {
		u.index--
		return strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[u.index] + "\n", ""
	} else {
		return "", strconv.Itoa(u.index+1) + "/" + strconv.Itoa(len(u.searchResults)) + "\n" + u.searchResults[0] + "\n"
	}
}
