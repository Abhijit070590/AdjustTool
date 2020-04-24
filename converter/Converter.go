package converter

type Converter interface {
	GetHash(src []byte) string
}
