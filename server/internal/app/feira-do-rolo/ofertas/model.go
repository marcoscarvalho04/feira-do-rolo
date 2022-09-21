package ofertas

import (
	"gorm.io/gorm"
)

type ofertaModel struct {
	gorm.Model
	TituloProduto    string `json:"titulo_produto"`
	DescricaoProduto string `json:"descricao_produto"`
	Contato          string `json:"contato"`
	Id               int32  `json:"id" gorm:"primaryKey"`
}

func (ofertaModel) TableName() string {
	return "offers"
}
