package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user *userRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {

	// if doesFileExist(cfg.Path + cfg.UserFileName) {
	// 	_, err := os.Create(cfg.UserFileName)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}


	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil
}

func (s *Store) CloseDB() {
	s.user.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

// func doesFileExist(fileName string) bool {
// 	_, error := os.Stat(fileName)
// 	return os.IsNotExist(error)
// }
