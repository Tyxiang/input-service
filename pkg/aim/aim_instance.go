package aim

var aim *Aim

func init() {
	aim = NewAim()
}

func Setup(config map[string]interface{}) error {
	return aim.Setup(config)
}

func GetSignal(channel int) (float64, error) {
	return aim.GetSignal(channel)
}
