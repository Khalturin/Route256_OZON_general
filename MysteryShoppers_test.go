package main

import "testing"

func TestMysteryShoppers1(te *testing.T) {
	m := 2 + 1
	n := []int{1, 4, 3}
	t := []int{2, 7, 6}

	res := MysteryShoppers(n, t, m)

	if res[0] != 1 {
		te.Error("res[0] expected 1, but ", res[0])
	}
	if res[1] != 3 {
		te.Error("res[1] expected 3, but ", res[1])
	}
	if res[2] != 0 {
		te.Error("res[2] expected 0, but ", res[2])
	}

}

func TestMysteryShoppers2(te *testing.T) {
	m := 1 + 1
	n := []int{1, 4}
	t := []int{2, 7}

	res := MysteryShoppers(n, t, m)

	if res[0] != 1 {
		te.Error("res[0] expected 1, but ", res[0])
	}
	if res[1] != 3 {
		te.Error("res[1] expected 3, but ", res[1])
	}
}

func TestMysteryShoppers3(te *testing.T) {
	m := 8 + 1
	n := []int{4, 2, 9, 2, 4, 7, 9, 14, 5}
	t := []int{7, 4, 8, 7, 6, 3, 14, 9, 2}

	res := MysteryShoppers(n, t, m)

	if res[0] != 3 {
		te.Error("res[0] expected 3, but ", res[0])
	}
	if res[1] != 2 {
		te.Error("res[1] expected 2, but ", res[1])
	}
	if res[2] != 0 {
		te.Error("res[2] expected 0, but ", res[2])
	}
	if res[3] != 2 {
		te.Error("res[3] expected 2, but ", res[3])
	}
	if res[4] != 0 {
		te.Error("res[4] expected 0, but ", res[4])
	}
	if res[5] != 0 {
		te.Error("res[5] expected 0, but ", res[5])
	}
	if res[6] != 0 {
		te.Error("res[6] expected 0, but ", res[6])
	}
	if res[7] != 0 {
		te.Error("res[7] expected 0, but ", res[7])
	}
	if res[8] != 0 {
		te.Error("res[8] expected 0, but ", res[8])
	}
}

func TestMysteryShoppers4(te *testing.T) {
	m := 1
	n := []int{14}
	t := []int{9}

	res := MysteryShoppers(n, t, m)

	if res[0] != 0 {
		te.Error("res[0] expected 0, but ", res[0])
	}
}
