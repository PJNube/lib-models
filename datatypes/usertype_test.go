package datatypes

import "testing"

func TestUserTypeValidate(t *testing.T) {
	for _, valid := range []UserType{"", UserTypeInteractive, UserTypeCloudRepresentative} {
		v := valid
		if err := v.Validate(); err != nil {
			t.Errorf("UserType %q should be valid: %v", valid, err)
		}
	}
	invalid := UserType("Robot")
	if err := invalid.Validate(); err == nil {
		t.Error("UserType Robot should be invalid")
	}
}

func TestUserTypeIsInteractive(t *testing.T) {
	if !UserType("").IsInteractive() {
		t.Error("zero value must count as interactive")
	}
	if !UserTypeInteractive.IsInteractive() {
		t.Error("Interactive must be interactive")
	}
	if UserTypeCloudRepresentative.IsInteractive() {
		t.Error("CloudRepresentative must not be interactive")
	}
}
