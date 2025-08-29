package enumerable

type signedIntegersNumbers interface {
	int | int8 | int16 | int32 | int64
}

type unsignedIntegerNumbers interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type floatingPointNumbers interface {
	float32 | float64
}

type allNumber interface {
	signedIntegersNumbers | unsignedIntegerNumbers | floatingPointNumbers
}
