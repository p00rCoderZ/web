package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "To pole nie może być puste!")
		}
	}
}

//MaxLength ..
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("To pole jest zbyt długie (maksymalna liczba znaków %d )", d))
	}
}
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("To pole jest zbyt krótkie (minimalna liczba znaków %d )", d))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "To pole jest niepoprawne")
	}
}

func (f *Form) ExistsEmail(field string) {
	f.Errors.Add(field, "Email już istnieje!")

}

func (f *Form) ExistsNick(field string) {
	f.Errors.Add(field, "Nick już istnieje!")

}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
