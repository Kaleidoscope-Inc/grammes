package traversal

// StringStep adds a custom string to the traversal. This is not a gremlin step.
// Added for conveninece when the grammes library does not provide a step.
// Sample usage string_step("count()")
// User is responsible for the correctness of the string.
func (g String) StringStep(traversal string) String {
	g.buffer.Reset()

	g.buffer.WriteString("." + traversal)

	g.string += g.buffer.String()
	return g
}
