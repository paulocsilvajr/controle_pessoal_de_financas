package controller

import "testing"

func TestPessoaIndex(t *testing.T) {
	tokenPessoaAdmin, _ := getToken("teste01", "123456")
	rota := "/pessoas"
	res, body, err := get(rota, tokenPessoaAdmin)
	if err != nil {
		t.Error(err)
		return
	}

	status := res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}

	tokenPessoaComum, _ := getToken("paulo", "123456")
	res, body, _ = get(rota, tokenPessoaComum)
	status = res.StatusCode
	if status != 200 {
		t.Error(res, string(body))
	}
}
