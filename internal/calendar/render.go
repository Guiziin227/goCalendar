package calendar

import (
	"fmt"
	"strings"
	"time"
)

const (
	mesLargura  = 27
	ansiInverte = "\033[7m"
	ansiReseta  = "\033[0m"
)

// RenderMes transforma um MesInfo em 8 linhas
func RenderMes(info MesInfo, hoje time.Time) []string {

	linhas := make([]string, 0, 8)

	titulo := fmt.Sprintf("%s %d", info.Mes.String(), info.Ano)
	linhas = append(linhas, centerText(titulo, mesLargura))

	linhas = append(linhas, "Dom Seg Ter Qua Qui Sex Sab")

	var row strings.Builder
	row.WriteString(strings.Repeat("    ", info.PrimeiroDiaSemana))

	col := info.PrimeiroDiaSemana
	for dia := 1; dia <= info.DiaNoMes; dia++ {
		diaStr := fmt.Sprintf("%3d", dia)

		if info.Ano == hoje.Year() && info.Mes == hoje.Month() && dia == hoje.Day() {
			diaStr = ansiInverte + diaStr + ansiReseta
		}

		if col == 6 {
			row.WriteString(diaStr)
		} else {
			row.WriteString(diaStr + " ")
		}
		col++
		if col == 7 {
			linhas = append(linhas, row.String())
			row.Reset()
			col = 0
		}
	}

	if col > 0 {
		missingCols := 7 - col
		row.WriteString(strings.Repeat("    ", missingCols-1) + "   ")
		linhas = append(linhas, row.String())
	}

	linhaVazia := strings.Repeat(" ", mesLargura)
	for len(linhas) < 8 {
		linhas = append(linhas, linhaVazia)
	}

	return linhas

}

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	leftPad := (width - len(text)) / 2
	rightPad := width - len(text) - leftPad
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}

const espacoColunas = "   " // 3 espaços separando os meses na horizontal

// RenderAno organiza os 12 meses de um ano em um grid 3x4.
func RenderAno(ano int, hoje time.Time) []string {
	var linhasAno []string

	// 1. Cabeçalho geral do ano centralizado na largura total:
	// (3 meses * 27 chars) + (2 espaços entre colunas * 3 chars) = 87 caracteres
	larguraTotal := (mesLargura * 3) + (len(espacoColunas) * 2)
	cabecalhoAno := fmt.Sprintf("%d", ano)
	linhasAno = append(linhasAno, centerText(cabecalhoAno, larguraTotal))
	linhasAno = append(linhasAno, "") // Linha em branco abaixo do ano

	// 2. Percorre os 4 trimestres (linhas de meses: 0, 1, 2, 3)
	for trimestre := range 4 {
		var bloco [3][]string

		// Renderiza os 3 meses deste trimestre
		for col := range 3 {
			numMes := time.Month(trimestre*3 + col + 1)
			info := GetMesInfo(ano, numMes) // Ajuste o nome caso sua função no calendar.go tenha outro nome
			bloco[col] = RenderMes(info, hoje)
		}

		// Junta linha a linha (de 0 a 7) os 3 meses lado a lado
		for idxLinha := range 8 {
			linhaCombinada := bloco[0][idxLinha] + espacoColunas + bloco[1][idxLinha] + espacoColunas + bloco[2][idxLinha]
			linhasAno = append(linhasAno, linhaCombinada)
		}
	}

	return linhasAno
}
