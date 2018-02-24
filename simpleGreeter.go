package greeter

// SimpleGreeter just says hello, and always says that it's rainy.
type SimpleGreeter struct {
	weather  string
	greeting string
}

// NewSimpleGreeter makes a new SimpleGreeter
func NewSimpleGreeter(w string, g string) SimpleGreeter {
	return SimpleGreeter{
		weather:  w,
		greeting: g,
	}
}

// SayHello returns the preset greeting
func (sg SimpleGreeter) SayHello() string {
	return sg.greeting
}

// TellWeather returns the preset weather
func (sg SimpleGreeter) TellWeather() string {
	return sg.weather
}
