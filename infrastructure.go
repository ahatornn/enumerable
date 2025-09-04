package enumerable

// NoElementsError represents an error that occurs when an operation expects
// at least one element in a sequence, but the sequence is empty.
type NoElementsError struct {
	message string
}

// Error implements the error interface for NoElementsError.
// This method returns a string representation of the error that describes
// why the operation failed due to an empty sequence.
func (e *NoElementsError) Error() string {
	return e.message
}

// MultipleElementsError represents an error that occurs when an operation expects
// exactly one element in a sequence, but the sequence contains multiple elements.
type MultipleElementsError struct {
	message string
}

// Error implements the error interface for MultipleElementsError.
// This method returns a string representation of the error that describes
// why the operation failed due to multiple elements in a sequence.
func (e *MultipleElementsError) Error() string {
	return e.message
}

var (
	// ErrNoElements is the predefined error instance returned when a sequence contains
	//  no elements.
	ErrNoElements = &NoElementsError{
		message: "sequence contains no elements",
	}

	// ErrMultipleElements is the predefined error instance returned when a sequence
	// contains more than one element, but exactly one was expected.
	ErrMultipleElements = &MultipleElementsError{
		message: "sequence contains more than one element",
	}
)

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
