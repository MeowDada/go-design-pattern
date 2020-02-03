package adapter

// Target is usually an invoking interface. Request() method is the target
// method that we want to invoke. 
type Target interface {
	Request() string
}

// Adaptee is usually an adaptee interface. SpecificRequest() method is the
// corresponding method to Target.Request(). We use an adapter to adapt this
// method.
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl is a concret object implementing Adaptee interface.
type AdapteeImpl struct {}

// SpecificRequests implements Adaptee interface for the concret Adaptee instance.
func (a *AdapteeImpl) SpecificRequest() string {
	return "Adaptee's method"
}

// NewAdaptee is a factory method to create Adaptee instance.
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

// Adapter is a middleware to connect adaptee and target.
type Adapter struct {
	Adaptee
}

// NewAdapter is a factory method for instantiating an adapter.
func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter {
		Adaptee: adaptee,
	}
}

// Request implements Target interface for the adapter.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}






