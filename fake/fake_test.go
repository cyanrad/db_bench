package fake

import (
	"encoding/json"
	"testing"
)

// print the contents of the obj
func PrettyPrint(t *testing.T, data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("%s \n", p)
}

func TestGenerateFakeApplicant(t *testing.T) {
	applicant := GenerateFakeApplicant()
	PrettyPrint(t, applicant)
}
