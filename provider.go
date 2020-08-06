package faker

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var providers = map[string]providerFunc{
	// misc
	providerKey("Bool", "bool"):          boolProvider,
	providerKey("", "bool"):              boolProvider,
	providerKey("PhoneNumber", "string"): phoneNumberProvider,
	providerKey("Uuid", "string"):        uuidProvider,
	// number
	providerKey("IntInRange", "int"):         intInRangeProvider,
	providerKey("Int", "int"):                intProvider,
	providerKey("", "int"):                   intProvider,
	providerKey("Int64InRange", "int64"):     int64InRangeProvider,
	providerKey("Int64", "int64"):            int64Provider,
	providerKey("", "int64"):                 int64Provider,
	providerKey("Int32InRange", "int32"):     int32InRangeProvider,
	providerKey("Int32", "int32"):            int32Provider,
	providerKey("", "int32"):                 int32Provider,
	providerKey("Int16InRange", "int16"):     int16InRangeProvider,
	providerKey("Int16", "int16"):            int16Provider,
	providerKey("", "int16"):                 int16Provider,
	providerKey("Int8InRange", "int8"):       int8InRangeProvider,
	providerKey("Int8", "int8"):              int8Provider,
	providerKey("", "int8"):                  int8Provider,
	providerKey("UintInRange", "uint"):       uintInRangeProvider,
	providerKey("Uint", "uint"):              uintProvider,
	providerKey("", "uint"):                  uintProvider,
	providerKey("Uint64InRange", "uint64"):   uint64InRangeProvider,
	providerKey("Uint64", "uint64"):          uint64Provider,
	providerKey("", "uint64"):                uint64Provider,
	providerKey("Uint32InRange", "uint32"):   uint32InRangeProvider,
	providerKey("Uint32", "uint32"):          uint32Provider,
	providerKey("", "uint32"):                uint32Provider,
	providerKey("Uint16InRange", "uint16"):   uint16InRangeProvider,
	providerKey("Uint16", "uint16"):          uint16Provider,
	providerKey("", "uint16"):                uint16Provider,
	providerKey("Uint8InRange", "uint8"):     uint8InRangeProvider,
	providerKey("Uint8", "uint8"):            uint8Provider,
	providerKey("", "uint8"):                 uint8Provider,
	providerKey("Float64InRange", "float64"): float64InRangeProvider,
	providerKey("Float64", "float64"):        float64Provider,
	providerKey("", "float64"):               float64Provider,
	providerKey("Float32InRange", "float32"): float32InRangeProvider,
	providerKey("Float32", "float32"):        float32Provider,
	providerKey("", "float32"):               float32Provider,
	// string
	providerKey("StringWithSize", "string"):  stringWithSizeProvider,
	providerKey("String", "string"):          stringProvider,
	providerKey("", "string"):                stringProvider,
	providerKey("DigitsWithSize", "string"):  digitsWithSizeProvider,
	providerKey("Digits", "string"):          digitsProvider,
	providerKey("LettersWithSize", "string"): lettersWithSizeProvider,
	providerKey("Letters", "string"):         lettersProvider,
	providerKey("Lexify", "string"):          lexifyProvider,
	providerKey("Numerify", "string"):        numerifyProvider,
	providerKey("Parameterize", "string"):    parameterizeProvider,
	providerKey("Pick", "string"):            pickProvider,
	// time
	providerKey("DurationInRange", "time.Duration"): durationInRangeProvider,
	providerKey("Duration", "time.Duration"):        durationProvider,
	providerKey("", "time.Duration"):                durationProvider,
	providerKey("Time", "time.Time"):                timeProvider,
	providerKey("", "time.Time"):                    timeProvider,
	providerKey("NanoSecond", "int"):                nanoSecondProvider,
	providerKey("Second", "int"):                    secondProvider,
	providerKey("Minute", "int"):                    minuteProvider,
	providerKey("Hour", "int"):                      hourProvider,
	providerKey("Day", "int"):                       dayProvider,
	providerKey("WeekDay", "string"):                weekDayProvider,
	providerKey("Month", "string"):                  monthProvider,
	providerKey("Year", "int"):                      yearProvider,
	providerKey("TimeZone", "string"):               timeZoneProvider,
	providerKey("TimeZoneAbbr", "string"):           timeZoneAbbrProvider,
	providerKey("TimeZoneFull", "string"):           timeZoneFullProvider,
	providerKey("TimeZoneOffset", "float32"):        timeZoneOffsetProvider,
	providerKey("TimeZoneRegion", "string"):         timeZoneRegionProvider,
	// country
	providerKey("CountryName", "string"):        countryNameProvider,
	providerKey("CountryAlpha2", "string"):      countryAlpha2Provider,
	providerKey("CountryAlpha3", "string"):      countryAlpha3Provider,
	providerKey("CountryNationality", "string"): countryNationalityProvider,
	providerKey("CountryFlag", "string"):        countryFlagProvider,
	// lang
	providerKey("LangName", "string"): langNameProvider,
	providerKey("LangCode", "string"): langCodeProvider,
	// currency
	providerKey("CurrencyName", "string"):   currencyNameProvider,
	providerKey("CurrencyCode", "string"):   currencyCodeProvider,
	providerKey("CurrencySymbol", "string"): currencySymbolProvider,
	// address
	providerKey("AddressCity", "string"):             addressCityProvider,
	providerKey("AddressState", "string"):            addressStateProvider,
	providerKey("AddressStateCode", "string"):        addressStateCodeProvider,
	providerKey("AddressStreetName", "string"):       addressStreetNameProvider,
	providerKey("AddressStreetNumber", "string"):     addressStreetNumberProvider,
	providerKey("AddressSecondaryAddress", "string"): addressSecondaryAddressProvider,
	providerKey("AddressZip", "string"):              addressZipProvider,
	providerKey("AddressFull", "string"):             addressFullProvider,
	// name
	providerKey("MaleFirstName", "string"):    maleFirstNameProvider,
	providerKey("FemaleFirstName", "string"):  femaleFirstName,
	providerKey("NeutralFirstName", "string"): neutralFirstNameProvider,
	providerKey("FirstName", "string"):        firstNameProvider,
	providerKey("LastName", "string"):         lastNameProvider,
	providerKey("NamePrefix", "string"):       namePrefixProvider,
	providerKey("NameSuffix", "string"):       nameSuffixProvider,
	providerKey("FullName", "string"):         fullNameProvider,
	providerKey("NameInitials", "string"):     nameInitialsProvider,
	// gender
	providerKey("Gender", "string"):            genderProvider,
	providerKey("BinaryGender", "string"):      binaryGenderProvider,
	providerKey("ShortBinaryGender", "string"): shortBinaryGenderProvider,
	// sentence
	providerKey("Sentence", "string"):                   sentenceProvider,
	providerKey("ParagraphWithSentenceCount", "string"): paragraphWithSentenceCountProvider,
	providerKey("Paragraph", "string"):                  paragraphProvider,
	providerKey("ArticleWithParagraphCount", "string"):  articleWithParagraphCountProvider,
	providerKey("Article", "string"):                    articleProvider,
	// internet
	providerKey("Username", "string"):  usernameProvider,
	providerKey("Domain", "string"):    domainProvider,
	providerKey("Email", "string"):     emailProvider,
	providerKey("FreeEmail", "string"): freeEmailProvider,
	providerKey("SafeEmail", "string"): safeEmailProvider,
	providerKey("Slug", "string"):      slugProvider,
	providerKey("Url", "string"):       urlProvider,
}

const skipTag = "-"
const uniqueTag = "unique"

var (
	tagFnCallRegexp = regexp.MustCompile(`(.+?)\((.+?)\)`)
	tagLenRegexp    = regexp.MustCompile(`len=(\d+)`)
)

type fakerTag struct {
	funcName string
	unique   bool
	length   int
	params   []string
}

func (tag *fakerTag) mustSkip() bool {
	return tag.funcName == skipTag
}

func decodeTag(tagString string) *fakerTag {
	tag := &fakerTag{}
	for _, token := range strings.Split(tagString, ";") {
		if token == skipTag {
			tag.funcName = skipTag
			return tag
		}
		if token == uniqueTag {
			tag.unique = true
			continue
		}
		if m := tagLenRegexp.FindStringSubmatch(token); len(m) == 2 {
			tag.length, _ = strconv.Atoi(m[1])
			continue
		}
		if tag.funcName == "" {
			if m := tagFnCallRegexp.FindStringSubmatch(token); len(m) == 3 {
				tag.funcName = m[1]
				tag.params = strings.Split(m[2], ",")
				continue
			}
			tag.funcName = token
		}
	}
	return tag
}

type providerFunc func(...string) (interface{}, error)

func providerKey(providerName, providerType string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", providerName, providerType))
}

func RegisterProvider(providerName, providerType string, fn providerFunc) error {
	key := providerKey(providerName, providerType)
	if _, ok := providers[key]; ok {
		return errors.New("Provider already registered")
	}
	providers[key] = fn
	return nil
}

func UnregisterProvider(providerName, providerType string) error {
	key := providerKey(providerName, providerType)
	if _, ok := providers[key]; !ok {
		return errors.New("Provider not registered")
	}
	delete(providers, key)
	return nil
}
