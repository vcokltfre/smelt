package repository

import (
	"errors"
	"os"
	"path"
)

func Init(name, at string) (*Repository, error) {
	fd, err := os.Stat(at)
	if err != nil {
		return nil, err
	}

	if !fd.IsDir() {
		return nil, errors.New("not a directory")
	}

	_, err = os.Stat(path.Join(at, ".smelt"))
	if err == nil {
		return nil, errors.New("repository already exists")
	}

	err = os.Mkdir(path.Join(at, ".smelt"), 0755)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		Name: name,
		Root: at,
	}

	err = repo.Save()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
