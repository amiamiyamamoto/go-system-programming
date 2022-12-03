type StringFuture struct {
	receiver chan string
	cache    string
}

func NewStringFuture() (*StringFuture, chan string) {
	f := &StringFuture{
		receiver: make(chan string)
	}
	return f, f.receiver
}

func (f *StringFuture) Get() string {
	r, ok := <- f.receiver
	if ok {
		close(f.receiver)
		f.cache = r
	}
}

func (f * StringFuture) close() {
	close(f.receiver)
}