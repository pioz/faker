package faker

import (
	"strconv"
	"time"
)

var timezoneData = PoolGroup{
	"offset": {"-12", "-11", "-10", "-8", "-7", "-6", "-5", "-4.5", "-4", "-3", "-2.5", "-2", "-1", "0", "-1", "0", "1", "2", "3", "4", "4.5", "5", "5.5", "5.75", "6", "6.5", "7", "8", "9", "9.5", "10", "11", "12", "13"},
	"abbr":   {"DST", "U", "HST", "AKDT", "PDT", "PDT", "PST", "UMST", "MDT", "MDT", "CAST", "CDT", "CDT", "CCST", "SPST", "EDT", "UEDT", "VST", "PYT", "ADT", "CBST", "SWST", "PSST", "NDT", "ESAST", "AST", "SEST", "GDT", "MST", "BST", "U", "MDT", "ADT", "CVST", "MDT", "UTC", "GMT", "BST", "GDT", "GST", "WEDT", "CEDT", "RDT", "CEDT", "WCAST", "NST", "GDT", "MEDT", "EST", "SDT", "EEDT", "SAST", "FDT", "TDT", "JDT", "LST", "JST", "AST", "KST", "AST", "EAST", "MSK", "SAMT", "IDT", "AST", "ADT", "MST", "GST", "CST", "AST", "WAST", "YEKT", "PKT", "IST", "SLST", "NST", "CAST", "BST", "MST", "SAST", "NCAST", "CST", "NAST", "MPST", "WAST", "TST", "UST", "NAEST", "JST", "KST", "CAST", "ACST", "EAST", "AEST", "WPST", "TST", "YST", "CPST", "VST", "NZST", "U", "FST", "MST", "KDT", "TST", "SST"},
	"text":   {"Dateline Standard Time", "UTC-11", "Hawaiian Standard Time", "Alaskan Standard Time", "Pacific Standard Time (Mexico)", "Pacific Daylight Time", "Pacific Standard Time", "US Mountain Standard Time", "Mountain Standard Time (Mexico)", "Mountain Standard Time", "Central America Standard Time", "Central Standard Time", "Central Standard Time (Mexico)", "Canada Central Standard Time", "SA Pacific Standard Time", "Eastern Standard Time", "US Eastern Standard Time", "Venezuela Standard Time", "Paraguay Standard Time", "Atlantic Standard Time", "Central Brazilian Standard Time", "SA Western Standard Time", "Pacific SA Standard Time", "Newfoundland Standard Time", "E. South America Standard Time", "Argentina Standard Time", "SA Eastern Standard Time", "Greenland Standard Time", "Montevideo Standard Time", "Bahia Standard Time", "UTC-02", "Mid-Atlantic Standard Time", "Azores Standard Time", "Cape Verde Standard Time", "Morocco Standard Time", "UTC", "Greenwich Mean Time", "British Summer Time", "GMT Standard Time", "Greenwich Standard Time", "W. Europe Standard Time", "Central Europe Standard Time", "Romance Standard Time", "Central European Standard Time", "W. Central Africa Standard Time", "Namibia Standard Time", "GTB Standard Time", "Middle East Standard Time", "Egypt Standard Time", "Syria Standard Time", "E. Europe Standard Time", "South Africa Standard Time", "FLE Standard Time", "Turkey Standard Time", "Israel Standard Time", "Libya Standard Time", "Jordan Standard Time", "Arabic Standard Time", "Kaliningrad Standard Time", "Arab Standard Time", "E. Africa Standard Time", "Moscow Standard Time", "Samara Time", "Iran Standard Time", "Arabian Standard Time", "Azerbaijan Standard Time", "Mauritius Standard Time", "Georgian Standard Time", "Caucasus Standard Time", "Afghanistan Standard Time", "West Asia Standard Time", "Yekaterinburg Time", "Pakistan Standard Time", "India Standard Time", "Sri Lanka Standard Time", "Nepal Standard Time", "Central Asia Standard Time", "Bangladesh Standard Time", "Myanmar Standard Time", "SE Asia Standard Time", "N. Central Asia Standard Time", "China Standard Time", "North Asia Standard Time", "Singapore Standard Time", "W. Australia Standard Time", "Taipei Standard Time", "Ulaanbaatar Standard Time", "North Asia East Standard Time", "Japan Standard Time", "Korea Standard Time", "Cen. Australia Standard Time", "AUS Central Standard Time", "E. Australia Standard Time", "AUS Eastern Standard Time", "West Pacific Standard Time", "Tasmania Standard Time", "Yakutsk Standard Time", "Central Pacific Standard Time", "Vladivostok Standard Time", "New Zealand Standard Time", "UTC+12", "Fiji Standard Time", "Magadan Standard Time", "Kamchatka Standard Time", "Tonga Standard Time", "Samoa Standard Time"},
	"full":   {"(UTC-12:00) International Date Line West", "(UTC-11:00) Coordinated Universal Time-11", "(UTC-10:00) Hawaii", "(UTC-09:00) Alaska", "(UTC-08:00) Baja California", "(UTC-07:00) Pacific Time (US & Canada)", "(UTC-08:00) Pacific Time (US & Canada)", "(UTC-07:00) Arizona", "(UTC-07:00) Chihuahua, La Paz, Mazatlan", "(UTC-07:00) Mountain Time (US & Canada)", "(UTC-06:00) Central America", "(UTC-06:00) Central Time (US & Canada)", "(UTC-06:00) Guadalajara, Mexico City, Monterrey", "(UTC-06:00) Saskatchewan", "(UTC-05:00) Bogota, Lima, Quito", "(UTC-05:00) Eastern Time (US & Canada)", "(UTC-05:00) Indiana (East)", "(UTC-04:30) Caracas", "(UTC-04:00) Asuncion", "(UTC-04:00) Atlantic Time (Canada)", "(UTC-04:00) Cuiaba", "(UTC-04:00) Georgetown, La Paz, Manaus, San Juan", "(UTC-04:00) Santiago", "(UTC-03:30) Newfoundland", "(UTC-03:00) Brasilia", "(UTC-03:00) Buenos Aires", "(UTC-03:00) Cayenne, Fortaleza", "(UTC-03:00) Greenland", "(UTC-03:00) Montevideo", "(UTC-03:00) Salvador", "(UTC-02:00) Coordinated Universal Time-02", "(UTC-02:00) Mid-Atlantic - Old", "(UTC-01:00) Azores", "(UTC-01:00) Cape Verde Is.", "(UTC) Casablanca", "(UTC) Coordinated Universal Time", "(UTC) Edinburgh, London", "(UTC+01:00) Edinburgh, London", "(UTC) Dublin, Lisbon", "(UTC) Monrovia, Reykjavik", "(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna", "(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague", "(UTC+01:00) Brussels, Copenhagen, Madrid, Paris", "(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb", "(UTC+01:00) West Central Africa", "(UTC+01:00) Windhoek", "(UTC+02:00) Athens, Bucharest", "(UTC+02:00) Beirut", "(UTC+02:00) Cairo", "(UTC+02:00) Damascus", "(UTC+02:00) E. Europe", "(UTC+02:00) Harare, Pretoria", "(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius", "(UTC+03:00) Istanbul", "(UTC+02:00) Jerusalem", "(UTC+02:00) Tripoli", "(UTC+03:00) Amman", "(UTC+03:00) Baghdad", "(UTC+03:00) Kaliningrad, Minsk", "(UTC+03:00) Kuwait, Riyadh", "(UTC+03:00) Nairobi", "(UTC+03:00) Moscow, St. Petersburg, Volgograd", "(UTC+04:00) Samara, Ulyanovsk, Saratov", "(UTC+03:30) Tehran", "(UTC+04:00) Abu Dhabi, Muscat", "(UTC+04:00) Baku", "(UTC+04:00) Port Louis", "(UTC+04:00) Tbilisi", "(UTC+04:00) Yerevan", "(UTC+04:30) Kabul", "(UTC+05:00) Ashgabat, Tashkent", "(UTC+05:00) Yekaterinburg", "(UTC+05:00) Islamabad, Karachi", "(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi", "(UTC+05:30) Sri Jayawardenepura", "(UTC+05:45) Kathmandu", "(UTC+06:00) Astana", "(UTC+06:00) Dhaka", "(UTC+06:30) Yangon (Rangoon)", "(UTC+07:00) Bangkok, Hanoi, Jakarta", "(UTC+07:00) Novosibirsk", "(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi", "(UTC+08:00) Krasnoyarsk", "(UTC+08:00) Kuala Lumpur, Singapore", "(UTC+08:00) Perth", "(UTC+08:00) Taipei", "(UTC+08:00) Ulaanbaatar", "(UTC+09:00) Irkutsk", "(UTC+09:00) Osaka, Sapporo, Tokyo", "(UTC+09:00) Seoul", "(UTC+09:30) Adelaide", "(UTC+09:30) Darwin", "(UTC+10:00) Brisbane", "(UTC+10:00) Canberra, Melbourne, Sydney", "(UTC+10:00) Guam, Port Moresby", "(UTC+10:00) Hobart", "(UTC+10:00) Yakutsk", "(UTC+11:00) Solomon Is., New Caledonia", "(UTC+11:00) Vladivostok", "(UTC+12:00) Auckland, Wellington", "(UTC+12:00) Coordinated Universal Time+12", "(UTC+12:00) Fiji", "(UTC+12:00) Magadan", "(UTC+12:00) Petropavlovsk-Kamchatsky - Old", "(UTC+13:00) Nuku'alofa", "(UTC+13:00) Samoa"},
	"region": {"Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers", "Africa/Asmara", "Africa/Bamako", "Africa/Bangui", "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville", "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta", "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti", "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone", "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala", "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos", "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi", "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane", "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena", "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo", "Africa/Sao_Tome", "Africa/Timbuktu", "Africa/Tripoli", "Africa/Tunis", "Africa/Windhoek", "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua", "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca", "America/Argentina/ComodRivadavia", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis", "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion", "America/Atikokan", "America/Atka", "America/Bahia", "America/Bahia_Banderas", "America/Barbados", "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota", "America/Boise", "America/Buenos_Aires", "America/Cambridge_Bay", "America/Campo_Grande", "America/Cancun", "America/Caracas", "America/Catamarca", "America/Cayenne", "America/Cayman", "America/Chicago", "America/Chihuahua", "America/Coral_Harbour", "America/Cordoba", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao", "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit", "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Ensenada", "America/Fort_Nelson", "America/Fort_Wayne", "America/Fortaleza", "America/Glace_Bay", "America/Godthab", "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala", "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo", "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg", "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac", "America/Indianapolis", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Jujuy", "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Knox_IN", "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Louisville", "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot", "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Mendoza", "America/Menominee", "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton", "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau", "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Ojinaga", "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port_of_Spain", "America/Port-au-Prince", "America/Porto_Acre", "America/Porto_Velho", "America/Puerto_Rico", "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina", "America/Resolute", "America/Rio_Branco", "America/Rosario", "America/Santa_Isabel", "America/Santarem", "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock", "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia", "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule", "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver", "America/Virgin", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife", "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok", "Arctic/Longyearbyen", "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat", "Asia/Ashkhabad", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul", "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Chita", "Asia/Choibalsan", "Asia/Chongqing", "Asia/Chungking", "Asia/Colombo", "Asia/Dacca", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai", "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong", "Asia/Hovd", "Asia/Irkutsk", "Asia/Istanbul", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul", "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macao", "Asia/Macau", "Asia/Magadan", "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk", "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qyzylorda", "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai", "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Tel_Aviv", "Asia/Thimbu", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ujung_Pandang", "Asia/Ulaanbaatar", "Asia/Ulan_Bator", "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faeroe", "Atlantic/Faroe", "Atlantic/Jan_Mayen", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena", "Atlantic/Stanley", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Canberra", "Australia/Currie", "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/Lindeman", "Australia/Lord_Howe", "Australia/Melbourne", "Australia/Perth", "Australia/Sydney", "Australia/Yancowinna", "Etc/GMT", "Etc/GMT+0", "Etc/GMT+1", "Etc/GMT+10", "Etc/GMT+11", "Etc/GMT+12", "Etc/GMT+2", "Etc/GMT+3", "Etc/GMT+4", "Etc/GMT+5", "Etc/GMT+6", "Etc/GMT+7", "Etc/GMT+8", "Etc/GMT+9", "Etc/GMT0", "Etc/GMT-0", "Etc/GMT-1", "Etc/GMT-10", "Etc/GMT-11", "Etc/GMT-12", "Etc/GMT-13", "Etc/GMT-14", "Etc/GMT-2", "Etc/GMT-3", "Etc/GMT-4", "Etc/GMT-5", "Etc/GMT-6", "Etc/GMT-7", "Etc/GMT-8", "Etc/GMT-9", "Etc/UTC", "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belfast", "Europe/Belgrade", "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen", "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki", "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta", "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Asia/Nicosia", "Europe/Oslo", "Europe/Paris", "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo", "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane", "Europe/Tiraspol", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius", "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich", "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe", "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion", "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter", "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos", "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway", "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau", "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Ponape", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan", "Pacific/Samoa", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Truk", "Pacific/Wake", "Pacific/Wallis", "Pacific/Yap"},
}

// DurationInRange will build a random time.Duration between min and max
// included.
func DurationInRange(min, max time.Duration) time.Duration {
	return time.Duration(Int64InRange(int64(min), int64(max)))
}

// Duration will build a random time.Duration.
func Duration() time.Duration {
	return time.Duration(Int64())
}

// TimeInRange will build a random time.Time between min and max included.
func TimeInRange(min, max time.Time) time.Time {
	if min.After(max) {
		return min
	}
	d := DurationInRange(0, max.Sub(min))
	return min.Add(d)
}

// Time will build a random time.Time.
func Time() time.Time {
	return time.Now().Add(Duration())
}

// TimeNow will build the current time.
func TimeNow() time.Time {
	return time.Now()
}

// NanoSecond will build a random nano second.
func NanoSecond() int {
	return IntInRange(0, 999999999)
}

// Second will build a random second.
func Second() int {
	return IntInRange(0, 59)
}

// Minute will build a random minute.
func Minute() int {
	return IntInRange(0, 59)
}

// Hour will build a random hour in military time.
func Hour() int {
	return IntInRange(0, 23)
}

// Day will build a random day between 1 - 31.
func Day() int {
	return IntInRange(1, 31)
}

// WeekDay will build a random weekday string (Monday-Sunday).
func WeekDay() string {
	return time.Weekday(IntInRange(0, 6)).String()
}

// Month will build a random month string.
func Month() string {
	return time.Month(IntInRange(1, 12)).String()
}

// Year will build a random year between 1900 - 2100.
func Year() int {
	return IntInRange(1900, 2100)
}

// TimeZone will build a random timezone string.
func TimeZone() string {
	timezone, _ := GetData("timezone", "text")
	return timezone.(string)
}

// TimeZoneAbbr will build a random abbreviated timezone string.
func TimeZoneAbbr() string {
	timezone, _ := GetData("timezone", "abbr")
	return timezone.(string)
}

// TimeZoneFull will build a random full timezone string.
func TimeZoneFull() string {
	timezone, _ := GetData("timezone", "full")
	return timezone.(string)
}

// TimeZoneOffset will build a random timezone offset.
func TimeZoneOffset() float32 {
	offsetString, _ := GetData("timezone", "offset")
	offset, err := strconv.ParseFloat(offsetString.(string), 32)
	if err != nil {
		panic(err)
	}
	return float32(offset)
}

// TimeZoneRegion will build a random timezone region string.
func TimeZoneRegion() string {
	timezone, _ := GetData("timezone", "region")
	return timezone.(string)
}

// Builder functions

func durationInRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxDuration(params...)
	if err != nil {
		return nil, err
	}
	return DurationInRange(min, max), nil
}

func durationBuilder(params ...string) (interface{}, error) {
	return Duration(), nil
}

func timeBuilder(params ...string) (interface{}, error) {
	return Time(), nil
}

func timeNowBuilder(params ...string) (interface{}, error) {
	return TimeNow(), nil
}

func nanoSecondBuilder(params ...string) (interface{}, error) {
	return NanoSecond(), nil
}

func secondBuilder(params ...string) (interface{}, error) {
	return Second(), nil
}

func minuteBuilder(params ...string) (interface{}, error) {
	return Minute(), nil
}

func hourBuilder(params ...string) (interface{}, error) {
	return Hour(), nil
}

func dayBuilder(params ...string) (interface{}, error) {
	return Day(), nil
}

func weekDayBuilder(params ...string) (interface{}, error) {
	return WeekDay(), nil
}

func monthBuilder(params ...string) (interface{}, error) {
	return Month(), nil
}

func yearBuilder(params ...string) (interface{}, error) {
	return Year(), nil
}

func timeZoneBuilder(params ...string) (interface{}, error) {
	return TimeZone(), nil
}

func timeZoneAbbrBuilder(params ...string) (interface{}, error) {
	return TimeZoneAbbr(), nil
}

func timeZoneFullBuilder(params ...string) (interface{}, error) {
	return TimeZoneFull(), nil
}

func timeZoneOffsetBuilder(params ...string) (interface{}, error) {
	return TimeZoneOffset(), nil
}

func timeZoneRegionBuilder(params ...string) (interface{}, error) {
	return TimeZoneRegion(), nil
}
