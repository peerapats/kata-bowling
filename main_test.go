package main

import "testing"
import "github.com/stretchr/testify/assert"

// ล้างท่อ
func TestZero(t *testing.T) {
	var input = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	};

	var result = playGame(input);
	var totalScore = sumScore(result);

	assert.Equal(t, 0, totalScore, "Shoud be 0");
}

// แต้มธรรมดา
func TestNormal1(t *testing.T) {
	var input = [10]string{
		"54",
		"54",
		"54",
		"54",
		"54",
		"54",
		"54",
		"54",
		"54",
		"54",
	};

	var result = playGame(input);
	var totalScore = sumScore(result);

	assert.Equal(t, 90, totalScore, "Shoud be 90");
}

func TestFrame1_5(t *testing.T) {
	var input1 = [10]string{
		"4/",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}

func TestFrame1_6(t *testing.T) {
	var input1 = [10]string{
		"X-",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}

func TestFrame2_1(t *testing.T) {
	var input1 = [10]string{
		"X-", //20
		"X-", //10
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 30, totalScore, "Shoud be 30")

}

func TestFrame2_2(t *testing.T) {
	var input1 = [10]string{
		"X-",
		"X-",
		"-/",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 50, totalScore, "Shoud be 50")

}

func TestFrame2_3(t *testing.T) {
	var input1 = [10]string{
		"X-",
		"4/",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 30, totalScore, "Shoud be 30")

}

// F10 ST
func TestFrame10_1(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"X--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}

// F10 SP
func TestFrame10_2(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"-/-",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}

// F10 Num
func TestFrame10_3(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"34",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 7, totalScore, "Shoud be 7")

}

// F10 2ST
func TestFrame10_4(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"XX-",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 20, totalScore, "Shoud be 20")

}

// F10 3ST
func TestFrame10_5(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"XXX",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 30, totalScore, "Shoud be 30")

}

// F9 ST
func TestFrame9_1(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"X-",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}

// F9 SP
func TestFrame9_2(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"4/",
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 10, totalScore, "Shoud be 10")

}
// F9 ST-ST
func TestFrame9_4(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"X-",
		"X--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 30, totalScore, "Shoud be 30")

}

// F9 ST-SP
func TestFrame9_3(t *testing.T) {
	var input1 = [10]string{
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"--",
		"X-",
		"5/-",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 30, totalScore, "Shoud be 30")

}
// //========================
// func TestFrame1_1(t *testing.T) {
// 	var input1 = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 	}

// 	var result = playGame(input1)
// 	var totalScore = sumScore(result)

// 	assert.Equal(t, 0, totalScore, "Shoud be 0")

// }

// func TestFrame1_2(t *testing.T) {
// 	var input1 = [10]string{
// 		"4-",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 	}

// 	var result = playGame(input1)
// 	var totalScore = sumScore(result)

// 	assert.Equal(t, 4, totalScore, "Shoud be 4")

// // }
// func TestFrame1_3(t *testing.T) {
// 	var input1 = [10]string{
// 		"-5",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 	}

// 	var result = playGame(input1)
// 	var totalScore = sumScore(result)

// 	assert.Equal(t, 5, totalScore, "Shoud be 5")

// }

// func TestFrame1_4(t *testing.T) {
// 	var input1 = [10]string{
// 		"45",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 	}

// 	var result = playGame(input1)
// 	var totalScore = sumScore(result)

// 	assert.Equal(t, 9, totalScore, "Shoud be 9")

// }

