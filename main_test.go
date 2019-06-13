package main

import "testing"
import "github.com/stretchr/testify/assert"

// func TestZero(t *testing.T) {
// 	var input = [10]string{
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
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 0, totalScore, "Shoud be 0");
// }

// func TestNormal1(t *testing.T) {
// 	var input = [10]string{
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 90, totalScore, "Shoud be 90");
// }

// func TestNormal1_Spare(t *testing.T) {
// 	var input = [10]string{
// 		"5/",  //15
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 96, totalScore, "Shoud be 96");
// }

// func TestNormal1_Strike(t *testing.T) {
// 	var input = [10]string{
// 		"X-",  // 19
// 		"54",  // 9
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 		"54",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 100, totalScore, "Shoud be 100");
// }

// func TestNormal1_Strike2(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"54",  // 9
// 		"5/",  //20
// 		"X-",  // 10
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 39, totalScore, "Shoud be 39");
// }

// func TestNormal1_EndGame1(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"X--",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 10, totalScore, "Shoud be 10");
// }

// func TestNormal1_EndGame2(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"XX-",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 20, totalScore, "Shoud be 20");
// }

// func TestNormal1_EndGame3(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"XXX",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 30, totalScore, "Shoud be 30");
// }

// func TestNormal1_EndGame4(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"X53",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 18, totalScore, "Shoud be 18");
// }

// func TestNormal1_EndGame5(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"X55",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 20, totalScore, "Shoud be 20");
// }

// func TestNormal1_EndGame6(t *testing.T) {
// 	var input = [10]string{
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"--",
// 		"5/9",
// 	};

// 	var result = playGame(input);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 19, totalScore, "Shoud be 19");
// }

// func TestPlayGame(t *testing.T) {
// 	var input1 = [10]string{
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"X-",
// 		"XXX",
// 	};

// 	var result = playGame(input1);
// 	var totalScore = sumScore(result);

// 	assert.Equal(t, 300, totalScore, "Shoud be 300");

// }

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

//========================
func TestFrame1_1(t *testing.T) {
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
		"--",
	}

	var result = playGame(input1)
	var totalScore = sumScore(result)

	assert.Equal(t, 0, totalScore, "Shoud be 0")

}

func TestFrame1_2(t *testing.T) {
	var input1 = [10]string{
		"4-",
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

	assert.Equal(t, 4, totalScore, "Shoud be 4")

}
func TestFrame1_3(t *testing.T) {
	var input1 = [10]string{
		"-5",
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

	assert.Equal(t, 5, totalScore, "Shoud be 5")

}

func TestFrame1_4(t *testing.T) {
	var input1 = [10]string{
		"45",
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

	assert.Equal(t, 9, totalScore, "Shoud be 9")

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
		"X-",
		"X-",
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