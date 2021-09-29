package exception

const ()

var (
	codemap = map[int]string{}
)

type Err struct {
	code int
	Msg  string
}

func (e *Err) Error() string {
	return codemap[e.code]
}
