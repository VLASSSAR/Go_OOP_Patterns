package base

// У пк присутствует функционал сканирования, а также определения сканнера.
type PC interface {
	Scan()
	AddScanner(scanner Scanner)
}
