package ofertas

import (
	"testing"
)

const (
	ERR_EXPECTED_ERROR = "expected error, got nothing"
)

func TestFieldsEmpty(t *testing.T) {

	modelToValidate := OfertaModel{}
	err := modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	modelToValidate.Contato = "test contato"
	err = modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	modelToValidate.DescricaoProduto = "test Description"
	err = modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	t.Logf("TestFieldsEmpty passed")
}

func TestFieldsMoreThanTwoThousandLen(t *testing.T) {
	modelToValidate := OfertaModel{}
	modelToValidate.TituloProduto = "test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  "

	err := modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	modelToValidate.TituloProduto = ""
	modelToValidate.DescricaoProduto = "test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  "

	err = modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	modelToValidate.DescricaoProduto = ""
	modelToValidate.Contato = "test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  test big string test big string  test big string  test big string  test big string  test big string  test big string  test big string  "

	err = modelToValidate.Save()
	if err == nil {
		t.Errorf(ERR_EXPECTED_ERROR)
		return
	}

	t.Logf("TestFieldsMoreThanTwoThousandLen passed")
}

func TestSaveSucessfully(t *testing.T) {

}
