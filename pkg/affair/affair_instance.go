package affair

var affair *Affair

func init() {
	affair = NewAffair()
}

func Setup(config map[string]interface{}) error {
	return affair.Setup(config)
}

func Send(cmd []byte) ([]byte, error) {
	return affair.Send(cmd)
}
