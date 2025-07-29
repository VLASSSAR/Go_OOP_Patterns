package data

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

// Создадим адаптер
type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	println("Отправка Xml данных!")
}
