/*
	Author: David Newswanger

	This is my implementation of r/dailyprogrammers rhyming challenge. This program
	takes in a number file of words which have been broken down into "phenomes"
	which describe how each word is pronounced. This program uses those phenomes
	to fined words that rhyme with whatever the user enters.

	https://www.reddit.com/r/dailyprogrammer/comments/4fnz37/20160420_challenge_263_intermediate_help_eminem/
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*

  Store data in "reverse tree", where the last sounds of each word are stored at the root of the tree

*/

//Node
type Node struct {
	children map[string]*Node
	word     string
}

func main() {
	var text string
	fmt.Print("Enter a word that you would like to rhyme: ")
	fmt.Scanf("%s", &text)

	text = strings.ToUpper(text)
	fmt.Println(text)

	root, words := parseData()
	results := getRhymes(words[text], root)

	//fmt.Println(text, words[text])
	fmt.Println("-------------- Results --------------")

	for i := 1; i <= len(results); i++ {

		fmt.Println("\n", len(results[i]), "Matches:")
		for _, v := range results[i] {
			fmt.Println("[", i+1, "]", v, ": ", words[v][1:])
		}
	}

}

func parseData() (Node, map[string][]string) {
	/* reads the words file and returns a tree */

	// Loads the pronunciation data from the file.
	dat, _ := ioutil.ReadFile("data/words.txt")
	words := strings.Split(string(dat), "\n")[56:]

	// Creates the tree's root node
	root := Node{children: make(map[string]*Node)}

	wordRef := make(map[string][]string)

	for _, x := range words {
		// Each word gets stored in a tree and addressed by the phenom, with the phenom for the last part of the
		// word being addressed first. Therefore, the root node of the tree contains the last phenom for each item in
		// the tree
		current := &root
		all := strings.Split(x, " ")
		word := all[0]
		s := all[1:]
		wordRef[word] = s

		// Loop through the word's phenoms in reverse order, starting with the last phenom first
		for i := len(s) - 1; i >= 0; i-- {

			// If a node for the phenom doesn't exist, create it
			if current.children[s[i]] == nil {
				current.children[s[i]] = &Node{children: make(map[string]*Node)}
			}

			// Set the current node to the current phenom
			current = current.children[s[i]]
		}

		// Once we've reached the last phenom for the word, save the word as a leaf in the tree.
		current.word = word

	}

	return root, wordRef
}

func getRhymes(pattern []string, root Node) map[int][]string {
	results := []string{}

	m := make(map[int][]string)
	listed := make(map[string]bool)

	// This loop traveses the tree until it reaches the node represented by the pattern that is passed in.
	// When the node is reached, the words under that node are returned and added to the dictionary.
	// The loop goes through every possible list of phenomes. If a phenome is X Y Z, it will travers to
	// X Y Z, then  X Y then Y in order to load the different numbers of matching phenomes
	for i := range pattern {
		current := &root

		// Each iteration cuts off the next initial phenome
		p := pattern[i:]

		// Break if there is only 1 phenome left
		if len(p)-1 == 0 {
			break
		}

		// Traverses until the given node, the passes the node to recTrav which returns all the words under
		// the given node.
		for j := len(p) - 1; j >= 0; j-- {
			current = current.children[p[j]]
		}

		results = recTrav(current)

		if len(results) > 1 {

			// Words that have already been added are saved to the listed array so that repeats are avoided.s
			for _, v := range results {
				if listed[v] == false {
					m[len(p)-1] = append(m[len(p)-1], v)
					listed[v] = true
				}
			}
		}

	}

	return m

}

// Returns an array with all of the words underneath a given node in the tree
func recTrav(node *Node) []string {
	if node.word != "" {
		return []string{node.word}
	}

	results := []string{}
	for _, c := range node.children {
		results = append(results, recTrav(c)...)
	}

	return results
}
