package fmt

const (
	commaSpaceString  = ", "
	nilAngleString    = "<nil>"
	nilParenString    = "(nil)"
	nilString         = "nil"
	mapString         = "map["
	percentBangString = "%!"
	misingString      = "(MISSING)"
	badIndexString    = "(BADINDEX)"
	panicString       = "(PANIC="
	extraString       = "%!(EXTRA "
	badWidthString    = "%!(BADWIDTH)"
	badPrecString     = "%!(BADPREC)"
	noVerbString      = "%!(NOVERB)"
	invReflectString  = "<invalid reflect.Value>"
)

type State interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, err error)
	Precision() (prec int, err error)
	Flag(c int) bool
}

type Formatter interface {
	Format(f State, c rune)
}

type Stringer interface {
	String() string
}

type GoString interface {
	GoString() string
}

type buffer []byte

func (b *buffer) Write(p []byte) {
	*b = append(*b, p...)
}

func (b *buffer) WriteString(s string) {
	*b = append(*b, s...)
}

func (b *buffer) WriteByte(c byte) {
	*b = append(*b, c)
}

func (b *buffer) WriteRune(r rune) {
	// TODO: 算法没看懂
}
