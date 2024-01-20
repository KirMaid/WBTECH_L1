/*Реализовать паттерн «адаптер» на любом примере. */
package main

import "fmt"

// Интерфейс Mac
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Интерфейс Windows
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// Адаптер для Windows
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

// Клиентский код
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// Интерфейс Computer
type Computer interface {
	InsertIntoLightningPort()
}

// Паттерн "Адаптер" позволяет объектам с несовместимыми интерфейсами работать вместе.
// Адаптер действует как обертка между двумя объектами.
// Он перехватывает вызовы для одного объекта и преобразует их в формат и интерфейс, которые распознает второй объект.
func main() {
	client := &Client{}
	mac := &Mac{}
	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{windowMachine: windows}

	client.InsertLightningConnectorIntoComputer(mac)
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
