package files

import (
	"HoldMyLink_Bot/lib/e"
	"HoldMyLink_Bot/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}
func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save page", err) }()
	fPath := filepath.Join(s.basePath, page.UserName)
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)
	if err != nil {
		return err
	}
	fPath = filepath.Join(fPath, fName)
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}
	return nil

}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("cant pick random", err) }()

	path := filepath.Join(s.basePath, userName)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, storage.ErrorsNoSavedPages
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))
	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("cant remove file", err)

	}
	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		errMsg := fmt.Sprintf("cant remove file %s", path)
		return e.Wrap(errMsg, err)
	}
	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("cant check file is exists", err)

	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		errMsg := fmt.Sprintf("cant check file %s is exists", path)
		return false, e.Wrap(errMsg, err)
	}
	return true, nil

}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("cant decode page", err)
	}
	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("cant decode page", err)

	}
	return &p, nil

}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
