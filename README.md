# OffensiveWords
run using: go run solution.go

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