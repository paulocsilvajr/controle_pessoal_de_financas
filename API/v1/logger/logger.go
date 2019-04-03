package logger

import (
	"controle_pessoal_de_financas/API/v1/helper"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ServeHTTPAndLog(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Access-Control-Allow-Origin", "*")

			start := time.Now()

			inner.ServeHTTP(w, r)

			helper.CriarDiretorioSeNaoExistir("logs")

			msg := fmt.Sprintf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)

			GeraLogFS(msg, start)
		})
}

func GeraLogFS(msg string, startTime time.Time) {
	// log em arquivo
	nomeArquivo := fmt.Sprintf("logs/%04d%02d%02d.log", startTime.Year(), startTime.Month(), startTime.Day())
	arquivo, err := os.OpenFile(nomeArquivo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Erro ao abrir arquivo de log[%s]", err)
	}
	defer arquivo.Close()

	// Saída múltipla: arquivo(logs/....log) e tela(Stdout)
	multiplaSaida := io.MultiWriter(os.Stdout, arquivo)
	log.SetOutput(multiplaSaida)

	// [f&s]: File e Screen
	log.Printf(
		"[f&s] %s", msg,
	)
}
