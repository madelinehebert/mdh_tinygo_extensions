package mdh_tinygo_extensions

import (
	"machine"
)

// Constants needed to control direction
const (
	FORWARDS  bool = true
	RIGHT     bool = true
	MOTOR0    bool = true
	BACKWARDS bool = false
	LEFT      bool = false
	MOTOR1    bool = false
)

// Motor struct - represents a single motor port on a motor controllers
type Motor struct {
	BrakePin     machine.Pin
	DirectionPin machine.Pin
	SpeedPin     machine.Pin
	PwmPin       machine.PWM
	PwmCh        uint8
}

// Vehicle struct - a vehicle is both ports on a motor controller
type Vehicle struct {
	M0 Motor
	M1 Motor
}

// Motor automagical configuration function for everything needed to get a single motor up and running ASAP
func (m Motor) ConfigureEverything() error {
	m.ConfigureAnalog()
	if err := m.ConfigurePWM(); err != nil {
		return err
	}
	if err := m.ConfigurePWMChannel(); err != nil {
		return err
	}

	m.SetDirection(FORWARDS)
	m.SetSpeed(255)

	return nil
}

// Motor automagical configuration function for just the pins
func (m Motor) ConfigureAll() error {
	m.ConfigureAnalog()

	if err := m.ConfigurePWM(); err != nil {
		return err
	}

	if err := m.ConfigurePWMChannel(); err != nil {
		return err
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
func (m Motor) ConfigurePWMChannel() error {
	if ch, err := m.PwmPin.Channel(m.SpeedPin); err != nil {
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

// Motor function to stop a motor
func (m Motor) Stop() {
	m.BrakePin.High()
}

// Motor function to start a motor from a stopped state
func (m Motor) Start() {
	m.BrakePin.Low()
}

// Vehicle Functions

// Vehicle function to turn right - assumes m0 is left motor and m1 is right motor
func (v Vehicle) TurnRight() {
	v.M0.SetDirection(FORWARDS)
	v.M1.SetDirection(BACKWARDS)
}

// Vehicle function to turn right - assumes m0 is left motor and m1 is right motor
func (v Vehicle) TurnLeft() {
	v.M0.SetDirection(BACKWARDS)
	v.M1.SetDirection(FORWARDS)
}

// Vehicle automagical configuration function
func (v Vehicle) ConfigureEverything() error {

	// Configure M0 Direction Pin
	v.M0.DirectionPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure M1 Direction Pin
	v.M1.DirectionPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Configure M0 PWM pin
	err := v.M0.PwmPin.Configure(machine.PWMConfig{})
	if err != nil {
		return err
	}

	// Configure channel on M0
	ch0, err := v.M0.PwmPin.Channel(v.M0.SpeedPin)
	if err != nil {
		return err
	} else {
		v.M0.PwmCh = ch0
	}

	// Configure M1 PWM pin
	err = v.M1.PwmPin.Configure(machine.PWMConfig{})
	if err != nil {
		return err
	}

	// Configure channel on M1
	ch1, err := v.M1.PwmPin.Channel(v.M1.SpeedPin)
	if err != nil {
		return err
	} else {
		v.M1.PwmCh = ch1
	}

	// Set direction for both motors
	v.M0.DirectionPin.High()
	v.M1.DirectionPin.High()

	// Toggle the motors!
	v.M0.PwmPin.Set(ch0, uint32(255))
	v.M1.PwmPin.Set(ch1, uint32(255))

	return nil
}

// Vehicle function to set direction for all motors
func (v Vehicle) SetDirectionAll(direction bool) {
	v.M0.SetDirection(direction)
	v.M1.SetDirection(direction)
}

// Vehicle function to set direction for each motor individually
func (v Vehicle) SetDirection(direction0 bool, direction1 bool) {
	v.M0.SetDirection(direction0)
	v.M1.SetDirection(direction1)
}
