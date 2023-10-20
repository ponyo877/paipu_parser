package repository

import (
	// "github.com/EndlessCheng/mahjong-helper/platform/mahjongSoul/internal/dao"
	"github.com/ponyo877/paipu_parser/entity"
)

// PaipuMahjongSoul http repository
type PaipuMahjongSoul struct {
}

// NewPaipuMahjongSoul create new repository
func NewPaipuMahjongSoul() *PaipuMahjongSoul {
	return &PaipuMahjongSoul{}
}

// Get
func (r *PaipuMahjongSoul) GetParsedPaipu(paipuURL string) (entity.ParsedPaipu, error) {
	return entity.ParsedPaipu{}, nil
}
