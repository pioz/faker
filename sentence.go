package faker

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

var sentenceData = PoolGroup{
	"noun":              {"alligator", "ant", "apple", "apricot", "banana", "bear", "bee", "bird", "blackberry", "blueberry", "camel", "cat", "cheetah", "cherry", "chicken", "chimpanzee", "cow", "cranberry", "crocodile", "currant", "deer", "dog", "dolphin", "duck", "eagle", "elephant", "fig", "fish", "fly", "fox", "frog", "giraffe", "goat", "goldfish", "grape", "grapefruit", "grapes", "hamster", "hippopotamus", "horse", "kangaroo", "kitten", "kiwi", "kumquat", "lemon", "lime", "lion", "lobster", "melon", "monkey", "nectarine", "octopus", "orange", "owl", "panda", "peach", "pear", "persimmon", "pig", "pineapple", "plum", "pomegranate", "prune", "puppy", "rabbit", "raspberry", "rat", "scorpion", "seal", "shark", "sheep", "snail", "snake", "spider", "squirrel", "strawberry", "tangerine", "tiger", "turtle", "watermelon", "wolf", "zebra"},
	"adjective":         {"adaptable", "adventurous", "affable", "affectionate", "agreeable", "alert", "alluring", "ambitious", "amiable", "amicable", "amused", "amusing", "boundless", "brave", "bright", "broad-minded", "calm", "capable", "careful", "charming", "cheerful", "coherent", "comfortable", "communicative", "compassionate", "confident", "conscientious", "considerate", "convivial", "cooperative", "courageous", "courteous", "creative", "credible", "cultured", "dashing", "dazzling", "debonair", "decisive", "decorous", "delightful", "detailed", "determined", "diligent", "diplomatic", "discreet", "dynamic", "eager", "easygoing", "efficient", "elated", "eminent", "emotional", "enchanting", "encouraging", "endurable", "energetic", "entertaining", "enthusiastic", "excellent", "excited", "exclusive", "exuberant", "fabulous", "fair", "fair-minded", "faithful", "fantastic", "fearless", "fine", "forceful", "frank", "friendly", "funny", "generous", "gentle", "glorious", "good", "gregarious", "happy", "hard-working", "harmonious", "helpful", "hilarious", "honest", "honorable", "humorous", "imaginative", "impartial", "independent", "industrious", "instinctive", "intellectual", "intelligent", "intuitive", "inventive", "jolly", "joyous", "kind", "kind-hearted", "knowledgeable", "level", "likeable", "lively", "lovely", "loving", "loyal", "lucky", "mature", "modern", "modest", "neat", "nice", "obedient", "optimistic", "painstaking", "passionate", "patient", "peaceful", "perfect", "persistent", "philosophical", "pioneering", "placid", "plausible", "pleasant", "plucky", "polite", "powerful", "practical", "pro-active", "productive", "protective", "proud", "punctual", "quick-witted", "quiet", "rational", "receptive", "reflective", "reliable", "relieved", "reserved", "resolute", "resourceful", "responsible", "rhetorical", "righteous", "romantic", "sedate", "seemly", "selective", "self-assured", "self-confident", "self-disciplined", "sensible", "sensitive", "shrewd", "shy", "silly", "sincere", "skillful", "smiling", "sociable", "splendid", "steadfast", "stimulating", "straightforward", "successful", "succinct", "sympathetic", "talented", "thoughtful", "thrifty", "tidy", "tough", "trustworthy", "unassuming", "unbiased", "understanding", "unusual", "upbeat", "versatile", "vigorous", "vivacious", "warm", "warmhearted", "willing", "wise", "witty", "wonderful"},
	"sentence_template": {"the {{noun}} is {{anoun}}", "{{anoun}} is {{anAdjective}} {{noun}}", "the first {{adjective}} {{noun}} is, in its own way, {{anoun}}", "their {{noun}} was, in this moment, {{anAdjective}} {{noun}}", "{{anoun}} is {{anoun}} from the right perspective", "the literature would have us believe that {{anAdjective}} {{noun}} is not but {{anoun}}", "{{anAdjective}} {{noun}} is {{anoun}} of the mind", "the {{adjective}} {{noun}} reveals itself as {{anAdjective}} {{noun}} to those who look", "authors often misinterpret the {{noun}} as {{anAdjective}} {{noun}}, when in actuality it feels more like {{anAdjective}} {{noun}}", "we can assume that any instance of {{anoun}} can be construed as {{anAdjective}} {{noun}}", "they were lost without the {{adjective}} {{noun}} that composed their {{noun}}", "the {{adjective}} {{noun}} comes from {{anAdjective}} {{noun}}", "{{anoun}} can hardly be considered {{anAdjective}} {{noun}} without also being {{anoun}}", "few can name {{anAdjective}} {{noun}} that isn't {{anAdjective}} {{noun}}", "some posit the {{adjective}} {{noun}} to be less than {{adjective}}", "{{anoun}} of the {{noun}} is assumed to be {{anAdjective}} {{noun}}", "{{anoun}} sees {{anoun}} as {{anAdjective}} {{noun}}", "the {{noun}} of {{anoun}} becomes {{anAdjective}} {{noun}}", "{{anoun}} is {{anoun}}'s {{noun}}", "{{anoun}} is the {{noun}} of {{anoun}}", "{{anAdjective}} {{noun}}'s {{noun}} comes with it the thought that the {{adjective}} {{noun}} is {{anoun}}", "{{nouns}} are {{adjective}} {{nouns}}", "{{adjective}} {{nouns}} show us how {{nouns}} can be {{nouns}}", "before {{nouns}}, {{nouns}} were only {{nouns}}", "those {{nouns}} are nothing more than {{nouns}}", "some {{adjective}} {{nouns}} are thought of simply as {{nouns}}", "one cannot separate {{nouns}} from {{adjective}} {{nouns}}", "the {{nouns}} could be said to resemble {{adjective}} {{nouns}}", "{{anAdjective}} {{noun}} without {{nouns}} is truly a {{noun}} of {{adjective}} {{nouns}}"},
	"starting_phrase":   {"to be more specific, ", "in recent years, ", "however, ", "by the way", "of course, ", "some assert that ", "if this was somewhat unclear, ", "unfortunately, that is wrong; on the contrary, ", "it's very tricky, if not impossible, ", "this could be, or perhaps ", "this is not to discredit the idea that ", "we know that ", "it's an undeniable fact, really; ", "framed in a different way, ", "what we don't know for sure is whether or not ", "as far as we can estimate, ", "as far as he is concerned, ", "the zeitgeist contends that ", "though we assume the latter, ", "far from the truth, ", "extending this logic, ", "nowhere is it disputed that ", "in modern times ", "in ancient times ", "recent controversy aside, ", "washing and polishing the car,", "having been a gymnast, ", "after a long day at school and work, ", "waking to the buzz of the alarm clock, ", "draped neatly on a hanger, ", "shouting with happiness, "},
}

// Sentence will build a random sentence string.
func Sentence() string {
	phrase := randomStartingPhrase()
	phrase += makeSentence()
	phrase = capitalize(phrase)
	phrase += pickLastPunc()
	return phrase
}

// ParagraphWithSentenceCount will build a random paragraph string consisting
// of n sentences.
func ParagraphWithSentenceCount(n int) string {
	p := make([]string, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, Sentence())
	}
	return strings.Join(p, " ")
}

// Paragraph will build a random paragraph.
func Paragraph() string {
	n := IntInRange(4, 11)
	return ParagraphWithSentenceCount(n)
}

// ArticleWithParagraphCount will build a random article string consisting
// of n paragraphs.
func ArticleWithParagraphCount(n int) string {
	p := make([]string, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, Paragraph())
	}
	return strings.Join(p, "\n")
}

// Article will build a random article.
func Article() string {
	n := IntInRange(4, 11)
	return ArticleWithParagraphCount(n)
}

// Builder functions

func sentenceBuilder(params ...string) (interface{}, error) {
	return Sentence(), nil
}

func paragraphWithSentenceCountBuilder(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return ParagraphWithSentenceCount(size), nil
}

func paragraphBuilder(params ...string) (interface{}, error) {
	return Paragraph(), nil
}

func articleWithParagraphCountBuilder(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return ArticleWithParagraphCount(size), nil
}

func articleBuilder(params ...string) (interface{}, error) {
	return Article(), nil
}

// Private functions

var (
	punctuation     = []string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "!", "?", "!", "?", ";"}
	pluralizeRegexp = regexp.MustCompile(`(ss|ish|ch|x|us)$`)
)

const vowels = "aeiouy"

func makeSentence() string {
	sentenceTemplate, _ := GetData("sentence", "sentence_template")

	funcMap := template.FuncMap{
		"noun": func() string {
			value, _ := GetData("sentence", "noun")
			return value.(string)
		},
		"anoun": func() string {
			value, _ := GetData("sentence", "noun")
			return articleize(value.(string))
		},
		"nouns": func() string {
			value, _ := GetData("sentence", "noun")
			return pluralize(value.(string))
		},
		"adjective": func() string {
			value, _ := GetData("sentence", "adjective")
			return value.(string)
		},
		"anAdjective": func() string {
			value, _ := GetData("sentence", "adjective")
			return articleize(value.(string))
		},
	}

	buf := bytes.NewBufferString("")
	tmpl := template.Must(template.New("sentence").Funcs(funcMap).Parse(sentenceTemplate.(string)))
	err := tmpl.Execute(buf, "")
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func randomStartingPhrase() string {
	if Float32InRange(0, 1) < 0.33 {
		phrase, err := GetData("sentence", "starting_phrase")
		if err != nil {
			panic(err)
		}
		return phrase.(string)
	}
	return ""
}

func pickLastPunc() string {
	i := IntInRange(0, len(punctuation)-1)
	return punctuation[i]
}

func pluralize(word string) string {
	if strings.HasSuffix(word, "s") {
		return word
	}

	if pluralizeRegexp.MatchString(word) {
		word += "e"
	} else {
		if strings.HasSuffix(word, "y") && !strings.Contains(vowels, string(word[len(word)-2])) {
			word = strings.TrimSuffix(word, string(word[len(word)-1]))
			word += "ie"
		}
	}
	return word + "s"
}

func articleize(word string) string {
	a := "a"
	if strings.Contains("aeiou", string(word[0])) {
		a = "an"
	}
	return fmt.Sprintf("%s %s", a, word)
}

func capitalize(word string) string {
	chars := []rune(word)
	chars[0] = unicode.ToUpper(chars[0])
	return string(chars)
}
