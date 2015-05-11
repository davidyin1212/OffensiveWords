package main

/*
For my solution I made a couple of assumptions first that the file names and path are fixed in their structure
because of this I hardcoded it in the main function and didn't use comand line args however I can do so if that's the
case. Also I assumend that the input files can be changed so I tried to keep my code as flexiable as possible and 
I tried modularizing my code by seperating the hash creation and getting the score of a file so that any programer
can generate a different hash based upon a different set of files and get different scores for the other inputs.
I implemented my solution by building a hash table from the phrases both the complete phrases which I issue a weight
to depending on the risk and also partials for phrases with multiple words. I use fragments so that its easier to
determine that I should continue looking ahead for more words. Then I tokenize all the words and strip punctuations
then by doing a lookup in the hashtable I can determine the weight of the word or determine the phrase weight. Then
by summing all together I can get the result for that file. Overall runtime should be O(n) where n is the sum of the 
length of high risk and low risk files and then another O(n) for lookup of each file where n is the number of words
of file. Though my solution is relativly flexible it doesn't handle subsets so well so if there is a case where one 
of the phrases is a subset of the another phrase it doesn't really handle that so something like "hello world" and "hello"
right now it double counts but I'm not sure if that's the intended logic 
*/

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"bytes"
)

//Check for error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Function to initialize hash with the words from the low risk and high risk files
//it also handle phrases by breaking them down into fragments and then hashing each individual one
func hash_init(high_risk_phrases string, low_risk_phrases string) (m map[string]int) {
	var risk_map map[string]int
	risk_map = make(map[string]int)

	file, err := os.Open(high_risk_phrases)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		high_risk_phrase := strings.ToLower(scanner.Text())
		risk_map[high_risk_phrase] = 2

		high_risk_words := strings.Split(high_risk_phrase, " ")
		if len(high_risk_words) > 1 {
			for s := range high_risk_words {
				risk_map[high_risk_words[s]] = 3
			}
		}
	}

	file, err = os.Open(low_risk_phrases)
	check(err)
	defer file.Close()

	scanner = bufio.NewScanner(file)
	check(scanner.Err())
	for scanner.Scan() {
		low_risk_phrase := strings.ToLower(scanner.Text())
		risk_map[low_risk_phrase] = 1

		low_risk_words := strings.Split(low_risk_phrase, " ")
		if len(low_risk_words) > 1 {
			for s := range low_risk_words {
				risk_map[low_risk_words[s]] = 3
			}
		}
	}
	return risk_map
}

//Provide the hash map and the path to the file and it will score the file
//based upon the hashed values from the high and low risk files
func get_file_score(risk_map map[string]int, file_path string) (score int) {
	file, err := os.Open(file_path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(scanner.Err())
	scanner.Split(bufio.ScanWords)
	var result int
	for scanner.Scan() {
		val := risk_map[strings.ToLower(stripchars(scanner.Text()))]
		if val == 3 {
			val = 0
			var buffer bytes.Buffer
			buffer.WriteString(strings.ToLower(scanner.Text()))
			for scanner.Scan() {
				buffer.WriteString(" ")
				buffer.WriteString(strings.ToLower(scanner.Text()))
				val = risk_map[stripchars(buffer.String())]
				if val == 1 || val == 2 {
					break
				}
				if risk_map[strings.ToLower(stripchars(scanner.Text()))] != 3 {
					val = risk_map[strings.ToLower(stripchars(scanner.Text()))]
					break
				}
			}
		}
		result += val
	}
	return result
}

//removes punctuations so that tokens only contain characters
func stripchars(str string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune("?!\",.@#", r) < 0 {
            return r
        }
        return -1
    }, str)
}

//main function executes all the operations
func main() {
	risk_map := hash_init("./offensive_text/high_risk_phrases.txt", "./offensive_text/low_risk_phrases.txt")

	file, err := os.Create("output.txt")
	check(err)
	bufferWriter := bufio.NewWriter(file)
	_, err = bufferWriter.WriteString("input01:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input01.txt")) + "\n")
	_, err = bufferWriter.WriteString("input02:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input02.txt")) + "\n")
	_, err = bufferWriter.WriteString("input03:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input03.txt")) + "\n")
	_, err = bufferWriter.WriteString("input04:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input04.txt")) + "\n")
	_, err = bufferWriter.WriteString("input05:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input05.txt")) + "\n")
	_, err = bufferWriter.WriteString("input06:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input06.txt")) + "\n")
	_, err = bufferWriter.WriteString("input07:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input07.txt")) + "\n")
	_, err = bufferWriter.WriteString("input08:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input08.txt")) + "\n")
	_, err = bufferWriter.WriteString("input09:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input09.txt")) + "\n")
	_, err = bufferWriter.WriteString("input10:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input10.txt")) + "\n")
	_, err = bufferWriter.WriteString("input11:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input11.txt")) + "\n")
	_, err = bufferWriter.WriteString("input12:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input12.txt")) + "\n")
	_, err = bufferWriter.WriteString("input13:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input13.txt")) + "\n")
	_, err = bufferWriter.WriteString("input14:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input14.txt")) + "\n")
	_, err = bufferWriter.WriteString("input15:" + strconv.Itoa(get_file_score(risk_map, "./offensive_text/input15.txt")) + "\n")
	bufferWriter.Flush()

}