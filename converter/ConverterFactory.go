package converter

type ConverterFactory struct{
	md5Cnv *Md5Converter
}

func (f ConverterFactory) GetConverter(algorithm string) Converter {
	switch algorithm {
	case "md5":
		return new(Md5Converter)
	default:
		return nil
	}
}