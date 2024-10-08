package mdh_tinygo_extensions

import (
	"machine"
)

// Constants needed to control direction
const (
	FORWARDS  bool = true
	RIGHT     bool = true
	MOTOR0    bool = true //Left Motor
	BACKWARDS bool = false
	LEFT      bool = false
	MOTOR1    bool = false //Right motor
)

// Motor struct - represents a single motor port on a motor controllers
type Motor struct {
	BrakePin         machine.Pin //The motor's brake pin
	DirectionPin     machine.Pin //The motor's direction pin
	SpeedPin         machine.Pin //The motor's speed pin
	PwmPin           machine.PWM //The motor's PWM timer
	PwmCh            uint8       //The PWM timer's channel
	ForwardDirection bool        //The direction to use as forward; Useful to correct polarity
}

// Motor automagical configuration function for everything needed to get a single motor up and running ASAP
func (m Motor) ConfigureEverything() error {
	m.ConfigureAnalog()
	if err := m.ConfigurePWM(); err != nil {
		return err
	}

	if ch, err := m.PwmPin.Channel(m.SpeedPin); err != nil {
		return err
	} else {
		m.PwmCh = ch
	}

	m.SetDirection(m.ForwardDirection)
	m.SetSpeed(255)

	return nil
}

// Motor automagical configuration function for just the pins
func (m Motor) ConfigurePins() error {
	m.ConfigureAnalog()

	if err := m.ConfigurePWM(); err != nil {
		return err
	}

	if ch, err := m.PwmPin.Channel(m.SpeedPin); err != nil {
		return err
	} else {
		m.PwmCh = ch
	}

	return nil
}

// Motor autoconfiguration function for analog pins
func (m Motor) ConfigureAnalog() {
	m.DirectionPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.SpeedPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.BrakePin.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Motor autoconfiguration function for pwm pins
func (m Motor) ConfigurePWM() error {
	if err := m.PwmPin.Configure(machine.PWMConfig{}); err != nil {
		return err
	} else {
		return nil
	}
}

// Motor autoconfiguration function for pwm channels
/*
func (m Motor) ConfigurePWMChannel() error {
	if ch, err := m.PwmPin.Channel(m.SpeedPin); err != nil {
		return err
	} else {
		m.PwmCh = ch
		return nil
	}
}
*/
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

// Motor function to stop a motor
func (m Motor) Stop() {
	m.BrakePin.High()
	m.SetSpeed(0)
}

// Motor function to start a motor from a stopped state
func (m Motor) Start() {
	m.BrakePin.Low()
	m.SetSpeed(255)
}
