package lexorank

import (
	"fmt"
	"strconv"
	"strings"
)

type LexoRank struct {
	Bucket int
	Value  string
}

const (
	DEFAULT_BUCKET = 1
)

var (
	minValue = "a"
	maxValue = "z"
)

func Min() LexoRank {
	return LexoRank{
		Bucket: DEFAULT_BUCKET,
		Value:  minValue,
	}
}

func Max() LexoRank {
	return LexoRank{
		Bucket: DEFAULT_BUCKET,
		Value:  maxValue,
	}
}

func Middle() LexoRank {
	return LexoRank{
		Bucket: DEFAULT_BUCKET,
		Value:  "h",
	}
}

func (lr LexoRank) String() string {
	return fmt.Sprintf("%d|%s", lr.Bucket, lr.Value)
}

func Parse(str string) (LexoRank, error) {
	parts := strings.SplitN(str, "|", 2)
	if len(parts) != 2 {
		return LexoRank{}, fmt.Errorf("invalid lexorank: %s", str)
	}

	bucket, err := strconv.Atoi(parts[0])
	if err != nil {
		return LexoRank{}, fmt.Errorf("invalid bucket: %s", parts[0])
	}

	if bucket < 0 || bucket > 2 {
		return LexoRank{}, fmt.Errorf("bucket out of range: %d", bucket)
	}

	value := parts[1]

	return LexoRank{Bucket: bucket, Value: value}, nil
}

const ALPHABET_SIZE = 26

// intPow computes base^exp for non-negative integers.
func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func Between(left, right *LexoRank) (LexoRank, error) {
	if left == nil && right == nil {
		return Middle(), nil
	} else if left == nil {
		min := Min()
		left = &min
	} else if right == nil {
		max := Max()
		right = &max
	}

	if left.Bucket != right.Bucket {
		return LexoRank{}, fmt.Errorf("left and right ranks must be in the same bucket")
	}

	leftValue := left.Value
	rightValue := right.Value

	if left.Value >= right.Value {
		return LexoRank{}, fmt.Errorf("left rank must be lower than right rank")
	}

	// Make ranks equal length
	for len(leftValue) < len(rightValue) {
		leftValue += "a"
	}
	for len(rightValue) < len(leftValue) { // Corrected condition
		rightValue += "a"
	}

	firstCodes := []byte(leftValue)
	secondCodes := []byte(rightValue)

	// Compute the difference of secondRank - firstRank in base 26
	difference := 0
	length := len(firstCodes)

	for i := length - 1; i >= 0; i-- {
		f := int(firstCodes[i])
		s := int(secondCodes[i])

		// If s < f, borrow from the previous position
		if s < f {
			s += ALPHABET_SIZE
			// Decrement the previous code (assuming i > 0)
			secondCodes[i-1] = byte(int(secondCodes[i-1]) - 1)
		}

		// Add to difference: (s - f) * 26^( length - i - 1 )
		power := intPow(ALPHABET_SIZE, length-i-1)
		difference += (s - f) * power
	}

	// If difference is too small, just append a mid character (like 'n')
	if difference <= 1 {
		midChar := 'a' + (ALPHABET_SIZE / 2) // 'a' + 13 = 'n'
		return LexoRank{Bucket: left.Bucket, Value: leftValue + string(midChar)}, nil
	}

	// Otherwise, we take half the difference and encode it in base 26 from the end
	difference /= 2

	var newElementBuilder strings.Builder
	offset := 0

	for i := 0; i < length; i++ {
		// Extract digit in base-26 for the i-th place (from right)
		digit := difference / intPow(ALPHABET_SIZE, i) % ALPHABET_SIZE

		// Start from the end of firstRank, add digit + offset
		idx := (length - 1) - i
		newCode := int(firstCodes[idx]) + digit + offset

		offset = 0
		// If newCode > 'z', wrap around
		if newCode > int('z') {
			offset++
			newCode -= ALPHABET_SIZE
		}

		// Prepend character (we'll reverse at the end, or build in reverse)
		newElementBuilder.WriteByte(byte(newCode))
	}

	// The builder has the new rank in reverse order, so reverse it
	newElementBytes := []byte(newElementBuilder.String())
	for i, j := 0, len(newElementBytes)-1; i < j; i, j = i+1, j-1 {
		newElementBytes[i], newElementBytes[j] = newElementBytes[j], newElementBytes[i]
	}

	return LexoRank{
		Bucket: left.Bucket,
		Value:  string(newElementBytes),
	}, nil

}
