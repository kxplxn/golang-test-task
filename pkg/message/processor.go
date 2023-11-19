package message

type Processor struct{}

func NewProcessor() Processor { return Processor{} }

func (p Processor) Proccess() error {
	// rmq, err := rabbitmq.Get()
	// if err != nil {
	// 	return fmt.Errorf("error getting rabbitmq channel: %s", err)
	// }
	return nil
}
