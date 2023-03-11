package ast

import "testing"

func TestRollRouterBack(t *testing.T) {
	RollRouterBack("ttt", "Testttt")
}

func TestRollGormBack(t *testing.T) {
	RollGormBack("ttt", "Testttt")
}
