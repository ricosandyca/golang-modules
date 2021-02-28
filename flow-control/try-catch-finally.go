package flowcontrol

// Exception defines exception data type
type Exception interface{}

// FlowControl is used to create flow control
// It has readable flow like the most in other languages
// Try method recovers codes that possibly have panic or throw inside
// Catch method is used to catch error. You need to use is if you want to avoid panic
// Finally method invokes at the end of the process
type FlowControl struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Throw is used to do panic or throw a code
// Throw is just working like panic function
func Throw(e Exception) {
	panic(e)
}

// Do function is used to start running the flow control
// You need to invoke this function
func (flow FlowControl) Do() {
	if flow.Finally != nil {
		defer flow.Finally()
	}

	if flow.Catch != nil {
		defer func() {
			if err := recover(); err != nil {
				flow.Catch(err)
			}
		}()
	}

	flow.Try()
}
