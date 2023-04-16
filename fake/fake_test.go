package fake

import (
	"main/util"
	"testing"
)

func TestGenerateFakeApplicant(t *testing.T) {
	applicant := GenerateFakeApplicant()
	util.PrettyPrint(t, applicant)
}

func TestGetRandomApplicant(t *testing.T) {
	applicants := []Applicant{
		GenerateFakeApplicant(),
		GenerateFakeApplicant(),
		GenerateFakeApplicant(),
	}

	for i := 0; i < 5; i++ {
		a, index := GetRandomApplicant(applicants)
		t.Log(index)
		util.PrettyPrint(t, a)
	}
}
