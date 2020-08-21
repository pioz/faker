package faker

var langData = PoolGroup{
	"name": {"Abkhazian", "Afar", "Afrikaans", "Akan", "Albanian", "Amharic", "Arabic", "Aragonese", "Armenian", "Assamese", "Avaric", "Avestan", "Aymara", "Azerbaijani", "Bambara", "Bashkir", "Basque", "Belarusian", "Bengali", "Bislama", "Bosnian", "Breton", "Bulgarian", "Burmese", "Catalan", "Central Khmer", "Chamorro", "Chechen", "Chinese", "Church Slavic", "Chuvash", "Cornish", "Corsican", "Cree", "Croatian", "Czech", "Danish", "Dhivehi", "Dutch", "Dzongkha", "English", "Esperanto", "Estonian", "Ewe", "Faroese", "Fijian", "Finnish", "French", "Fulah", "Galician", "Ganda", "Georgian", "German", "Guarani", "Gujarati", "Haitian", "Hausa", "Hebrew", "Herero", "Hindi", "Hiri Motu", "Hungarian", "Icelandic", "Ido", "Igbo", "Indonesian", "Interlingua (International Auxiliary Language Association)", "Interlingue", "Inuktitut", "Inupiaq", "Irish", "Italian", "Japanese", "Javanese", "Kalaallisut", "Kannada", "Kanuri", "Kashmiri", "Kazakh", "Kikuyu", "Kinyarwanda", "Kirghiz", "Komi", "Kongo", "Korean", "Kuanyama", "Kurdish", "Lao", "Latin", "Latvian", "Limburgan", "Lingala", "Lithuanian", "Luba-Katanga", "Luxembourgish", "Macedonian", "Malagasy", "Malay", "Malayalam", "Maltese", "Manx", "Maori", "Marathi", "Marshallese", "Modern Greek (1453-)", "Mongolian", "Nauru", "Navajo", "Ndonga", "Nepali", "North Ndebele", "Northern Sami", "Norwegian", "Norwegian Bokmål", "Norwegian Nynorsk", "Nyanja", "Occitan (post 1500)", "Ojibwa", "Oriya", "Oromo", "Ossetian", "Pali", "Panjabi", "Persian", "Polish", "Portuguese", "Pushto", "Quechua", "Romanian", "Romansh", "Rundi", "Russian", "Samoan", "Sango", "Sanskrit", "Sardinian", "Scottish Gaelic", "Serbian", "Serbo-Croatian", "Shona", "Sichuan Yi", "Sindhi", "Sinhala", "Slovak", "Slovenian", "Somali", "South Ndebele", "Southern Sotho", "Spanish", "Sundanese", "Swahili", "Swati", "Swedish", "Tagalog", "Tahitian", "Tajik", "Tamil", "Tatar", "Telugu", "Thai", "Tibetan", "Tigrinya", "Tonga (Tonga Islands)", "Tsonga", "Tswana", "Turkish", "Turkmen", "Twi", "Uighur", "Ukrainian", "Urdu", "Uzbek", "Venda", "Vietnamese", "Volapük", "Walloon", "Welsh", "Western Frisian", "Wolof", "Xhosa", "Yiddish", "Yoruba", "Zhuang", "Zulu"},
	"code": {"aa", "ab", "ae", "af", "ak", "am", "an", "ar", "as", "av", "ay", "az", "ba", "be", "bg", "bi", "bm", "bn", "bo", "br", "bs", "ca", "ce", "ch", "co", "cr", "cs", "cu", "cv", "cy", "da", "de", "dv", "dz", "ee", "el", "en", "eo", "es", "et", "eu", "fa", "ff", "fi", "fj", "fo", "fr", "fy", "ga", "gd", "gl", "gn", "gu", "gv", "ha", "he", "hi", "ho", "hr", "ht", "hu", "hy", "hz", "ia", "id", "ie", "ig", "ii", "ik", "io", "is", "it", "iu", "ja", "jv", "ka", "kg", "ki", "kj", "kk", "kl", "km", "kn", "ko", "kr", "ks", "ku", "kv", "kw", "ky", "la", "lb", "lg", "li", "ln", "lo", "lt", "lu", "lv", "mg", "mh", "mi", "mk", "ml", "mn", "mr", "ms", "mt", "my", "na", "nb", "nd", "ne", "ng", "nl", "nn", "no", "nr", "nv", "ny", "oc", "oj", "om", "or", "os", "pa", "pi", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "rw", "sa", "sc", "sd", "se", "sg", "sh", "si", "sk", "sl", "sm", "sn", "so", "sq", "sr", "ss", "st", "su", "sv", "sw", "ta", "te", "tg", "th", "ti", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ty", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "wo", "xh", "yi", "yo", "za", "zh", "zu"},
}

// LangName will build a random language name string.
func LangName() string {
	value, _ := GetData("lang", "name")
	return value.(string)
}

// LangCode will build a random language code string.
func LangCode() string {
	value, _ := GetData("lang", "code")
	return value.(string)
}

// Builder functions

func langNameBuilder(params ...string) (interface{}, error) {
	return LangName(), nil
}

func langCodeBuilder(params ...string) (interface{}, error) {
	return LangCode(), nil
}
