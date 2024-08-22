package mdh_tinygo_extensions

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

	if err := v.M0.ConfigureEverything(); err != nil {
		return err
	}

	if err := v.M1.ConfigureEverything(); err != nil {
		return err
	}

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
