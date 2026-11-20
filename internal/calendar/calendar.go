package calendar

import "time"

type MesInfo struct {
	Ano               int
	Mes               time.Month
	PrimeiroDiaSemana int // 0(domgin) ate 6(sabado)
	DiaNoMes          int
}

func GetMesInfo(ano int, mes time.Month) MesInfo {
	primeiroDia := time.Date(ano, mes, 1, 0, 0, 0, 0, time.Local)
	ultimoDia := time.Date(ano, mes+1, 0, 0, 0, 0, 0, time.Local)

	return MesInfo{
		Ano:               ano,
		Mes:               mes,
		PrimeiroDiaSemana: int(primeiroDia.Weekday()),
		DiaNoMes:          ultimoDia.Day(),
	}
}
