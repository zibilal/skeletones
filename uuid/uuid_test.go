package uuid

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestGenerateID(t *testing.T) {
	t.Log("Testing id")
	{
		id := GenerateID()

		strId := string(id)

		t.Log("Successfully convert to type string -->", strId)
		if id.Valid() {
			t.Logf("%s expected %s is valid", success, id)
		} else {
			t.Errorf("%s expected %s is valid, got %b", failed, id, id.Valid())
		}

		idNotValid := ID("12344577877999")
		if !idNotValid.Valid() {
			t.Logf("%s expected %s is not valid", success, idNotValid)
		} else {
			t.Errorf("%s expected %s is not valid got %b", failed, idNotValid, idNotValid.Valid())
		}
	}
}
