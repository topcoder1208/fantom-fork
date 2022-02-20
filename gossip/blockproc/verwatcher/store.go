package verwatcher

import (
	"github.com/MugamboBC/mugambo-base/kvdb"

	"github.com/topcoder1208/fantom-fork/logger"
)

// Store is a node persistent storage working over physical key-value database.
type Store struct {
	mainDB kvdb.Store

	logger.Instance
}

// NewStore creates store over key-value db.
func NewStore(mainDB kvdb.Store) *Store {
	s := &Store{
		mainDB:   mainDB,
		Instance: logger.MakeInstance(),
	}

	return s
}
