package basketball

import (
	"fmt"
	"strings"
)

// String returns a string representation of the basketball.
func (b *BasketBall) String() string {
	var builder strings.Builder

	builder.WriteString("Size: ")
	builder.WriteString(fmt.Sprintf("%d", b.Size))

	builder.WriteString("\n")

	builder.WriteString("Color: ")
	builder.WriteString(b.Color)

	builder.WriteString("\n")

	builder.WriteString("Weight: ")
	builder.WriteString(fmt.Sprintf("%d", b.Weight))

	return builder.String()
}
