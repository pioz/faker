# faker TODO

- Think to remove data sub package and move `var DB` and `var builders` in the same file (init.go)
- Wiki add new Builder

## Functions

Latitude() float64
LatitudeInRange(min, max float64) (float64, error)
Longitude() float64
LongitudeInRange(min, max float64) (float64, error)

### Code
Faker::Code.npi #=> "0000126252"
Faker::Code.isbn #=> "759021701-8"
Faker::Code.ean #=> "4600051000057"
Faker::Code.rut #=> "91389184-8"
Faker::Code.nric #=> "S5589083H"
Faker::Code.nric(min_age: 27, max_age: 34) #=> S8505970Z
Faker::Code.imei #= "546327785982623"
Faker::Code.asin #=> "B00000IGGJ"
Faker::Code.sin #=> "159160274"
Faker::Code.ssn

### PhoneNumber
Faker::PhoneNumber.phone_number #=> "397.693.1309 x4321"
Faker::PhoneNumber.cell_phone #=> "(186)285-7925"
Faker::PhoneNumber.cell_phone_in_e164 #=> "+944937040625"
Faker::PhoneNumber.area_code #=> "201"
Faker::PhoneNumber.exchange_code #=> "208"
Faker::PhoneNumber.subscriber_number #=> "3873"
Faker::PhoneNumber.subscriber_number(length: 2) #=> "39"
Faker::PhoneNumber.extension #=> "3764"
Faker::PhoneNumber.country_code #=> "+20"
Faker::PhoneNumber.phone_number_with_country_code #=> "+95 1-672-173-8153"
Faker::PhoneNumber.cell_phone_with_country_code #=> "+974 (190) 987-9034"

### Internet
Faker::Internet.email #=> "eliza@mann.net"
Faker::Internet.email(name: 'Nancy') #=> "nancy@terry.biz"
Faker::Internet.email(name: 'Janelle Santiago', separators: '+') #=> "janelle+santiago@becker.com"
Faker::Internet.email(domain: 'example') #=> "alice@example.name"
Faker::Internet.free_email #=> "freddy@gmail.com"
Faker::Internet.free_email(name: 'Nancy') #=> "nancy@yahoo.com"
Faker::Internet.safe_email #=> "christelle@example.org"
Faker::Internet.safe_email(name: 'Nancy') #=> "nancy@example.net"
Faker::Internet.username #=> "alexie"
Faker::Internet.username(specifier: 'Nancy') #=> "nancy"
Faker::Internet.username(specifier: 'Nancy Johnson', separators: %w(. _ -)) #=> "johnson-nancy"
Faker::Internet.username(specifier: 5..8)
Faker::Internet.username(specifier: 8)
Faker::Internet.password #=> "Vg5mSvY1UeRg7"
Faker::Internet.password(min_length: 8) #=> "YfGjIk0hGzDqS0"
Faker::Internet.password(min_length: 10, max_length: 20) #=> "EoC9ShWd1hWq4vBgFw"
Faker::Internet.password(min_length: 10, max_length: 20, mix_case: true) #=> "3k5qS15aNmG"
Faker::Internet.password(min_length: 10, max_length: 20, mix_case: true, special_characters: true) #=> "*%NkOnJsH4"
Faker::Internet.domain_name #=> "effertz.info"
Faker::Internet.domain_name(domain: "example") #=> "example.net"
Faker::Internet.domain_name(subdomain: true, domain: "example") #=> "horse.example.org"
Faker::Internet.domain_word #=> "haleyziemann"
Faker::Internet.domain_suffix #=> "info"
Faker::Internet.ip_v4_address #=> "24.29.18.175"
Faker::Internet.private_ip_v4_address #=> "10.0.0.1"
Faker::Internet.public_ip_v4_address #=> "24.29.18.175"
Faker::Internet.ip_v4_cidr #=> "24.29.18.175/21"
Faker::Internet.ip_v6_address #=> "ac5f:d696:3807:1d72:2eb5:4e81:7d2b:e1df"
Faker::Internet.ip_v6_cidr #=> "ac5f:d696:3807:1d72:2eb5:4e81:7d2b:e1df/78"
Faker::Internet.mac_address #=> "e6:0d:00:11:ed:4f"
Faker::Internet.mac_address(prefix: '55:44:33') #=> "55:44:33:02:1d:9b"
Faker::Internet.url #=> "http://thiel.com/chauncey_simonis"
Faker::Internet.url(host: 'example.com') #=> "http://example.com/clotilde.swift"
Faker::Internet.url(host: 'example.com', path: '/foobar.html') #=> "http://example.com/foobar.html"
Faker::Internet.slug #=> "pariatur_laudantium"
Faker::Internet.slug(words: 'foo bar') #=> "foo.bar"
Faker::Internet.slug(words: 'foo bar', glue: '-') #=> "foo-bar"
Faker::Internet.user_agent #=> "Mozilla/5.0 (compatible; MSIE 9.0; AOL 9.7; AOLBuild 4343.19; Windows NT 6.1; WOW64; Trident/5.0; FunWebProducts)"
Faker::Internet.user_agent(vendor: :firefox) #=> "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0"
Faker::Internet.uuid #=> "929ef6ef-b11f-38c9-111b-accd67a258b2"

URL() string
ImageURL(width int, height int) string
DomainName() string
DomainSuffix() string
IPv4Address() string
IPv6Address() string
StatusCode() string
SimpleStatusCode() int
LogLevel(logType string) string
HTTPMethod() string
UserAgent() string
ChromeUserAgent() string
FirefoxUserAgent() string
OperaUserAgent() string
SafariUserAgent() string

### Gender
Faker::Gender.type #=> "Non-binary"
Faker::Gender.binary_type #=> "Female"
Faker::Gender.short_binary_type #=> "f"

### Name ✅
Faker::Name.name             #=> "Tyshawn Johns Sr."
Faker::Name.name_with_middle #=> "Aditya Elton Douglas"
Faker::Name.first_name       #=> "Kaci"
Faker::Name.middle_name      #=> "Abraham"
Faker::Name.last_name        #=> "Ernser"
Faker::Name.male_first_name   #=> "Edward"
Faker::Name.female_first_name #=> "Natasha"
Faker::Name.prefix           #=> "Mr."
Faker::Name.suffix           #=> "IV"
Faker::Name.initials            #=> "NJM"
Faker::Name.initials(number: 2) #=> "NM"

### Words
Noun() string
Verb() string
Adverb() string
Preposition() string
Adjective() string
Word() string
Sentence(wordCount int) string
Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string
LoremIpsumWord() string
LoremIpsumSentence(wordCount int) string
LoremIpsumParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string
Question() string
Quote() string
Phrase() string

### Misc
Map() map[string]interface{}



STOP HERE FOR FIRST RELEASE


### Colors
Color() string
HexColor() string
RGBColor() []int
SafeColor() string

### Payments
Price(min, max float64) float64
CreditCard() *CreditCardInfo
CreditCardCvv() string
CreditCardExp() string
CreditCardNumber(*CreditCardOptions) string
CreditCardType() string
Currency() *CurrencyInfo
CurrencyLong() string
CurrencyShort() string
AchRouting() string
AchAccount() string
BitcoinAddress() string
BitcoinPrivateKey() string

### Company
BS() string
BuzzWord() string
Company() string
CompanySuffix() string
Job() *JobInfo
JobDescriptor() string
JobLevel() string
JobTitle() string

### App
AppName() string
AppVersion() string
AppAuthor() string

### Emoji
Emoji() string // 🤣
EmojiDescription() string // winking face
EmojiCategory() string // Smileys & Emotion
EmojiAlias() string // smiley
EmojiTag() string // happy

### Products

### Avatar (https://robohash.org)
Avatar()
AvatarWithParams(size, format, set)

### Placeholdit (https://placehold.it/300x300.png)
Placeholdit()
PlaceholditWithParams(...)

### Bank
BankAccountNumber
BankIban
BankIbanWithCountryCode
BankName
BankRoutingNumber
BankSwiftBic

### Blood
BloodGroup

### Coin
CoinFlip

### Marketing
Faker::Marketing.buzzwords #=> "rubber meets

### Measurement
Faker::Measurement.height #=> "6 inches"
Faker::Measurement.height(amount: 1.4) #=> "1.4 inches"
Faker::Measurement.height(amount: "none") #=> "inch"
Faker::Measurement.height(amount: "all") #=> "inches"
Faker::Measurement.length #=> "1 yard"
Faker::Measurement.volume #=> "10 cups"
Faker::Measurement.weight #=> "3 pounds"
Faker::Measurement.metric_height #=> "2 meters"
Faker::Measurement.metric_length #=> "0 decimeters"
Faker::Measurement.metric_volume #=> "1 liter"
Faker::Measurement.metric_weight #=> "8 grams"

### Job
Faker::Job.title #=> "Lead Accounting Associate"
Faker::Job.field #=> "Manufacturing"
Faker::Job.seniority #=> "Lead"
Faker::Job.position #=> "Supervisor"
Faker::Job.key_skill #=> "Teamwork"
Faker::Job.employment_type #=> "Full-time"
Faker::Job.education_level #=> "Bachelor"

### Funny Name

### File
Extension()
MimeType()

### Educator
Faker::Educator.university #=> "Mallowtown Technical College"
Faker::Educator.secondary_school #=> "Iceborough Secondary College"
Faker::Educator.degree #=> "Associate Degree in Criminology"
Faker::Educator.course_name #=> "Criminology 101"
Faker::Educator.subject #=> "Criminology"
Faker::Educator.campus #=> "Vertapple Campus"

### Device
Faker::Device.build_number #=> "5"
Faker::Device.manufacturer #=> "Apple"
Faker::Device.model_name #=> "iPhone 4"
Faker::Device.platform #=> "webOS"
Faker::Device.serial #=> "ejfjnRNInxh0363JC2WM"
Faker::Device.version #=> "4"

### Crypto
Faker::Crypto.md5 #=> "6b5ed240042e8a65c55ddb826c3408e6"
Faker::Crypto.sha1 #=> "4e99e31c51eef8b2d290e709f757f92e558a503f"
Faker::Crypto.sha256 #=> "51e4dbb424cd9db1ec5fb989514f2a35652ececef33f21c8dd1fd61bb8e3929d"

