package faker

import (
	"errors"
	"fmt"
	"strings"
)

var builders = map[string]builderFunc{
	// misc
	builderKey("Bool", "bool"):          boolBuilder,
	builderKey("", "bool"):              boolBuilder,
	builderKey("PhoneNumber", "string"): phoneNumberBuilder,
	builderKey("Uuid", "string"):        uuidBuilder,
	// number
	builderKey("IntInRange", "int"):         intInRangeBuilder,
	builderKey("Int", "int"):                intBuilder,
	builderKey("", "int"):                   intBuilder,
	builderKey("Int64InRange", "int64"):     int64InRangeBuilder,
	builderKey("Int64", "int64"):            int64Builder,
	builderKey("", "int64"):                 int64Builder,
	builderKey("Int32InRange", "int32"):     int32InRangeBuilder,
	builderKey("Int32", "int32"):            int32Builder,
	builderKey("", "int32"):                 int32Builder,
	builderKey("Int16InRange", "int16"):     int16InRangeBuilder,
	builderKey("Int16", "int16"):            int16Builder,
	builderKey("", "int16"):                 int16Builder,
	builderKey("Int8InRange", "int8"):       int8InRangeBuilder,
	builderKey("Int8", "int8"):              int8Builder,
	builderKey("", "int8"):                  int8Builder,
	builderKey("UintInRange", "uint"):       uintInRangeBuilder,
	builderKey("Uint", "uint"):              uintBuilder,
	builderKey("", "uint"):                  uintBuilder,
	builderKey("Uint64InRange", "uint64"):   uint64InRangeBuilder,
	builderKey("Uint64", "uint64"):          uint64Builder,
	builderKey("", "uint64"):                uint64Builder,
	builderKey("Uint32InRange", "uint32"):   uint32InRangeBuilder,
	builderKey("Uint32", "uint32"):          uint32Builder,
	builderKey("", "uint32"):                uint32Builder,
	builderKey("Uint16InRange", "uint16"):   uint16InRangeBuilder,
	builderKey("Uint16", "uint16"):          uint16Builder,
	builderKey("", "uint16"):                uint16Builder,
	builderKey("Uint8InRange", "uint8"):     uint8InRangeBuilder,
	builderKey("Uint8", "uint8"):            uint8Builder,
	builderKey("", "uint8"):                 uint8Builder,
	builderKey("Float64InRange", "float64"): float64InRangeBuilder,
	builderKey("Float64", "float64"):        float64Builder,
	builderKey("", "float64"):               float64Builder,
	builderKey("Float32InRange", "float32"): float32InRangeBuilder,
	builderKey("Float32", "float32"):        float32Builder,
	builderKey("", "float32"):               float32Builder,
	// string
	builderKey("StringWithSize", "string"):  stringWithSizeBuilder,
	builderKey("String", "string"):          stringBuilder,
	builderKey("", "string"):                stringBuilder,
	builderKey("DigitsWithSize", "string"):  digitsWithSizeBuilder,
	builderKey("Digits", "string"):          digitsBuilder,
	builderKey("LettersWithSize", "string"): lettersWithSizeBuilder,
	builderKey("Letters", "string"):         lettersBuilder,
	builderKey("Lexify", "string"):          lexifyBuilder,
	builderKey("Numerify", "string"):        numerifyBuilder,
	builderKey("Parameterize", "string"):    parameterizeBuilder,
	builderKey("Pick", "string"):            pickBuilder,
	// time
	builderKey("DurationInRange", "time.Duration"): durationInRangeBuilder,
	builderKey("Duration", "time.Duration"):        durationBuilder,
	builderKey("", "time.Duration"):                durationBuilder,
	builderKey("Time", "time.Time"):                timeBuilder,
	builderKey("", "time.Time"):                    timeBuilder,
	builderKey("NanoSecond", "int"):                nanoSecondBuilder,
	builderKey("Second", "int"):                    secondBuilder,
	builderKey("Minute", "int"):                    minuteBuilder,
	builderKey("Hour", "int"):                      hourBuilder,
	builderKey("Day", "int"):                       dayBuilder,
	builderKey("WeekDay", "string"):                weekDayBuilder,
	builderKey("Month", "string"):                  monthBuilder,
	builderKey("Year", "int"):                      yearBuilder,
	builderKey("TimeZone", "string"):               timeZoneBuilder,
	builderKey("TimeZoneAbbr", "string"):           timeZoneAbbrBuilder,
	builderKey("TimeZoneFull", "string"):           timeZoneFullBuilder,
	builderKey("TimeZoneOffset", "float32"):        timeZoneOffsetBuilder,
	builderKey("TimeZoneRegion", "string"):         timeZoneRegionBuilder,
	// country
	builderKey("CountryName", "string"):        countryNameBuilder,
	builderKey("CountryAlpha2", "string"):      countryAlpha2Builder,
	builderKey("CountryAlpha3", "string"):      countryAlpha3Builder,
	builderKey("CountryNationality", "string"): countryNationalityBuilder,
	builderKey("CountryFlag", "string"):        countryFlagBuilder,
	// lang
	builderKey("LangName", "string"): langNameBuilder,
	builderKey("LangCode", "string"): langCodeBuilder,
	// currency
	builderKey("CurrencyName", "string"):   currencyNameBuilder,
	builderKey("CurrencyCode", "string"):   currencyCodeBuilder,
	builderKey("CurrencySymbol", "string"): currencySymbolBuilder,
	// address
	builderKey("AddressCity", "string"):             addressCityBuilder,
	builderKey("AddressState", "string"):            addressStateBuilder,
	builderKey("AddressStateCode", "string"):        addressStateCodeBuilder,
	builderKey("AddressStreetName", "string"):       addressStreetNameBuilder,
	builderKey("AddressStreetNumber", "string"):     addressStreetNumberBuilder,
	builderKey("AddressSecondaryAddress", "string"): addressSecondaryAddressBuilder,
	builderKey("AddressZip", "string"):              addressZipBuilder,
	builderKey("AddressFull", "string"):             addressFullBuilder,
	// name
	builderKey("MaleFirstName", "string"):    maleFirstNameBuilder,
	builderKey("FemaleFirstName", "string"):  femaleFirstName,
	builderKey("NeutralFirstName", "string"): neutralFirstNameBuilder,
	builderKey("FirstName", "string"):        firstNameBuilder,
	builderKey("LastName", "string"):         lastNameBuilder,
	builderKey("NamePrefix", "string"):       namePrefixBuilder,
	builderKey("NameSuffix", "string"):       nameSuffixBuilder,
	builderKey("FullName", "string"):         fullNameBuilder,
	builderKey("NameInitials", "string"):     nameInitialsBuilder,
	// gender
	builderKey("Gender", "string"):            genderBuilder,
	builderKey("BinaryGender", "string"):      binaryGenderBuilder,
	builderKey("ShortBinaryGender", "string"): shortBinaryGenderBuilder,
	// sentence
	builderKey("Sentence", "string"):                   sentenceBuilder,
	builderKey("ParagraphWithSentenceCount", "string"): paragraphWithSentenceCountBuilder,
	builderKey("Paragraph", "string"):                  paragraphBuilder,
	builderKey("ArticleWithParagraphCount", "string"):  articleWithParagraphCountBuilder,
	builderKey("Article", "string"):                    articleBuilder,
	// internet
	builderKey("Username", "string"):  usernameBuilder,
	builderKey("Domain", "string"):    domainBuilder,
	builderKey("Email", "string"):     emailBuilder,
	builderKey("FreeEmail", "string"): freeEmailBuilder,
	builderKey("SafeEmail", "string"): safeEmailBuilder,
	builderKey("Slug", "string"):      slugBuilder,
	builderKey("Url", "string"):       urlBuilder,
}

type builderFunc func(...string) (interface{}, error)

func builderKey(builderName, builderType string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", builderName, builderType))
}

func RegisterBuilder(builderName, builderType string, fn builderFunc) error {
	key := builderKey(builderName, builderType)
	if _, ok := builders[key]; ok {
		return errors.New("Builder already registered")
	}
	builders[key] = fn
	return nil
}

func UnregisterBuilder(builderName, builderType string) error {
	key := builderKey(builderName, builderType)
	if _, ok := builders[key]; !ok {
		return errors.New("Builder not registered")
	}
	delete(builders, key)
	return nil
}
