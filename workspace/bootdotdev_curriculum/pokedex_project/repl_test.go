package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
	"github.com/osamaNazieh/pokedexcli/internal/util"
)



func TestCleanInput(t *testing.T) {
	cases := map[string]struct{
		input string 
		expected []string
	} {
		"regular":{input: " hello world ", expected: []string{"hello", "world"}},
		"no trimming":{input: "hello world", expected: []string {"hello", "world"}},
		"no sep":{input: "helloworld", expected: []string {"helloworld"}},
		"not a whitespace sep":{input: "hello_world",expected: []string {"hello_world"}},
		"no value":{input: "",expected: []string{"pokedex > "}},
		"all space":{input: "    ",expected: []string {"pokedex > "}},
	}
	for name, tc  := range cases {
		t.Run(name, func (t *testing.T)  {
			got := util.CleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)		
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
	
}

	
