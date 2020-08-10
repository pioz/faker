package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleSentece() {
	faker.SetSeed(1600)
	fmt.Println(faker.Sentence())
	// Output: A knowledgeable pomegranate without peaches is truly a panda of compassionate sharks.
}

func ExampleParagraphWithSentenceCount() {
	faker.SetSeed(1601)
	fmt.Println(faker.ParagraphWithSentenceCount(5))
	// Output: Their squirrel was, in this moment, an affable goat. A camel is a lion's panda! In modern times dazzling lemons show us how plums can be horses. This is not to discredit the idea that a cat can hardly be considered an amused fish without also being a lobster? An exuberant goldfish is a wolf of the mind.
}

func ExampleParagraph() {
	faker.SetSeed(1602)
	fmt.Println(faker.Paragraph())
	// Output: A currant is a deer from the right perspective. Having been a gymnast, authors often misinterpret the tangerine as an intelligent camel, when in actuality it feels more like a communicative bear; As far as he is concerned, a shark is the zebra of a dog. Those kumquats are nothing more than elephants. It's an undeniable fact, really; some posit the hilarious prune to be less than intellectual. A communicative zebra's goat comes with it the thought that the honest snake is a cow? Washing and polishing the car,excited cranberries show us how grapefruits can be elephants. One cannot separate bears from dynamic cows! One cannot separate cats from stimulating spiders. Extending this logic, a goldfish can hardly be considered a succinct spider without also being a cheetah.
}

func ExampleArticleWithParagraphCount() {
	faker.SetSeed(1603)
	fmt.Println(faker.ArticleWithParagraphCount(5))
	// Output: The fly of a pig becomes an alert blueberry. The vivacious apricot reveals itself as a powerful lime to those who look! Few can name a resolute turtle that isn't a happy kiwi; The pandas could be said to resemble fair-minded goats. Few can name a modest wolf that isn't a sensible fly! Some generous alligators are thought of simply as chimpanzees. Few can name a thrifty pineapple that isn't a pleasant apricot.
	// By the waybefore chickens, snakes were only lemons. If this was somewhat unclear, a kitten is a cow from the right perspective? Shouting with happiness, rabbits are romantic fishes. A watermelon is a humorous apricot. The willing deer comes from a diplomatic turtle? We can assume that any instance of a strawberry can be construed as an affable fish. Framed in a different way, the dolphins could be said to resemble receptive cherries. An intuitive fly's seal comes with it the thought that the relieved cranberry is a sheep. A blackberry is a panda's cheetah.
	// A turtle sees a panda as an efficient octopus. Seals are compassionate figs; Shouting with happiness, one cannot separate cherries from talented monkeys! Some skillful giraffes are thought of simply as elephants! Few can name an impartial cranberry that isn't a painstaking grapes. A cheetah can hardly be considered a resolute shark without also being a kitten. We know that they were lost without the helpful cherry that composed their giraffe. An energetic goldfish's grapes comes with it the thought that the eminent lemon is a blueberry! As far as he is concerned, their kumquat was, in this moment, a happy eagle. A careful grape's plum comes with it the thought that the industrious turtle is a duck.
	// Their strawberry was, in this moment, a delightful seal. A lively lion is a monkey of the mind. An ant of the plum is assumed to be a kind-hearted frog. It's an undeniable fact, really; they were lost without the brave fish that composed their snail. A mature tiger's zebra comes with it the thought that the resourceful rat is a sheep.
	// In modern times a hamster sees a peach as a helpful fig. In modern times a dog is the fox of a bee. One cannot separate chimpanzees from diplomatic snakes. Some funny seals are thought of simply as grapes; Having been a gymnast, they were lost without the encouraging nectarine that composed their pineapple. This is not to discredit the idea that punctual lions show us how zebras can be pears. Skillful chickens show us how chimpanzees can be lobsters! Few can name an inventive turtle that isn't an adaptable raspberry. Washing and polishing the car,we can assume that any instance of a lime can be construed as an intellectual currant. A cat can hardly be considered a peaceful scorpion without also being a crocodile;
}

func ExampleArticle() {
	faker.SetSeed(1604)
	fmt.Println(faker.Article())
	// Output: Few can name a sincere chimpanzee that isn't a selective cranberry. They were lost without the straightforward camel that composed their puppy; However, the zebras could be said to resemble tidy bears. Cultured hippopotamus show us how blackberries can be squirrels. In recent years, a cherry sees a nectarine as a resolute bear. This is not to discredit the idea that a crocodile is an emotional prune;
	// An amiable watermelon is a horse of the mind. To be more specific, a panda can hardly be considered an inventive hamster without also being a lime? A camel of the kangaroo is assumed to be a cooperative tiger; Before plums, kangaroos were only melons.
	// In modern times those puppies are nothing more than persimmons. The literature would have us believe that a silly octopus is not but a rabbit. Their blackberry was, in this moment, a protective banana. A puppy is the pear of a grapes. Some posit the amiable pig to be less than helpful.
	// Far from the truth, authors often misinterpret the owl as a calm crocodile, when in actuality it feels more like a placid orange. Extending this logic, some honest turtles are thought of simply as horses. The cherry of an apple becomes a brave chimpanzee; Draped neatly on a hanger, authors often misinterpret the octopus as a happy cranberry, when in actuality it feels more like a harmonious kiwi. The oranges could be said to resemble unusual lemons! Their melon was, in this moment, an impartial hippopotamus. Plums are diligent frogs. By the waya zebra is a thrifty fish.
	// The giraffes could be said to resemble broad-minded kangaroos. The first jolly plum is, in its own way, a grape! Extending this logic, the lobster is an orange. Shouting with happiness, the affable grape comes from a thrifty kiwi. However, their panda was, in this moment, an upbeat spider. An inventive deer's kumquat comes with it the thought that the excited cranberry is a spider.
	// Those bears are nothing more than goldfishes. Authors often misinterpret the bird as an imaginative persimmon, when in actuality it feels more like an intellectual apricot! An apple is an owl from the right perspective. A camel is the fig of an octopus. What we don't know for sure is whether or not foxes are faithful lions. The literature would have us believe that a pleasant tiger is not but a shark.
}

func TestSentenceBuild(t *testing.T) {
	faker.SetSeed(1720)
	s := &struct {
		Field1 string `faker:"Sentence"`
		Field2 string `faker:"ParagraphWithSentenceCount(3)"`
		Field3 string `faker:"Paragraph"`
		Field4 string `faker:"ArticleWithParagraphCount(2)"`
		Field5 string `faker:"Article"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "As far as he is concerned, the faithful zebra comes from a peaceful raspberry.", s.Field1)
	assert.Equal(t, "The successful cranberry comes from a neat eagle. To be more specific, authors often misinterpret the grape as an obedient crocodile, when in actuality it feels more like a plucky blueberry; In ancient times few can name a practical sheep that isn't a hilarious fish.", s.Field2)
	assert.Equal(t, "In ancient times the fishes could be said to resemble amused foxes. The first detailed kitten is, in its own way, a cranberry. What we don't know for sure is whether or not the raspberry is an ant. The zeitgeist contends that a chimpanzee is the grapes of a bear. They were lost without the romantic melon that composed their fly! A proud nectarine without birds is truly a dog of polite giraffes; Recent controversy aside, the alluring goldfish comes from a knowledgeable spider.", s.Field3)
	assert.Equal(t, "A straightforward currant's lime comes with it the thought that the placid cow is a cheetah. The literature would have us believe that a painstaking kangaroo is not but a cherry? Of course, they were lost without the obedient fox that composed their alligator. A cat of the shark is assumed to be a humorous pig.\nDraped neatly on a hanger, the happy lobster reveals itself as an optimistic hippopotamus to those who look. However, the glorious alligator reveals itself as a lively monkey to those who look? An alligator is an ant from the right perspective. Unfortunately, that is wrong; on the contrary, a pro-active squirrel is a fish of the mind;", s.Field4)
	assert.Equal(t, "As far as he is concerned, the wolf of a crocodile becomes a talented goldfish. An adaptable owl's hippopotamus comes with it the thought that the cooperative bird is a seal. The first amusing kitten is, in its own way, a monkey? Some assert that the lion is a cow. The ants could be said to resemble comfortable watermelons. This is not to discredit the idea that the zebras could be said to resemble creative turtles. The first eminent elephant is, in its own way, a cheetah. The peaches could be said to resemble responsible rabbits; Far from the truth, few can name a decorous kumquat that isn't a sincere cat!\nOf course, some intelligent blackberries are thought of simply as ants. To be more specific, skillful monkeys show us how goldfishes can be grapes. We can assume that any instance of a nectarine can be construed as an agreeable crocodile. The hamster of an elephant becomes a joyous lime. Draped neatly on a hanger, some posit the amused deer to be less than reflective. Though we assume the latter, hamsters are industrious apricots? Before fishes, goldfishes were only scorpions; Recent controversy aside, a camel can hardly be considered a stimulating frog without also being a pineapple. Some posit the harmonious wolf to be less than sincere.\nThe literature would have us believe that an eager pig is not but a rabbit; A kitten can hardly be considered an imaginative sheep without also being a kiwi. Recent controversy aside, the first energetic camel is, in its own way, a blackberry. Authors often misinterpret the grape as an agreeable snake, when in actuality it feels more like a stimulating blackberry. A grape is an intellectual grape. Of course, those nectarines are nothing more than lemons! They were lost without the generous rabbit that composed their strawberry. This could be, or perhaps the raspberry of a pineapple becomes a splendid raspberry. Of course, amiable birds show us how lobsters can be deers. The literature would have us believe that an amicable blackberry is not but an apple?\nThe eagles could be said to resemble intelligent currants. Fabulous crocodiles show us how kiwis can be bees. A raspberry can hardly be considered a loving lemon without also being a grapes? One cannot separate plums from resolute blackberries. We can assume that any instance of a turtle can be construed as an amicable seal. The orange of a puppy becomes a wise sheep. What we don't know for sure is whether or not some posit the sensitive chimpanzee to be less than proud? The scorpion is a rat. Those snakes are nothing more than snails? A lemon is the kitten of a currant!\nSome posit the boundless monkey to be less than fine. A succinct tiger's zebra comes with it the thought that the enchanting hamster is an elephant. A bee is an elated monkey? The powerful alligator reveals itself as a righteous kitten to those who look. The first gentle lion is, in its own way, a blackberry. Recent controversy aside, watermelons are tough cows? The instinctive scorpion reveals itself as a broad-minded deer to those who look;\nA kumquat is the strawberry of a blueberry. The literature would have us believe that an affable apple is not but a pomegranate. The literature would have us believe that an optimistic pomegranate is not but a pig. A persimmon is a credible currant; A panda sees an orange as a quick-witted ant; Nowhere is it disputed that before ants, rats were only frogs; This is not to discredit the idea that the first modern kangaroo is, in its own way, a dolphin? Shouting with happiness, a relieved horse's lobster comes with it the thought that the fabulous cherry is a fig. This is not to discredit the idea that an alluring cherry without snakes is truly a scorpion of dynamic grapefruits! The literature would have us believe that a faithful bird is not but a giraffe. Draped neatly on a hanger, dazzling melons show us how goldfishes can be ducks;\nThe hamsters could be said to resemble encouraging deers. The frog of a squirrel becomes an amiable apricot. Some upbeat giraffes are thought of simply as bananas. If this was somewhat unclear, an exclusive frog's strawberry comes with it the thought that the fair-minded horse is a hippopotamus!", s.Field5)
}
