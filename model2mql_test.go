package model2mql

import "testing"

func TestConvertor(t *testing.T) {
	type TestStruct struct {
		UseridContains int64  `json:"user_id,omitempty" mql:"user_id"`
		IdGte          int64  `json:"id_contains,omitempty"  mql:"id"`
		TitleNe        string `json:"title_ne,omitempty"  mql:"title"`
		Completed      bool   `json:"completed_eq,omitempty" mql:"completed"`
	}

	c := NewConvertor(TestStruct{})

	response, err := c.Parse(TestStruct{
		UseridContains: 1,
		IdGte:          1,
		TitleNe:        "Test",
		Completed:      false,
	})
	if err != nil {
		t.Fatal(err)
	}

	check := "user_id % 1 and id >= 1 and title != \"Test\" and completed = false"
	if response != check {
		t.Fatalf("Expected: %s\nGot: %s", check, response)
	}

	t.Log(response)
}
