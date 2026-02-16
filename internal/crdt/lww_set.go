package crdt

import "time"

// Char represents a single character with metadata for conflict resolution.
type Char struct {
	ID        string    `json:"id"`
	Value     string    `json:"value"`
	Position  float64   `json:"position"`
	Timestamp time.Time `json:"timestamp"`
	Deleted   bool      `json:"deleted"`
}

// GeneratePosition finds a value between the 'before' and 'after' positions.
// If inserting at the very beginning, 'before' is 0.
// If inserting at the very end, 'after' is a high upper bound.
func GeneratePosition(before, after float64) float64 {
	if after == 0 {
		return before + 1.0 // Appending to the end
	}
	return (before + after) / 2.0 // Inserting in between
}
