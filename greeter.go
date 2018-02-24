package greeter

// Greeter says hello, and tells the weather
type Greeter interface {
	SayHello() string
	TellWeather() string
}
