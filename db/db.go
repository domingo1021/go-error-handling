package db

type MockDb struct {
	persons map[string]Person
}

// NewMockDb creates a new instance of MockDb.
func NewMockDb() *MockDb {
	return &MockDb{
		persons: make(map[string]Person),
	}
}
