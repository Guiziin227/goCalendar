package calendar

import (
	"fmt"
	"strings"
)

const mesLargura = 27

// RenderMes transforma um MesInfo em 8 linhas
func RenderMes(info MesInfo) []string {

	linhas := make([]string, 0, 8)

	titulo := fmt.Sprintf("%s %d", info.Mes.String(), info.Ano)
	linhas = append(linhas, centerText(titulo, mesLargura))

	linhas = append(linhas, "Dom Seg Ter Qua Qui Sex Sab")

	var row strings.Builder
	row.WriteString(strings.Repeat("    ", info.PrimeiroDiaSemana))

	col := info.PrimeiroDiaSemana
	for dia := 1; dia <= info.DiaNoMes; dia++ {
		if col == 6 {
			row.WriteString(fmt.Sprintf("%3d", dia))
		} else {
			row.WriteString(fmt.Sprintf("%3d ", dia))
		}
		col++
		if col == 7 {
			linhas = append(linhas, row.String())
			row.Reset()
			col = 0
		}
	}

	if col > 0 {
		padding := fmt.Sprintf("%-*s", mesLargura, row.String())
		linhas = append(linhas, padding)
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
