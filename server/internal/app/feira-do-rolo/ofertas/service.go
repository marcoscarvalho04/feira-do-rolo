package ofertas

import (
	"gorm.io/gorm"
	"net/http"
)

type OfertaDTO struct {
	Id               int32  `json:"id"`
	TituloProduto    string `json:"tituloProduto"`
	DescricaoProduto string `json:"descricaoProduto"`
	Contato          string `json:"contato"`
}

type OfertaService struct {
	OfertaRepository *ofertaRepository
	OfertaDTO        *OfertaDTO
}

func NewOfertaService(db *gorm.DB) *OfertaService {
	return &OfertaService{
		OfertaRepository: newOfertaRepository(db),
		OfertaDTO:        &OfertaDTO{},
	}
}

func (s OfertaService) Migrate() error {
	return s.OfertaRepository.Migrate()
}

func (s OfertaService) Save(ofertaDTO OfertaDTO) (OfertaDTO, error) {
	s.OfertaRepository.Model = s.convertDTOToModel(ofertaDTO)
	ofertaModelResponse := ofertaModel{}
	ofertaModelResponse, err := s.OfertaRepository.save()
	if err != nil {
		return ofertaDTO, err
	}
	return s.convertModelToDTO(ofertaModelResponse), nil
}

func (s OfertaService) GetPaginated(r *http.Request) ([]OfertaDTO, error) {
	result := make([]OfertaDTO, 0)
	resultModel, err := s.OfertaRepository.GetPaginate(r)
	if err != nil {
		return result, err
	}
	for _, value := range resultModel {
		result = append(result, s.convertModelToDTO(value))
	}
	return result, nil
}

func (s OfertaService) convertDTOToModel(dto OfertaDTO) ofertaModel {
	return ofertaModel{
		Contato:          dto.Contato,
		Id:               dto.Id,
		DescricaoProduto: dto.DescricaoProduto,
		TituloProduto:    dto.TituloProduto,
	}
}
func (s OfertaService) convertModelToDTO(model ofertaModel) OfertaDTO {
	return OfertaDTO{
		Id:               model.Id,
		TituloProduto:    model.TituloProduto,
		Contato:          model.Contato,
		DescricaoProduto: model.DescricaoProduto,
	}
}
