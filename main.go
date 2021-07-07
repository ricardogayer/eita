package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var indiceColunaExcluida int = -1
var registro []string
var registros [][]string
var descricaoColunaExcluida string = "CREDICARD_NUMBER"
var diretorioOrigem string = "origem/"
var diretorioDestino string = "destino/"

func main() {

	// Abre o diretório onde os arquivos foram colocados (arquivos csv separados por ";")
	f, err := os.Open(diretorioOrigem)
	if err != nil {
		log.Fatal(err)
	}

	// Coloca os arquivos em um array para tratamento individual
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Para cada arquivo encontrado, chama-se a rotina para remover a coluna
	for _, file := range files {
		fmt.Println(file.Name())
		removeColuna(file.Name())
	}


}

func removeColuna(arquivo string) {

	// Limpeza das variáveis
	indiceColunaExcluida = -1
	registro = nil
	registros = nil

	// Abertura do arquivo de origem
	f, err := os.Open(diretorioOrigem + arquivo)

	if err != nil {
		log.Fatal("erro abrindo o arquivo")
	}

	// Leitura dos campos no formato csv
	r := csv.NewReader(f)
	r.Comma = 59
	r.LazyQuotes = true
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal("erro lendo os registros do arquivo",arquivo)
	}

	// Depois da leitura o arquivo é fechado p/ liberar recursos
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Para cada coluna, é realizado uma análise para verificar se existe a coluna procurado dentro do arquivo
	for _, record := range records {

		for i := 0; i < len(record); i++ {
			if record[i] == descricaoColunaExcluida {
				fmt.Println("Coluna a ser excluída encontrada no arquivo ",arquivo)
				indiceColunaExcluida = i
			}
		}

		for i := 0; i < len(record); i++ {
			if i != indiceColunaExcluida {
				registro = append(registro, record[i])
			}
		}

		registros = append(registros,registro)
		registro = nil

	}


	// Criar um arquivo novo com mesmo nome no diretório destino
	d, err := os.Create(diretorioDestino + arquivo)

	if err != nil {

		log.Fatalln("erro durante a abertura do arquivo", err)
	}

	w := csv.NewWriter(d)
	w.Comma = 59

	for _, registro := range registros {
		if err := w.Write(registro); err != nil {
			log.Fatalln("erro na gravação do arquivo", err)
		}
	}

	w.Flush()
	d.Close()

}


