package timestamp

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var ts Timestamp
	if err := json.Unmarshal([]byte("12345"), &ts); err != nil {
		t.Fatalf("Failure unmarshalling unquoted number: %s", err)
	}
	if ts != 12345 {
		t.Fatalf("t = %f instead of 12345", ts)
	}
	if err := json.Unmarshal([]byte("\"12345\""), &ts); err != nil {
		t.Fatalf("Failure unmarshalling quoted number: %s", err)
	}
	if ts != 12345 {
		t.Fatalf("t = %f instead of 12345", ts)
	}
}
