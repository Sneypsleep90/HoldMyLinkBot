package storage

import (
	"HoldMyLink_Bot/lib/e"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrorsNoSavedPages = errors.New("no saved page")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("cant calculate hash", err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("cant calculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
