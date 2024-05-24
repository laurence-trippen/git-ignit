package gitignit

import "testing"

func TestCompleteIgnoreName(t *testing.T) {
	result := CompleteIgnoreName("Go")

	if result != "Go.gitignore" {
		t.Errorf("Un-expected result %s. Expected Go.gitignore as result!", result)
	}
}

func TestMeantToFail(t *testing.T) {
	t.Error("Meant to fail!")
}
