package stupidadder_test

import (
	"log"
	stupidadder "stupid_adder"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStupidAdder(t *testing.T) {
	testname := "Fuck you"

	log.Println(testname)
	fuckingResult := stupidadder.Add(1, 2)
	require.Equal(t, int64(3), fuckingResult)

}
