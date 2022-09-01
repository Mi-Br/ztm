package main

import (
	"testing"
)

func TestAddHealth(t *testing.T) {
	testPlayer := MakePlayer("Test player", 100, 100)
	type testcase struct {
		add    uint
		result uint
	}

	cases := make([]testcase, 0, 3)
	cases = append(cases,
		testcase{10, 60},
		testcase{20, 80},
		testcase{30, 100},
	)

	for _, c := range cases {
		add, res := c.add, c.result
		testPlayer.addHealth(add)
		if testPlayer.health != res {
			t.Errorf("TestAddHealth(%v)=%v, expect %v", add, testPlayer.health, res)
			t.Fail()
		}
	}

}

func TestPlayer_applyDamage(t *testing.T) {
	type fields struct {
		health    uint
		maxHealth uint
		energy    uint
		maxEnergy uint
		name      string
	}
	type args struct {
		points uint
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantDead bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				health:    tt.fields.health,
				maxHealth: tt.fields.maxHealth,
				energy:    tt.fields.energy,
				maxEnergy: tt.fields.maxEnergy,
				name:      tt.fields.name,
			}
			if gotDead := p.applyDamage(tt.args.points); gotDead != tt.wantDead {
				t.Errorf("Player.applyDamage() = %v, want %v", gotDead, tt.wantDead)
			}
		})
	}
}
