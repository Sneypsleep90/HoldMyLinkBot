package telegram

import (
	"HoldMyLink_Bot/lib/e"
	"HoldMyLink_Bot/storage"
	"errors"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s' ", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)

	default:
		return p.tg.SendMessages(chatID, msgUnknownCommand)

	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("cant do command: save page", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	IsExist, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if IsExist {
		return p.tg.SendMessages(chatID, msgAlreadyExists)
	}
	if err := p.storage.Save(page); err != nil {
		return err
	}
	if err := p.tg.SendMessages(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("cant command: cant send random", err) }()
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrorsNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrorsNoSavedPages) {
		return p.tg.SendMessages(chatID, msgNoSavedPages)

	}

	if err := p.tg.SendMessages(chatID, page.URL); err != nil {
		return err
	}
	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessages(chatID, msgHelp)
}
func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessages(chatID, msgHello)

}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""

}
