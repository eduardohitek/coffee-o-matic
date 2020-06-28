// Package p contains an HTTP Cloud Function.
package coffe

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func gerarNumeroRandomico(size int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(size)
}

func retornarPropriedades() []interface{} {
	ret := make([]interface{}, 7)
	torra := []string{"clara", "média", "escura"}
	corpo := []string{"baixo", "médio", "alto"}
	acidez := []string{"cítrica", "málica", "lática", "fosfórica", "acética"}
	aroma := []string{"florado", "cítrico", "achocolatado", "frutado"}
	docura := []string{"nula", "baixa", "alta"}
	amargor := []string{"leve", "moderado", "forte", "muito forte"}
	finalizacao := []string{"chocolate ao leite", "chocolate amargo", "chocolate meio amargo"}
	propriedades := [][]string{torra, corpo, acidez, aroma, docura, amargor, finalizacao}

	for i := 0; i < len(propriedades); i++ {
		propriedade := propriedades[i]
		rn := gerarNumeroRandomico(len(propriedade))
		ret[i] = propriedade[rn]
	}
	return ret

}

// CoffeeReview prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func CoffeeReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	propriedades := retornarPropriedades()
	retorno := fmt.Sprintf("Um café de Torra %s e de corpo %s, com uma acidez %s porém possui um aroma %s, docura %s e amargor %s. Sugestão de finalização seria um %s.",
		propriedades...)
	fmt.Fprint(w, retorno)
}
