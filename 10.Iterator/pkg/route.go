package pkg

//У данного файла мы будем определять структуру маршрута, интерфейс итератора, потом структуру массива с обходом элементов и также
//добавим функционал по выбору следующего элемента и проверке того, можно ли перейти к следующему элементу.

type Route struct {
	Name       string
	TravelTime int
}

type Iterator interface {
	HasNext() bool //Признак того, есть ли следующий элемент в нашей коллекции.
	GetNext() *Route
}

type Routes struct {
	index  int     //Чтобы определять на каком шаге находится наш итератор
	Routes []Route //Массив маршрутов, который мы будем обходить
}

func (r *Routes) HasNext() bool {
	return r.index < len(r.Routes)
}

func (r *Routes) GetNext() *Route {
	defer func() {
		r.index++
	}()
	if r.HasNext() {
		return &r.Routes[r.index]
	}
	return nil
}
