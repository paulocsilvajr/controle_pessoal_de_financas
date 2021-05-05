package logger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/helper"
)

const dirLOG = "logs"

// ServeHTTPAndLog retorna o handler que serve/disponibiliza o HTTP para API
func ServeHTTPAndLog(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Access-Control-Allow-Origin", "*")

			start := time.Now()

			inner.ServeHTTP(w, r)

			// IP:Porta	MétodoHTTP	Rota	NomeRota	Duração

			mascara := strings.Join([]string{
				"%s",
				"%s",
				"%s",
				"%s",
				"%s",
			},
				"\t")
			msg := fmt.Sprintf(
				mascara,
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				name,
				time.Since(start),
			)

			GeraLogFS(msg, start)
		})
}

// GeraLogFS gera um arquivo de log baseado na mensagem(msg) e data/hora informada. Logs são gerados no formato yyyyMMdd.log
func GeraLogFS(msg string, startTime time.Time) {
	dirAbsLOG := helper.CriarDiretorioAbs(dirLOG)

	// log em arquivo
	nomeArquivo := fmt.Sprintf("%s/%04d%02d%02d.log", dirAbsLOG, startTime.Year(), startTime.Month(), startTime.Day())
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
