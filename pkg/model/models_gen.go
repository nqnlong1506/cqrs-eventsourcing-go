// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Genre string

const (
	GenreRomantic   Genre = "ROMANTIC"
	GenreFiction    Genre = "FICTION"
	GenreNonfiction Genre = "NONFICTION"
	GenreStory      Genre = "STORY"
)

var AllGenre = []Genre{
	GenreRomantic,
	GenreFiction,
	GenreNonfiction,
	GenreStory,
}

func (e Genre) IsValid() bool {
	switch e {
	case GenreRomantic, GenreFiction, GenreNonfiction, GenreStory:
		return true
	}
	return false
}

func (e Genre) String() string {
	return string(e)
}

func (e *Genre) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Genre(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Genre", str)
	}
	return nil
}

func (e Genre) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusAvailable Status = "AVAILABLE"
	StatusBorrowed  Status = "BORROWED"
)

var AllStatus = []Status{
	StatusAvailable,
	StatusBorrowed,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusAvailable, StatusBorrowed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}