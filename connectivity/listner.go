package connectivity

type Listener interface {
	GenNewEvent() string
}

type listener struct {

}

func newListener() *listener  {
	return &listener{

	}
}

func NewListener() Listener {
	return newListener()
}

func (l *listener) GenNewEvent() string {
	return ""
}