package faker

var currencyData = PoolGroup{
	"name":   {"Afghan Afghani", "Albanian Lek", "Algerian Dinar", "Angolan Kwanza", "Argentine Peso", "Armenian Dram", "Aruban Florin", "Australian Dollar", "Azerbaijani Manat", "Bahamian Dollar", "Bahraini Dinar", "Bangladeshi Taka", "Barbadian Dollar", "Belarusian Ruble", "Belize Dollar", "Bermudian Dollar", "Bhutanese Ngultrum", "Bitcoin", "Bitcoin Cash", "Bolivian Boliviano", "Bosnia and Herzegovina Convertible Mark", "Botswana Pula", "Brazilian Real", "British Penny", "British Pound", "Brunei Dollar", "Bulgarian Lev", "Burundian Franc", "Cambodian Riel", "Canadian Dollar", "Cape Verdean Escudo", "Cayman Islands Dollar", "Central African Cfa Franc", "Cfp Franc", "Chilean Peso", "Chinese Renminbi Yuan", "Chinese Renminbi Yuan Offshore", "Codes specifically reserved for testing purposes", "Colombian Peso", "Comorian Franc", "Congolese Franc", "Costa Rican Colón", "Croatian Kuna", "Cuban Convertible Peso", "Cuban Peso", "Czech Koruna", "Danish Krone", "Djiboutian Franc", "Dominican Peso", "East Caribbean Dollar", "Egyptian Pound", "Eritrean Nakfa", "Estonian Kroon", "Ethiopian Birr", "Euro", "European Composite Unit", "European Monetary Unit", "European Unit of Account 17", "European Unit of Account 9", "Falkland Pound", "Fijian Dollar", "Gambian Dalasi", "Georgian Lari", "Ghanaian Cedi", "Gibraltar Pound", "Gold (Troy Ounce)", "Guatemalan Quetzal", "Guernsey Pound", "Guinean Franc", "Guyanese Dollar", "Haitian Gourde", "Honduran Lempira", "Hong Kong Dollar", "Hungarian Forint", "Icelandic Króna", "Indian Rupee", "Indonesian Rupiah", "Iranian Rial", "Iraqi Dinar", "Isle of Man Pound", "Israeli New Sheqel", "Jamaican Dollar", "Japanese Yen", "Jersey Pound", "Jordanian Dinar", "Kazakhstani Tenge", "Kenyan Shilling", "Kuwaiti Dinar", "Kyrgyzstani Som", "Lao Kip", "Latvian Lats", "Lebanese Pound", "Lesotho Loti", "Liberian Dollar", "Libyan Dinar", "Lithuanian Litas", "Macanese Pataca", "Macedonian Denar", "Malagasy Ariary", "Malawian Kwacha", "Malaysian Ringgit", "Maldivian Rufiyaa", "Maltese Lira", "Mauritanian Ouguiya", "Mauritian Rupee", "Mexican Peso", "Moldovan Leu", "Mongolian Tögrög", "Moroccan Dirham", "Mozambican Metical", "Myanmar Kyat", "Namibian Dollar", "Nepalese Rupee", "Netherlands Antillean Gulden", "New Taiwan Dollar", "New Zealand Dollar", "Nicaraguan Córdoba", "Nigerian Naira", "North Korean Won", "Norwegian Krone", "Omani Rial", "Pakistani Rupee", "Palladium", "Panamanian Balboa", "Papua New Guinean Kina", "Paraguayan Guaraní", "Peruvian Sol", "Philippine Peso", "Platinum", "Polish Złoty", "Qatari Riyal", "Romanian Leu", "Russian Ruble", "Rwandan Franc", "Saint Helenian Pound", "Salvadoran Colón", "Samoan Tala", "Saudi Riyal", "Serbian Dinar", "Seychellois Rupee", "Sierra Leonean Leone", "Silver (Troy Ounce)", "Singapore Dollar", "Slovak Koruna", "Solomon Islands Dollar", "Somali Shilling", "South African Rand", "South Korean Won", "South Sudanese Pound", "Special Drawing Rights", "Sri Lankan Rupee", "Sudanese Pound", "Surinamese Dollar", "Swazi Lilangeni", "Swedish Krona", "Swiss Franc", "Syrian Pound", "São Tomé and Príncipe Dobra", "Tajikistani Somoni", "Tanzanian Shilling", "Thai Baht", "Tongan Paʻanga", "Trinidad and Tobago Dollar", "Tunisian Dinar", "Turkish Lira", "Turkmenistani Manat", "UIC Franc", "Ugandan Shilling", "Ukrainian Hryvnia", "Unidad de Fomento", "United Arab Emirates Dirham", "United States Dollar", "Uruguayan Peso", "Uzbekistan Som", "Vanuatu Vatu", "Venezuelan Bolívar", "Venezuelan Bolívar Soberano", "Vietnamese Đồng", "West African Cfa Franc", "Yemeni Rial", "Zambian Kwacha", "Zimbabwean Dollar"},
	"code":   {"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BCH", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTC", "BTN", "BWP", "BYN", "BYR", "BZD", "CAD", "CDF", "CHF", "CLF", "CLP", "CNH", "CNY", "COP", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EEK", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GBX", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR", "ISK", "JEP", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LVL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SKK", "SLL", "SOS", "SRD", "SSP", "STD", "SVC", "SYP", "SZL", "THB", "TJS", "TMM", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VEF", "VES", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XFU", "XOF", "XPD", "XPF", "XPT", "XTS", "YER", "ZAR", "ZMK", "ZMW", "ZWD", "ZWL", "ZWN", "ZWR"},
	"symbol": {"$", "Ar", "B/.", "Br", "Bs", "Bs.", "Bs.F", "C$", "CHF", "D", "Db", "E", "FRw", "Fdj", "Fr", "Ft", "G", "K", "KR", "KSh", "Kz", "Kč", "L", "Le", "Lei", "Ls", "Lt", "MK", "MTn", "MVR", "Nfk", "Nu.", "P", "Q", "R", "R$", "RM", "Rp", "S/", "SDR", "Sh", "Sk", "T", "T$", "UF", "UM", "USh", "Vt", "ZK", "kn", "kr", "kr.", "m", "oz t", "so'm", "som", "zł", "£", "£S", "¥", "ƒ", "ЅМ", "КМ", "РСД", "ден", "лв.", "դր.", "؋", "ب.د", "ج.م", "د.إ", "د.ا", "د.ت", "د.ج", "د.ك", "د.م.", "ر.س", "ر.ع.", "ر.ق", "ع.د", "ل.د", "ل.ل", "৳", "฿", "ლ", "៛", "₡", "₤", "₦", "₨", "₩", "₪", "₫", "€", "₭", "₮", "₱", "₲", "₴", "₵", "₸", "₹", "₺", "₼", "₽", "₿", "﷼"},
}

// CurrencyName will build a random currency name string.
func CurrencyName() string {
	value, _ := GetData("currency", "name")
	return value.(string)
}

// CurrencyCode will build a random currency code string.
func CurrencyCode() string {
	value, _ := GetData("currency", "code")
	return value.(string)
}

// CurrencySymbol will build a random currency symbol string.
func CurrencySymbol() string {
	value, _ := GetData("currency", "symbol")
	return value.(string)
}

// Builder functions

func currencyNameBuilder(params ...string) (interface{}, error) {
	return CurrencyName(), nil
}

func currencyCodeBuilder(params ...string) (interface{}, error) {
	return CurrencyCode(), nil
}

func currencySymbolBuilder(params ...string) (interface{}, error) {
	return CurrencySymbol(), nil
}
