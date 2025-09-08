package main

	func PassPass(a, b, c *int) {
		*c = *b
		*b = *a
		*a = 0
		*c = *b
		*b = 0
}
