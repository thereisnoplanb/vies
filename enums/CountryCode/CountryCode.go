package CountryCode

type Enum string

const Austria Enum = "AT"
const Belgium Enum = "BE"
const Bulgaria Enum = "BG"
const Croatia Enum = "HR"
const Cyprus Enum = "CY"
const Czechia Enum = "CZ"
const Denmark Enum = "DK"
const Estonia Enum = "EE"
const Finland Enum = "FI"
const France Enum = "FR"
const Germany Enum = "DE"
const Greece Enum = "EL"
const Hungary Enum = "HU"
const Ireland Enum = "IE"
const Italy Enum = "IT"
const Latvia Enum = "LV"
const Lithuania Enum = "LT"
const Luxembourg Enum = "LU"
const Malta Enum = "MT"
const Netherlands Enum = "NL"
const NorthernIreland Enum = "XI"
const Poland Enum = "PL"
const Portugal Enum = "PT"
const Romania Enum = "RO"
const Slovakia Enum = "SK"
const Slovenia Enum = "SI"
const Spain Enum = "ES"
const Sweden Enum = "SE"

var names map[Enum]string = map[Enum]string{
	Austria:         "AT",
	Belgium:         "BE",
	Bulgaria:        "BG",
	Croatia:         "HR",
	Cyprus:          "CY",
	Czechia:         "CZ",
	Denmark:         "DK",
	Estonia:         "EE",
	Finland:         "FI",
	France:          "FR",
	Germany:         "DE",
	Greece:          "EL",
	Hungary:         "HU",
	Ireland:         "IE",
	Italy:           "IT",
	Latvia:          "LV",
	Lithuania:       "LT",
	Luxembourg:      "LU",
	Malta:           "MT",
	Netherlands:     "NL",
	NorthernIreland: "XI",
	Poland:          "PL",
	Portugal:        "PT",
	Romania:         "RO",
	Slovakia:        "SK",
	Slovenia:        "SI",
	Spain:           "ES",
	Sweden:          "SE",
}

func (enum Enum) String() string {
	if name, ok := names[enum]; ok {
		return name
	}
	return "Unknown"
}

func (enum Enum) IsDefined() bool {
	_, ok := names[enum]
	return ok
}

func GetValues() []Enum {
	return []Enum{
		Austria,
		Belgium,
		Bulgaria,
		Croatia,
		Cyprus,
		Czechia,
		Denmark,
		Estonia,
		Finland,
		France,
		Greece,
		Spain,
		Netherlands,
		Ireland,
		Lithuania,
		Luxembourg,
		Latvia,
		Malta,
		Germany,
		Poland,
		Portugal,
		Slovakia,
		Slovenia,
		Sweden,
		Hungary,
		Italy,
		NorthernIreland,
	}
}
