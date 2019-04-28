package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/russelltsherman/new-age-bullshit/vocab"
)

var sentencePool [][]string
var sentencesGenerated int

// Sentence function
func Sentence() string {
	// if sentence pool is empty refill it
	if len(sentencePool) == 0 {
		initializeSentencePool()
	}

	topic := randInt(len(sentencePool) - 1)
	patternNumber := randInt(len(sentencePool[topic]) - 1)
	pattern := sentencePool[topic][patternNumber]
	removePatternFromPool(topic, patternNumber)
	pattern = processPattern(pattern)
	sentencesGenerated++
	// log.Printf("%v sentences generated", sentencesGenerated)

	return pattern
}

// Text generates a number of sentences
func Text(sentences int) []string {
	body := []string{}
	for i := 0; i < sentences; i++ {
		body = append(body, Sentence())
	}
	return body
}

// initialize the sentence pool
func initializeSentencePool() {
	sentencePool = vocab.NewSentencePool()
}

// process replacement patterns
func processPattern(pattern string) string {
	// insert a space before . , ; ?
	var re = regexp.MustCompile(`([\.,;\?])`)
	pattern = re.ReplaceAllString(pattern, ` $1`)

	// split string into array
	words := strings.Fields(pattern)

	var result bytes.Buffer
	for _, word := range words {
		// replace word if matches one of the placeholder words (e.g. nPerson),
		result.WriteString(vocab.Replace(word))
		result.WriteString(" ")
	}
	output := result.String()

	// replace 'a [vowel]' with 'an [vowel]'
	re = regexp.MustCompile(`(^|\W)([Aa]) ([aeiou])`)
	output = re.ReplaceAllString(output, `$1$2n $3`)

	// remove spaces before . , ; ?
	re = regexp.MustCompile(` ([,\.;\?])`)
	output = re.ReplaceAllString(output, `$1`)

	// take care of prefixes (delete the space after the hyphen)
	re = regexp.MustCompile(`- `)
	output = re.ReplaceAllString(output, `$1`)

	// add space after question marks if they're mid-sentence
	re = regexp.MustCompile(`\?(\w)`)
	output = re.ReplaceAllString(output, `? $1`)

	// return sentence with first letter capitalized
	return ucFirst(output)
}

// return a random int up to max
func randInt(max int) int {
	if max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// remove a sentence parttern from the pool to prevent reuse
func removePatternFromPool(topic int, idx int) {
	sentencePool[topic] = append(sentencePool[topic][:idx], sentencePool[topic][idx+1:]...)
	if len(sentencePool[topic]) == 0 {
		removeTopicFromPool(topic)
	}
}

// remove a topic from pool when empty
func removeTopicFromPool(topic int) {
	sentencePool = append(sentencePool[:topic], sentencePool[topic+1:]...)
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
