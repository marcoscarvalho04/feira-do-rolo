package ofertas

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ofertaRepository struct {
	Model ofertaModel
	db    *gorm.DB
}

func newOfertaRepository(db *gorm.DB) *ofertaRepository {
	return &ofertaRepository{
		db:    db,
		Model: ofertaModel{},
	}
}

func (s ofertaRepository) save() (ofertaModel, error) {
	err := s.validateAllFields()
	if err != nil {
		return s.Model, err
	}

	s.Model, err = s.saveToDatabase()
	if err != nil {
		return s.Model, err
	}
	return s.Model, nil
}

func (s ofertaRepository) validateAllFields() error {
	if len(s.Model.Contato) == 0 || len(s.Model.DescricaoProduto) == 0 || len(s.Model.TituloProduto) == 0 {
		return errors.New("fields: Contato, Descricao ou titulo must not be empty")
	}
	if len(s.Model.Contato) > 100 || len(s.Model.DescricaoProduto) > 200 || len(s.Model.TituloProduto) > 200 {
		return errors.New("fiels: Contato, Descricao ou titulo must not be over 200")
	}
	return nil
}

func (s ofertaRepository) saveToDatabase() (ofertaModel, error) {
	err := s.db.Save(&s.Model)
	if err.Error != nil {
		return s.Model, err.Error
	}

	return s.Model, nil
}

func (s ofertaRepository) Migrate() error {
	return s.db.AutoMigrate(&ofertaModel{})
}

func (s ofertaRepository) GetPaginate(r *http.Request) ([]ofertaModel, error) {
	allOffers := make([]ofertaModel, 0)
	err := s.db.Scopes(Paginate(r, s.db)).Model(&ofertaModel{}).Order("id desc").Find(&allOffers)
	if err.Error != nil {
		return allOffers, err.Error
	}
	return allOffers, nil

}

func Paginate(r *http.Request, originalDb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(originalDb *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))

		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("page_size"))

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return originalDb.Offset(offset).Limit(pageSize)
	}

}
