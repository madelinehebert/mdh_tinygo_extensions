package mdh_tinygo_extensions

import "machine"

// Constants needed to control direction
const (
	FORWARDS  bool = true
	BACKWARDS bool = false
)

// Motor structs
type Motor struct {
	DirectionPin machine.Pin
	SpeedPin     machine.Pin
	PwmPin       machine.PWM
	PwmCh        uint8
}

// Motor automagical configuration function
func (m Motor) ConfigureAll() error {
	m.ConfigureAnalog()
	m.ConfigurePWM()
	if err := m.ConfigurePWMChannel(); err != nil {
		return err
	}

	return nil
}

// Motor autoconfiguration function for analog pins
func (m Motor) ConfigureAnalog() {
	m.DirectionPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.SpeedPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Motor autoconfiguration function for pwm pins
func (m Motor) ConfigurePWM() {
	pwm := m.PwmPin
	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println(err.Error())
	}
}

// Motor autoconfiguration function for pwm channels
func (m Motor) ConfigurePWMChannel() error {
	ch, err := m.PwmPin.Channel(m.SpeedPin)
	if err != nil {
		println(err.Error())
		return err
	} else {
		m.PwmCh = ch
		return nil
	}
}

// Motor direction function
func (m Motor) SetDirection(direction bool) {
	// Fork logic based on input
	if direction {
		//Map High to True (Forwards)
		m.DirectionPin.High()
	} else {
		//Map Low to False (Backwards)
		m.DirectionPin.Low()
	}
}

// Motor direction function
func (m Motor) GetDirection() bool {
	//Return current value of directionPin
	return m.DirectionPin.Get()
}

// Motor function to set speed
func (m Motor) SetSpeed(speed uint32) {
	m.PwmPin.Set(m.PwmCh, speed)
}
