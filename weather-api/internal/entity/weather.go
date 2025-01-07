package entity

type Weather struct {
	City  string
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (w *Weather) Convert() {
	w.TempF = w.TempC*1.8 + 32
	w.TempK = w.TempC + 273.15
}
