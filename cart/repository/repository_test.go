package repository_test

import "testing"

// integration test cases
func TestRepository(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// connect to real caching system - redis
	// execute repository methods
	// asssert the result
}
