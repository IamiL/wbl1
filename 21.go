package main

import "fmt"

// разъём автомобилей Tesla
type TeslaConnector interface {
	TeslaChargerConnect()
}

// Разъём в автомобиле Tesla.
type Tesla struct {
}

func (s *Tesla) Charge(charger TeslaConnector) {
	fmt.Println("Пробуем зарядить автомобиль Tesla.")
	charger.TeslaChargerConnect()
}

// зарядное устройство Tesla.
type TeslaCharger struct {
}

func (p *TeslaCharger) TeslaChargerConnect() {
	fmt.Println("ЗУ Тесла подключено.")
}

// стандартное зарядное устройство
type StandartCharger struct {
}

func (p *StandartCharger) StandartChargerConnect() {
	fmt.Println("Стандартное ЗУ подключено.")
}

// Adapter — адаптер евровилки для американской розетки.
type Adapter struct {
	Standart *StandartCharger
}

// адаптер, поддерживающий TeslaConnector-интерфейс.
func (a *Adapter) TeslaChargerConnect() {
	fmt.Println("Адаптер в американской розетке.")
	a.Standart.StandartChargerConnect()
}

func main() {
	tesla := &Tesla{}
	tesla.Charge(&TeslaCharger{})

	standartCharger := &StandartCharger{}
	adapter := &Adapter{
		Standart: standartCharger,
	}

	tesla.Charge(adapter)
}
