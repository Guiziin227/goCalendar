package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/guiziin227/goCalendar/internal/calendar"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "cal: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	hoje := time.Now()
	var linhas []string

	switch len(args) {
	case 0:
		// Modo 1: sem argumentos -> exibe o mês atual
		info := calendar.GetMesInfo(hoje.Year(), hoje.Month())
		linhas = calendar.RenderMes(info, hoje)

	case 1:
		// Modo 2: 1 argumento -> exibe o ano inteiro (ex: cal 2026)
		ano, err := parseAno(args[0])
		if err != nil {
			return err
		}
		linhas = calendar.RenderAno(ano, hoje)

	case 2:
		// Modo 3: 2 argumentos -> exibe mês e ano específicos (ex: cal 11 2026)
		mes, err := parseMes(args[0])
		if err != nil {
			return err
		}
		ano, err := parseAno(args[1])
		if err != nil {
			return err
		}
		info := calendar.GetMesInfo(ano, mes)
		linhas = calendar.RenderMes(info, hoje)

	default:
		return errors.New("uso: cal [[mes] ano]")
	}

	for _, linha := range linhas {
		fmt.Println(linha)
	}
	return nil
}

func parseMes(s string) (time.Month, error) {
	mes, err := strconv.Atoi(s)
	if err != nil || mes < 1 || mes > 12 {
		return 0, fmt.Errorf("%q não é um mês válido (1..12)", s)
	}
	return time.Month(mes), nil
}

func parseAno(s string) (int, error) {
	ano, err := strconv.Atoi(s)
	if err != nil || ano < 1 || ano > 9999 {
		return 0, fmt.Errorf("%q não é um ano válido (1..9999)", s)
	}
	return ano, nil
}
