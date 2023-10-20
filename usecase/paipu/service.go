package paipu

import "github.com/ponyo877/paipu_parser/entity"

// Service Paipu usecase
type Service struct {
	repository Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

// GetParsedPaipu get parsed paipu
func (s *Service) GetParsedPaipu(paipuURL string) (entity.ParsedPaipu, error) {
	parsedPaipu, err := s.repository.GetParsedPaipu(paipuURL)
	if err != nil {
		return entity.ParsedPaipu{}, err
	}
	return parsedPaipu, nil
}
