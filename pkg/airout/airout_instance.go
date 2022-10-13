package airout

var airout *Airout

func init() {
	airout = NewAirout()
}

func Setup(config map[string]interface{}) error {
	return airout.Setup(config)
}

func Start() {
	airout.Start()
	return 
}

func Exit() chan struct{} {
	return airout.Exit()
}
func Err() chan error {
	return airout.Err()
}