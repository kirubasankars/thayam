package maps

import (
	"testing"
)

func TestRoute0(t *testing.T) {
	m := NewMap7()
	n, _:= m.GetNextPosition(0, 2, 12)
	if n != 14 {
		t.Log("expected 14, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(0, 14, 12)
	if n != 24 {
		t.Log("expected 24, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(0, 24, 12)
	if n != 36 {
		t.Log("expected 36, got", n)
		t.Fail()
	}

	n, _= m.GetNextPosition(0, 36, 12)
	if n != 48 {
		t.Log("expected 48, got", n)
		t.Fail()
	}
}

func TestRoute1(t *testing.T) {
	m := NewMap7()
	n, _:= m.GetNextPosition(1, 8, 12)
	if n != 20 {
		t.Log("expected 20, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(1, 20, 6)
	if n != 2 {
		t.Log("expected 2, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(1, 2, 6)
	if n != 36 {
		t.Log("expected 36, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(1, 36, 12)
	if n != 32 {
		t.Log("expected 32, got", n)
		t.Fail()
	}

	n, _= m.GetNextPosition(1, 32, 12)
	if n != 48 {
		t.Log("expected 48, got", n)
		t.Fail()
	}
}

func TestRoute2(t *testing.T) {
	m := NewMap7()
	n, _:= m.GetNextPosition(2, 14, 12)
	if n != 2 {
		t.Log("expected 2, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(2, 2, 12)
	if n != 32 {
		t.Log("expected 32, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(2, 32, 12)
	if n != 28 {
		t.Log("expected 28, got", n)
		t.Fail()
	}

	n, _= m.GetNextPosition(2, 28, 12)
	if n != 48 {
		t.Log("expected 48, got", n)
		t.Fail()
	}
}

func TestRoute3(t *testing.T) {
	m := NewMap7()
	n, _:= m.GetNextPosition(3, 20, 12)
	if n != 8 {
		t.Log("expected 8, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(3, 8, 12)
	if n != 28 {
		t.Log("expected 28, got", n)
		t.Fail()
	}
	n, _= m.GetNextPosition(3, 28, 12)
	if n != 24 {
		t.Log("expected 28, got", n)
		t.Fail()
	}

	n, _= m.GetNextPosition(3, 24, 12)
	if n != 48 {
		t.Log("expected 48, got", n)
		t.Fail()
	}
}
