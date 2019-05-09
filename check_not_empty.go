package main

func checkNotEmpty(strings ...string) bool {
	for _, s := range strings {
		if s == "" {
			return false
		}
	}
	return true
}
