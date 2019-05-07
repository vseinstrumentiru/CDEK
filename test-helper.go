package cdek

func strLink(s string) *string {
	return &s
}

func intLink(i int) *int {
	return &i
}

func float64Link(f float64) *float64 {
	return &f
}

func boolLink(b bool) *bool {
	return &b
}
