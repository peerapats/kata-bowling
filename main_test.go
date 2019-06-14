package main 

import "testing" 
// import "github.com/stretchr/testify/assert" 

type Person struct { 
	name string 
	} 
			func TestAlwaysReturn0(t *testing.T) { 
				var input = []string{"--", "--", "--", "--", "--", "--", "--" , "--" , "--" , "--"} 
				result,_ := getScore(input) 
				if result != 0 { 
				t.Error("Result should be 0") 
				} 
				} 

func TestAlwaysReturnOpen(t *testing.T) { 
var input = []string{"11", "11", "11", "11", "11", "11", "11" , "11" , "11" , "11"} 
result,_ := getScore(input) 
if result != 20 { 
t.Error("Result should be 20") 
} 
} 

func TestAlwaysReturnSpare(t *testing.T) { 
var input = []string{"--", "--", "1/", "--", "--", "--", "--", "--", "--", "--"} 
result,_ := getScore(input) 
if result !=10 { 
t.Error("Result should be 10") 
} 
} 

func TestAlwaysReturnDoubleSpare(t *testing.T) { 
	var input = []string{"--", "--", "1/", "1/", "--", "--", "--", "--", "--", "--"} 
	result,_ := getScore(input) 
	if result !=21 { 
	t.Error("Result should be 21") 
	} 
	} 

func TestAlwaysReturnStrike(t *testing.T) { 
	var input = []string{"--", "--", "--", "--", "x", "--","--","--","--","--"} 
	result,_ := getScore(input) 
	if result != 10 { 
	t.Error("Result should be 10") 
	} 
	} 

func TestAlwaysReturnDoubleStrike(t *testing.T) { 
var input = []string{"--", "--", "--", "--", "x", "x","--","--","--","--"} 
result,_ := getScore(input) 
if result != 30 { 
t.Error("Result should be 30") 
} 
} 


	

func TestAlwaysReturnAllStrike(t *testing.T) { 
var input = []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "xxx"}
result,_ := getScore(input) 
if result != 300 { 
t.Error("Result should be 300") 
} 
} 

func TestAlwaysReturnLastFrameStrike(t *testing.T) { 
	var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "x12"}
	result,_ := getScore(input) 
	if result != 13 { 
	t.Error("Result should be 13") 
	} 
	} 

	func TestAlwaysReturnSpareLastFrameSpare(t *testing.T) { 
		var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "1/2"}
		result,_ := getScore(input) 
		if result != 12 { 
		t.Error("Result should be 12") 
		} 
	} 

		func TestAlwaysReturnFrame9Strike(t *testing.T) { 
			var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "x", "x11"}
			result,_ := getScore(input) 
			if result != 33 { 
			t.Error("Result should be 33") 
			} 
			} 

			func TestAlwaysReturnFrame9Spare(t *testing.T) { 
				var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "1/", "22"}
				result,_ := getScore(input) 
				if result != 16  { 
				t.Error("Result should be 16") 
				} 
				} 
			func TestAlwaysReturnFrame9Open(t *testing.T) { 
				var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "11", "--"}
				result,_ := getScore(input) 
				if result != 2  { 
				t.Error("Result should be 2") 
				} 
				} 

// func TestAlwaysReturn43WithAssert(t *testing.T) { 
// var input = []string{"22", "44", "1/", "12", "x", "52"} 
// result,_ := getScore(input) 
// assert.Equal(t, 43, result, "Result should be 43") 
// }