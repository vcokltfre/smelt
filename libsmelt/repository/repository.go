package repository

import (
	"encoding/json"
	"os"
	"path"
)

type Repository struct {
	Name string
	Root string
}

type RepositoryConfig struct {
	Name string `json:"name"`
}

func (r *Repository) Save() error {
	repoConfig := &RepositoryConfig{
		Name: r.Name,
	}

	data, err := json.Marshal(repoConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(r.Root, ".smelt", "repository"), data, 0644)
}
