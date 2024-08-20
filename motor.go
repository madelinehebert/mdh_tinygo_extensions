package mdh_tinygo_extensions

import "machine"

// Constants needed to control direction
const (
	FORWARDS  bool = true
	BACKWARDS bool = false
)

// Motor structs
type motor struct {
	directionPin machine.Pin
	speedPin     machine.Pin
	pwmPin       machine.PWM
	pwmCh        uint8
}

// Motor automagical configuration function
func (m motor) ConfigureAll() error {
	m.ConfigureAnalog()
	m.ConfigurePWM()
	if err := m.ConfigurePWMChannel(); err != nil {
		return err
	}

	return nil
}

// Motor autoconfiguration function for analog pins
func (m motor) ConfigureAnalog() {
	m.directionPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.speedPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Motor autoconfiguration function for pwm pins
func (m motor) ConfigurePWM() {
	pwm := m.pwmPin
	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println(err.Error())
	}
}

// Motor autoconfiguration function for pwm channels
func (m motor) ConfigurePWMChannel() error {
	ch, err := m.pwmPin.Channel(m.speedPin)
	if err != nil {
		println(err.Error())
		return err
	} else {
		m.pwmCh = ch
		return nil
	}
}

// Motor direction function
func (m motor) SetDirection(direction bool) {
	// Fork logic based on input
	if direction {
		//Map High to True (Forwards)
		m.directionPin.High()
	} else {
		//Map Low to False (Backwards)
		m.directionPin.Low()
	}
}

// Motor direction function
func (m motor) GetDirection() bool {
	//Return current value of directionPin
	return m.directionPin.Get()
}

// Motor function to set speed
func (m motor) SetSpeed(speed uint32) {
	m.pwmPin.Set(m.pwmCh, speed)
}
