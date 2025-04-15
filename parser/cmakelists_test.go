package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCMakelists(t *testing.T) {

	str := `
	cmake_minimum_required(VERSION 2.8...3.15)
project(RoaringBitmap
  DESCRIPTION "Roaring bitmaps in C (and C++)"
  LANGUAGES CXX C
)
set (CMAKE_C_STANDARD 11)#hello
set (CMAKE_CXX_STANDARD 11)

include(GNUInstallDirs)
#这是注释`

	commands, err := ParseCMakelists(str)
	assert.Nil(t, err)
	assert.Equal(t, len(commands), 5)
	fmt.Printf("%s\n", commands[0])

	commands, err = ParseCMakelists("")
	assert.Nil(t, err)
	assert.Equal(t, len(commands), 0)

	commands, err = ParseCMakelists("#")
	assert.Nil(t, err)
	assert.Equal(t, len(commands), 0)

	commands, err = ParseCMakelists("####")
	assert.Nil(t, err)
	assert.Equal(t, len(commands), 0)

	commands, err = ParseCMakelists(" # abcd\n")
	assert.Nil(t, err)
	assert.Equal(t, len(commands), 0)

}
