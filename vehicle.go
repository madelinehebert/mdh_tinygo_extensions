package mdh_tinygo_extensions

import "machine"

// Vehicle struct - a vehicle is both ports on a motor controller
type Vehicle struct {
	M0 Motor
	M1 Motor
}

// Vehicle Functions

// Vehicle function to turn right - assumes m0 is left motor and m1 is right motor
func (v Vehicle) TurnRight() {
	v.M0.SetDirection(v.M0.ForwardDirection)
	v.M1.SetDirection(!v.M1.ForwardDirection)
}

// Vehicle function to turn right - assumes m0 is left motor and m1 is right motor
func (v Vehicle) TurnLeft() {
	v.M0.SetDirection(!v.M0.ForwardDirection)
	v.M1.SetDirection(v.M1.ForwardDirection)
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
	v.M0.SetDirection(v.M0.ForwardDirection)
	v.M1.SetDirection(v.M1.ForwardDirection)

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

// Vehicle function to move a pair of motors in the forwards direction
func (v Vehicle) GoForwards() {
	v.SetSpeed(255)
	v.SetDirection(v.M0.ForwardDirection, v.M1.ForwardDirection)
}

// Vehicle function to move a pair of motors in the backwards direction
func (v Vehicle) GoBackwards() {
	v.SetSpeed(255)
	v.SetDirection(!v.M0.ForwardDirection, !v.M1.ForwardDirection)
}

// Vehicle function to stop a pair of motors
func (v Vehicle) Stop() {
	v.M0.Stop()
	v.M1.Stop()
}

// Vehicle function to start a pair of stopped motors
func (v Vehicle) Start() {
	v.M0.Start()
	v.M1.Start()
}

// Vehicle function to set speed
func (v Vehicle) SetSpeed(speed uint32) {
	v.M0.SetSpeed(speed)
	v.M1.SetSpeed(speed)
}
