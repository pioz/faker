package faker

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

func Sentence() string {
	phrase := randomStartingPhrase()
	phrase += makeSentence()
	phrase = capitalize(phrase)
	phrase += pickLastPunc()
	return phrase
}

func ParagraphWithSentenceCount(n int) string {
	p := make([]string, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, Sentence())
	}
	return strings.Join(p, " ")
}

func Paragraph() string {
	n := IntInRange(4, 11)
	return ParagraphWithSentenceCount(n)
}

func ArticleWithParagraphCount(n int) string {
	p := make([]string, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, Paragraph())
	}
	return strings.Join(p, "\n")
}

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
	sentenceTemplate, err := GetData("sentence", "sentence_template")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"noun": func() string {
			v, err := GetData("sentence", "noun")
			if err != nil {
				panic(err)
			}
			return v.(string)
		},
		"anoun": func() string {
			v, err := GetData("sentence", "noun")
			if err != nil {
				panic(err)
			}
			return articleize(v.(string))
		},
		"nouns": func() string {
			v, err := GetData("sentence", "noun")
			if err != nil {
				panic(err)
			}
			return pluralize(v.(string))
		},
		"adjective": func() string {
			v, err := GetData("sentence", "adjective")
			if err != nil {
				panic(err)
			}
			return v.(string)
		},
		"anAdjective": func() string {
			v, err := GetData("sentence", "adjective")
			if err != nil {
				panic(err)
			}
			return articleize(v.(string))
		},
	}

	buf := bytes.NewBufferString("")
	tmpl := template.Must(template.New("sentence").Funcs(funcMap).Parse(sentenceTemplate.(string)))
	err = tmpl.Execute(buf, "")
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
