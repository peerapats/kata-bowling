
package main 

import ( 
"fmt" 
"strconv" 
"unicode/utf8" 
"strings" 
) 

func getScore(input []string) (int, []int) { 
var score []int 
var total int 
for i, element := range input { 
var sum int 

if(i == 9 ) { 
var a int; 
var b int; 
var c int; 
a=0
b=0
c=0
res := strings.Split(element, "") 
if(res[0] == "x"){
	a= 10
}else{
	a,_ = strconv.Atoi(res[0]); 
}

if(res[1] == "x"){
	b= 10
}else if(res[1] == "/"){
	b = 10 - a
}else{
	b,_ = strconv.Atoi(res[1]); 
}
if(len(element) == 3){
	if(res[2] == "x"){
	  c = 10
	}else{
	  c,_ = strconv.Atoi(res[2]); 
	}
}
var sum = a+b+c; 
score = append(score, sum); 
} else { 

if(element == "x") { 

sum = 10; 
score = append(score, 10); 
// fmt.Println("input",input[i+1]) 
// ให้เอา 2 ตาต่อไปมาบวก 
//บวกรวบแรก 
if(input[i+1] == "x") {
score[i] += 10; 
var isSet = false; 
//เหลืออีก 2 รอบ 
if(i < len(input) - 2) && (input[i+2][0] == 'x'){ 
score[i] += 10; 
} else { 
for _, c := range input[i+2] { 
buf := make([]byte, 1) 
_ = utf8.EncodeRune(buf, c) 
value, _ := strconv.Atoi(string(buf)) 
if(!isSet) { 
score[i] += value; 
isSet = true; 
} 
} 
} 
} else { 
for _, c := range input[i+1] { 
buf := make([]byte, 1) 
_ = utf8.EncodeRune(buf, c) 
value, _ := strconv.Atoi(string(buf)) 
score[i] += value; 
} 
} 

if(i == 8){
	res := strings.Split(input[i+1], "") 
	var a int; 
    a=0
	if(input[i] == "x") {
		score[i] = 10;
		if(res[0] == "x"){
			score[i] += 10;	
		} else {
			a,_ = strconv.Atoi(res[0]); 
			score[i] += a
		}
		if(res[1] == "x"){
			score[i] += 10;	
		}else if(res[1] == "/"){
			a,_ = strconv.Atoi(res[0]); 
			score[i] += 10 - a;	
		}else{
			a,_ = strconv.Atoi(res[1]); 
			score[i] += a;	
		}
	}
}


} else { 
var isSp = false; 
for _, c := range element { 
buf := make([]byte, 1) 
_ = utf8.EncodeRune(buf, c) 
value, _ := strconv.Atoi(string(buf)) 

if(c == '/') { 
isSp = true; 
sum = 10; 
} else { 
sum += value; 
} 

} 

score = append(score, sum); 

if(isSp && i <= 9) { 
if(i < len(input) - 1) && (input[i+1] == "X"){ 
score[i] += 10; 
} else { 
var isSet = false; 
for _, c := range input[i+1] { 
buf := make([]byte, 1) 
_ = utf8.EncodeRune(buf, c) 
value, _ := strconv.Atoi(string(buf)) 
if(!isSet) { 
score[i] += value; 
isSet = true; 
} 
} 
} 
} 
} 

} 

} 

// fmt.Println("Input",input); 
for _, v := range score { 
total += v 
//fmt.Println("Frame ", i + 1 ,"[" , input[i], "] : ", v); 
} 
//fmt.Println("total ", total); 
return total, score; 
} 

func main () { 
	// var input = []string{"--", "--", "1/", "--", "--", "--", "--", "--", "--", "--"} 
	var input = []string{"--", "--", "--", "--", "--", "--", "--", "--", "x", "x11"}

	 total,score :=getScore(input)
	 fmt.Println(input)
	 for i, v := range score {  
		fmt.Println("Frame ", i + 1 ,"[" , input[i], "] : ", v); 
		} 
		fmt.Println("Total is ",total)
}