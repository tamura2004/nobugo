package step

import "testing"

func TestStepString(t *testing.T) {
	cases := []struct {
		step Step
		want string
	}{
		{PREPARE, "準備フェイズ"},
		{MARCH, "行軍フェイズ"},
		{EMPLOY, "調略フェイズ"},
		{BATTLE, "合戦フェイズ"},
		{CHECK, "終了フェイズ"},
	}

	for _, c := range cases {
		got := c.step.String()
		if got != c.want {
			t.Errorf("bad step string, got %v want %v", got, c.want)
		}
	}
}
