package paipu

import (
	"github.com/ponyo877/paipu_parser/entity"
)

// Reader interface
type Reader interface {
	GetParsedPaipu(paipuURL string) (entity.ParsedPaipu, error)
}

// Writer interface
type Writer interface {
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetParsedPaipu(paipuURL string) (entity.ParsedPaipu, error)
}
