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
	if e.message == "" {
		return "sequence contains no elements"
	}
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
	if e.message == "" {
		return "sequence contains more than one element"
	}
	return e.message
}

var (
	// ErrNoElements is the predefined error instance returned when a sequence contains
	//  no elements.
	ErrNoElements = &NoElementsError{}

	// ErrMultipleElements is the predefined error instance returned when a sequence
	// contains more than one element, but exactly one was expected.
	ErrMultipleElements = &MultipleElementsError{}
)

// NewNoElementsError creates a new NoElementsError with a custom message.
// This function is useful when you need to provide additional context
// about why no elements were found in a sequence.
func NewNoElementsError(message string) error {
	return &NoElementsError{message: message}
}

// NewMultipleElementsError creates a new MultipleElementsError with a custom message.
// This function is useful when you need to provide additional context about
// why multiple elements were found when exactly one was expected.
func NewMultipleElementsError(message string) error {
	return &MultipleElementsError{message: message}
}

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
