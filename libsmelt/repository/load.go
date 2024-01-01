package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Load(at string) (*Repository, error) {
	at, err := filepath.Abs(at)
	if err != nil {
		return nil, err
	}

	pathParts := strings.Split(at, string(os.PathSeparator))

	for {
		if len(pathParts) == 0 {
			return nil, errors.New("not in a smelt repository")
		}

		_, err := os.Stat(string(os.PathSeparator) + path.Join(append(pathParts, ".smelt", "repository")...))
		if err == nil {
			break
		}

		pathParts = pathParts[:len(pathParts)-1]
	}

	data, err := os.ReadFile(string(os.PathSeparator) + path.Join(append(pathParts, ".smelt", "repository")...))
	if err != nil {
		return nil, err
	}

	repoConfig := &RepositoryConfig{}
	err = json.Unmarshal(data, repoConfig)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Name: repoConfig.Name,
		Root: path.Join(pathParts...),
	}, nil
}
