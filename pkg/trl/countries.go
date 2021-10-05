package trl

import "sort"

// Countries is contains translations of country names by ISO code
var Countries = map[string]S{
	"AF":             {"en": "Afghanistan", "fr": "Afghanistan", "de": "Afghanistan", "it": "Afghanistan"},
	"AL":             {"en": "Albania", "fr": "Albanie", "de": "Albanien", "it": "Albania"},
	"DZ":             {"en": "Algeria", "fr": "Algérie", "de": "Algerien", "it": "Algeria"},
	"AS":             {"en": "American Samoa", "fr": "Samoa américaines", "de": "Andorra", "it": "Andorra"},
	"AD":             {"en": "Andorra", "fr": "Andorre", "de": "Andorra", "it": "Andorra"},
	"AO":             {"en": "Angola", "fr": "Angola", "de": "Angola", "it": "Angola"},
	"AG":             {"en": "Antigua and Barbuda", "fr": "Antigua-et-Barbuda", "de": "Antigua und Barbuda", "it": "Antigua e Barbuda"},
	"AR":             {"en": "Argentina", "fr": "Argentine", "de": "Argentinien", "it": "Argentina"},
	"AM":             {"en": "Armenia", "fr": "Arménie", "de": "Armenien", "it": "Armenia"},
	"AW":             {"en": "Aruba", "fr": "Aruba", "de": "Aruba", "it": "Aruba"},
	"AU":             {"en": "Australia", "fr": "Australie", "de": "Australien", "it": "Australia"},
	"AT":             {"en": "Austria", "fr": "Autriche", "de": "Österreich", "it": "Austria"},
	"AZ":             {"en": "Azerbaijan", "fr": "Azerbaïdjan", "de": "Aserbaidschan", "it": "Azerbaigian"},
	"BS":             {"en": "Bahamas, The", "fr": "Bahamas", "de": "Bahamas", "it": "Bahamas"},
	"BH":             {"en": "Bahrain", "fr": "Bahreïn", "de": "Bahrain", "it": "Bahrein"},
	"BD":             {"en": "Bangladesh", "fr": "Bangladesh", "de": "Bangladesch", "it": "Bangladesh"},
	"BB":             {"en": "Barbados", "fr": "Barbade", "de": "Barbados", "it": "Barbados"},
	"BY":             {"en": "Belarus", "fr": "Bélarus", "de": "Belarus", "it": "Bielorussia"},
	"BE":             {"en": "Belgium", "fr": "Belgique", "de": "Belgien", "it": "Belgio"},
	"BZ":             {"en": "Belize", "fr": "Belize", "de": "Belize", "it": "Belize"},
	"BJ":             {"en": "Benin", "fr": "Bénin", "de": "Benin", "it": "Benin"},
	"BM":             {"en": "Bermuda", "fr": "Bermudes", "de": "Bermuda", "it": "Bermuda"},
	"BT":             {"en": "Bhutan", "fr": "Bhoutan", "de": "Bhutan", "it": "Bhutan"},
	"BO":             {"en": "Bolivia", "fr": "Bolivie", "de": "Bolivien", "it": "Bolivia"},
	"BA":             {"en": "Bosnia and Herzegovina", "fr": "Bosnie-Herzégovine", "de": "Bosnien und Herzegowina", "it": "Bosnia ed Erzegovina"},
	"BW":             {"en": "Botswana", "fr": "Botswana", "de": "Botsuana", "it": "Botswana"},
	"BR":             {"en": "Brazil", "fr": "Brésil", "de": "Brasilien", "it": "Brasile"},
	"VG":             {"en": "British Virgin Islands", "fr": "Îles Vierges britanniques", "de": "Britische Jungferninseln", "it": "Isole Vergini britanniche"},
	"BN":             {"en": "Brunei Darussalam", "fr": "Brunéi Darussalam", "de": "Brunei Darussalam", "it": "Brunei"},
	"BG":             {"en": "Bulgaria", "fr": "Bulgarie", "de": "Bulgarien", "it": "Bulgaria"},
	"BF":             {"en": "Burkina Faso", "fr": "Burkina Faso", "de": "Burkina Faso", "it": "Burkina Faso"},
	"BI":             {"en": "Burundi", "fr": "Burundi", "de": "Burundi", "it": "Burundi"},
	"CV":             {"en": "Cabo Verde", "fr": "Cabo Verde", "de": "Cabo Verde", "it": "Capo Verde"},
	"KH":             {"en": "Cambodia", "fr": "Cambodge", "de": "Kambodscha", "it": "Cambogia"},
	"CM":             {"en": "Cameroon", "fr": "Cameroun", "de": "Kamerun", "it": "Camerun"},
	"CA":             {"en": "Canada", "fr": "Canada", "de": "Kanada", "it": "Canada"},
	"KY":             {"en": "Cayman Islands", "fr": "Îles Caïmans", "de": "Caymaninseln", "it": "Isole Cayman"},
	"CF":             {"en": "Central African Republic", "fr": "République centrafricaine", "de": "Zentralafrikanische Republik", "it": "Repubblica Centrafricana"},
	"TD":             {"en": "Chad", "fr": "Tchad", "de": "Tschad", "it": "Ciad"},
	"GG":             {"en": "Channel Islands", "fr": "Îles Anglo-Normandes", "de": "Kanalinseln", "it": "Isole del Canale"},
	"CL":             {"en": "Chile", "fr": "Chili", "de": "Chile", "it": "Cile"},
	"CN":             {"en": "China", "fr": "Chine", "de": "China", "it": "Cina"},
	"CO":             {"en": "Colombia", "fr": "Colombie", "de": "Kolumbien", "it": "Colombia"},
	"KM":             {"en": "Comoros", "fr": "Comores", "de": "Komoren", "it": "Comore"},
	"CD":             {"en": "Congo, Dem. Rep.", "fr": "Congo, République démocratique du", "de": "Kongo, Demokratische Republik", "it": "Repubblica Democratica del Congo"},
	"CG":             {"en": "Congo, Rep.", "fr": "Congo, République du", "de": "Kongo", "it": "Repubblica del Congo"},
	"CK":             {"en": "Cook Islands", "fr": "Îles Cook", "de": "Cookinseln", "it": "Isole Cook"},
	"CR":             {"en": "Costa Rica", "fr": "Costa Rica", "de": "Costa Rica", "it": "Costa Rica"},
	"CI":             {"en": "Cote d'Ivoire", "fr": "Côte d'Ivoire", "de": "Côte d'Ivoire", "it": "Costa d'Avorio"},
	"HR":             {"en": "Croatia", "fr": "Croatie", "de": "Kroatien", "it": "Croazia"},
	"CU":             {"en": "Cuba", "fr": "Cuba", "de": "Kuba", "it": "Cuba"},
	"CW":             {"en": "Curacao", "fr": "Curacao", "de": "Curacao", "it": "Curacao"},
	"CY":             {"en": "Cyprus", "fr": "Chypre", "de": "Zypern", "it": "Cipro"},
	"CZ":             {"en": "Czech Republic", "fr": "République tchèque", "de": "Tschechische Republik", "it": "Repubblica Ceca"},
	"DK":             {"en": "Denmark", "fr": "Danemark", "de": "Dänemark", "it": "Danimarca"},
	"DJ":             {"en": "Djibouti", "fr": "Djibouti", "de": "Dschibuti", "it": "Gibuti"},
	"DM":             {"en": "Dominica", "fr": "Dominique", "de": "Dominica", "it": "Dominica"},
	"DO":             {"en": "Dominican Republic", "fr": "République dominicaine", "de": "Dominikanische Republik", "it": "Repubblica Dominicana"},
	"EC":             {"en": "Ecuador", "fr": "Équateur", "de": "Ecuador", "it": "Ecuador"},
	"EG":             {"en": "Egypt, Arab Rep.", "fr": "Égypte, République arabe d’", "de": "Ägypten", "it": "Egitto"},
	"SV":             {"en": "El Salvador", "fr": "El Salvador", "de": "El Salvador", "it": "El Salvador"},
	"GQ":             {"en": "Equatorial Guinea", "fr": "Guinée équatoriale", "de": "Äquatorialguinea", "it": "Guinea Equatoriale"},
	"ER":             {"en": "Eritrea", "fr": "Érythrée", "de": "Eritrea", "it": "Eritrea"},
	"EE":             {"en": "Estonia", "fr": "Estonie", "de": "Estland", "it": "Estonia"},
	"SZ":             {"en": "Eswatini", "fr": "Eswatini", "de": "Swasiland", "it": "Swaziland"},
	"ET":             {"en": "Ethiopia", "fr": "Éthiopie", "de": "Äthiopien", "it": "Etiopia"},
	"FO":             {"en": "Faroe Islands", "fr": "Îles Féroé", "de": "Färöer-Inseln", "it": "Isole Faroe"},
	"FJ":             {"en": "Fiji", "fr": "Fidji", "de": "Fidschi", "it": "Figi"},
	"FI":             {"en": "Finland", "fr": "Finlande", "de": "Finnland", "it": "Finlandia"},
	"FR":             {"en": "France", "fr": "France", "de": "Frankreich", "it": "Francia"},
	"PF":             {"en": "French Polynesia", "fr": "Polynésie française", "de": "Französisch-Polynesien", "it": "Polinesia francese"},
	"GA":             {"en": "Gabon", "fr": "Gabon", "de": "Gabun", "it": "Gabon"},
	"GM":             {"en": "Gambia, The", "fr": "Gambie", "de": "Gambia", "it": "Gambia"},
	"GE":             {"en": "Georgia", "fr": "Géorgie", "de": "Georgien", "it": "Georgia"},
	"DE":             {"en": "Germany", "fr": "Allemagne", "de": "Deutschland", "it": "Germania"},
	"GH":             {"en": "Ghana", "fr": "Ghana", "de": "Ghana", "it": "Ghana"},
	"GI":             {"en": "Gibraltar", "fr": "Gibraltar", "de": "Gibraltar", "it": "Gibilterra"},
	"GR":             {"en": "Greece", "fr": "Grèce", "de": "Griechenland", "it": "Grecia"},
	"GL":             {"en": "Greenland", "fr": "Groenland", "de": "Grönland", "it": "Groenlandia"},
	"GD":             {"en": "Grenada", "fr": "Grenade", "de": "Grenada", "it": "Grenada"},
	"GU":             {"en": "Guam", "fr": "Guam", "de": "Guam", "it": "Guam"},
	"GT":             {"en": "Guatemala", "fr": "Guatemala", "de": "Guatemala", "it": "Guatemala"},
	"GN":             {"en": "Guinea", "fr": "Guinée", "de": "Guinea", "it": "Guinea"},
	"GW":             {"en": "Guinea-Bissau", "fr": "Guinée-Bissau", "de": "Guinea-Bissau", "it": "Guinea-Bissau"},
	"GY":             {"en": "Guyana", "fr": "Guyane", "de": "Guyana", "it": "Guyana"},
	"HT":             {"en": "Haiti", "fr": "Haïti", "de": "Haiti", "it": "Haiti"},
	"HN":             {"en": "Honduras", "fr": "Honduras", "de": "Honduras", "it": "Honduras"},
	"HK":             {"en": "Hong Kong SAR, China", "fr": "Chine, RAS de Hong Kong", "de": "Hongkong", "it": "Hong Kong"},
	"HU":             {"en": "Hungary", "fr": "Hongrie", "de": "Ungarn", "it": "Ungheria"},
	"IS":             {"en": "Iceland", "fr": "Islande", "de": "Island", "it": "Islanda"},
	"IN":             {"en": "India", "fr": "Inde", "de": "Indien", "it": "India"},
	"ID":             {"en": "Indonesia", "fr": "Indonésie", "de": "Indonesien", "it": "Indonesia"},
	"IR":             {"en": "Iran, Islamic Rep.", "fr": "Iran, République islamique d’", "de": "Iran", "it": "Iran"},
	"IQ":             {"en": "Iraq", "fr": "Iraq", "de": "Irak", "it": "Iraq"},
	"IE":             {"en": "Ireland", "fr": "Irlande", "de": "Irland", "it": "Irlanda"},
	"IM":             {"en": "Isle of Man", "fr": "Île de Man", "de": "Isle of Man", "it": "Isola di Man"},
	"IL":             {"en": "Israel", "fr": "Israël", "de": "Israel", "it": "Israele"},
	"IT":             {"en": "Italy", "fr": "Italie", "de": "Italien", "it": "Italia"},
	"JM":             {"en": "Jamaica", "fr": "Jamaïque", "de": "Jamaika", "it": "Giamaica"},
	"JP":             {"en": "Japan", "fr": "Japon", "de": "Japan", "it": "Giappone"},
	"JO":             {"en": "Jordan", "fr": "Jordanie", "de": "Jordanien", "it": "Giordania"},
	"KZ":             {"en": "Kazakhstan", "fr": "Kazakhstan", "de": "Kasachstan", "it": "Kazakistan"},
	"KE":             {"en": "Kenya", "fr": "Kenya", "de": "Kenia", "it": "Kenya"},
	"KI":             {"en": "Kiribati", "fr": "Kiribati", "de": "Kiribati", "it": "Kiribati"},
	"KP":             {"en": "Korea, Dem. People’s Rep.", "fr": "Corée, République démocratique de", "de": "Nordkorea", "it": "Corea del Nord"},
	"KR":             {"en": "Korea, Rep.", "fr": "Corée, République de", "de": "Südkorea", "it": "Corea del Sud"},
	"XK":             {"en": "Kosovo", "fr": "Kosovo", "de": "Kosovo", "it": "Kosovo"},
	"KW":             {"en": "Kuwait", "fr": "Koweït", "de": "Kuwait", "it": "Kuwait"},
	"KG":             {"en": "Kyrgyzstan", "fr": "Kirghizistan", "de": "Kirgisistan", "it": "Kirghizistan"},
	"LA":             {"en": "Lao PDR", "fr": "République démocratique populaire lao", "de": "Laos", "it": "Laos"},
	"LV":             {"en": "Latvia", "fr": "Lettonie", "de": "Lettland", "it": "Lettonia"},
	"LB":             {"en": "Lebanon", "fr": "Liban", "de": "Libanon", "it": "Libano"},
	"LS":             {"en": "Lesotho", "fr": "Lesotho", "de": "Lesotho", "it": "Lesotho"},
	"LR":             {"en": "Liberia", "fr": "Libéria", "de": "Liberia", "it": "Liberia"},
	"LY":             {"en": "Libya", "fr": "Libye", "de": "Libyen", "it": "Libia"},
	"LI":             {"en": "Liechtenstein", "fr": "Liechtenstein", "de": "Liechtenstein", "it": "Liechtenstein"},
	"LT":             {"en": "Lithuania", "fr": "Lituanie", "de": "Litauen", "it": "Lituania"},
	"LU":             {"en": "Luxembourg", "fr": "Luxembourg", "de": "Luxemburg", "it": "Lussemburgo"},
	"MO":             {"en": "Macao SAR, China", "fr": "Région administrative spéciale de Macao, Chine", "de": "Macau", "it": "Macao"},
	"MK":             {"en": "Macedonia, FYR", "fr": "Macédoine, ex-République yougoslave de", "de": "Mazedonien", "it": "Macedonia"},
	"MG":             {"en": "Madagascar", "fr": "Madagascar", "de": "Madagaskar", "it": "Madagascar"},
	"MW":             {"en": "Malawi", "fr": "Malawi", "de": "Malawi", "it": "Malawi"},
	"MY":             {"en": "Malaysia", "fr": "Malaisie", "de": "Malaysia", "it": "Malesia"},
	"MV":             {"en": "Maldives", "fr": "Maldives", "de": "Malediven", "it": "Maldive"},
	"ML":             {"en": "Mali", "fr": "Mali", "de": "Mali", "it": "Mali"},
	"MT":             {"en": "Malta", "fr": "Malte", "de": "Malta", "it": "Malta"},
	"MH":             {"en": "Marshall Islands", "fr": "Îles Marshall", "de": "Marshallinseln", "it": "Isole Marshall"},
	"MR":             {"en": "Mauritania", "fr": "Mauritanie", "de": "Mauretanien", "it": "Mauritania"},
	"MU":             {"en": "Mauritius", "fr": "Maurice", "de": "Mauritius", "it": "Mauritius"},
	"MX":             {"en": "Mexico", "fr": "Mexique", "de": "Mexiko", "it": "Messico"},
	"FM":             {"en": "Micronesia, Fed. Sts.", "fr": "Micronésie, États fédérés de", "de": "Mikronesien", "it": "Micronesia"},
	"MD":             {"en": "Moldova", "fr": "Moldova", "de": "Moldau", "it": "Moldavia"},
	"MC":             {"en": "Monaco", "fr": "Monaco", "de": "Monaco", "it": "Monaco"},
	"MN":             {"en": "Mongolia", "fr": "Mongolie", "de": "Mongolei", "it": "Mongolia"},
	"ME":             {"en": "Montenegro", "fr": "Monténégro", "de": "Montenegro", "it": "Montenegro"},
	"MA":             {"en": "Morocco", "fr": "Maroc", "de": "Marokko", "it": "Marocco"},
	"MZ":             {"en": "Mozambique", "fr": "Mozambique", "de": "Mosambik", "it": "Mozambico"},
	"MM":             {"en": "Myanmar", "fr": "Myanmar", "de": "Myanmar", "it": "Myanmar o Birmania"},
	"NA":             {"en": "Namibia", "fr": "Namibie", "de": "Namibia", "it": "Namibia"},
	"NR":             {"en": "Nauru", "fr": "Nauru", "de": "Nauru", "it": "Nauru"},
	"NP":             {"en": "Nepal", "fr": "Népal", "de": "Nepal", "it": "Nepal"},
	"NL":             {"en": "Netherlands", "fr": "Pays-Bas", "de": "Niederlande", "it": "Paesi Bassi"},
	"NC":             {"en": "New Caledonia", "fr": "Nouvelle-Calédonie", "de": "Neukaledonien", "it": "Caledonia"},
	"NZ":             {"en": "New Zealand", "fr": "Nouvelle-Zélande", "de": "Neuseeland", "it": "Nuova Zelanda"},
	"NI":             {"en": "Nicaragua", "fr": "Nicaragua", "de": "Nicaragua", "it": "Nicaragua"},
	"NE":             {"en": "Niger", "fr": "Niger", "de": "Niger", "it": "Niger"},
	"NG":             {"en": "Nigeria", "fr": "Nigéria", "de": "Nigeria", "it": "Nigeria"},
	"MP":             {"en": "Northern Mariana Islands", "fr": "Mariannes", "de": "Nördliche Marianen", "it": "Isole Marianne Settentrionali"},
	"NO":             {"en": "Norway", "fr": "Norvège", "de": "Norwegen", "it": "Norvegia"},
	"OM":             {"en": "Oman", "fr": "Oman", "de": "Oman", "it": "Oman"},
	"PK":             {"en": "Pakistan", "fr": "Pakistan", "de": "Pakistan", "it": "Pakistan"},
	"PW":             {"en": "Palau", "fr": "Palaos", "de": "Palau", "it": "Palau"},
	"PA":             {"en": "Panama", "fr": "Panama", "de": "Panama", "it": "Panama"},
	"PG":             {"en": "Papua New Guinea", "fr": "Papouasie-Nouvelle-Guinée", "de": "Papua-Neuguinea", "it": "Papua Nuova Guinea"},
	"PY":             {"en": "Paraguay", "fr": "Paraguay", "de": "Paraguay", "it": "Paraguay"},
	"PE":             {"en": "Peru", "fr": "Pérou", "de": "Peru", "it": "Perù"},
	"PH":             {"en": "Philippines", "fr": "Philippines", "de": "Philippinen", "it": "Filippine"},
	"PL":             {"en": "Poland", "fr": "Pologne", "de": "Polen", "it": "Polonia"},
	"PT":             {"en": "Portugal", "fr": "Portugal", "de": "Portugal", "it": "Portogallo"},
	"PR":             {"en": "Puerto Rico", "fr": "Porto Rico", "de": "Puerto Rico", "it": "Porto Rico"},
	"QA":             {"en": "Qatar", "fr": "Qatar", "de": "Katar", "it": "Qatar"},
	"RO":             {"en": "Romania", "fr": "Roumanie", "de": "Rumänien", "it": "Romania"},
	"RU":             {"en": "Russian Federation", "fr": "Fédération de Russie", "de": "Russische Föderation", "it": "Russia"},
	"RW":             {"en": "Rwanda", "fr": "Rwanda", "de": "Ruanda", "it": "Ruanda"},
	"WS":             {"en": "Samoa", "fr": "Samoa", "de": "Samoa", "it": "Samoa"},
	"SM":             {"en": "San Marino", "fr": "Saint-Marin", "de": "San Marino", "it": "San Marino"},
	"ST":             {"en": "Sao Tome and Principe", "fr": "Sao Tomé-et-Principe", "de": "São Tomé und Príncipe", "it": "São Tomé e Príncipe"},
	"SA":             {"en": "Saudi Arabia", "fr": "Arabie saoudite", "de": "Saudi-Arabien", "it": "Arabia Saudita"},
	"SN":             {"en": "Senegal", "fr": "Sénégal", "de": "Senegal", "it": "Senegal"},
	"RS":             {"en": "Serbia", "fr": "Serbie", "de": "Serbien", "it": "Serbia"},
	"SC":             {"en": "Seychelles", "fr": "Seychelles", "de": "Seychellen", "it": "Seychelles"},
	"SL":             {"en": "Sierra Leone", "fr": "Sierra Leone", "de": "Sierra Leone", "it": "Sierra Leone"},
	"SG":             {"en": "Singapore", "fr": "Singapour", "de": "Singapur", "it": "Singapore"},
	"SX":             {"en": "Sint Maarten (Dutch part)", "fr": "Sint Maarten (Dutch part)", "de": "Sint Maarten", "it": "Sint Maarten"},
	"SK":             {"en": "Slovak Republic", "fr": "République slovaque", "de": "Slowakei", "it": "Slovacchia"},
	"SI":             {"en": "Slovenia", "fr": "Slovénie", "de": "Slowenien", "it": "Slovenia"},
	"SB":             {"en": "Solomon Islands", "fr": "Îles Salomon", "de": "Salomonen", "it": "Isole Salomone"},
	"SO":             {"en": "Somalia", "fr": "Somalie", "de": "Somalia", "it": "Somalia"},
	"ZA":             {"en": "South Africa", "fr": "Afrique du Sud", "de": "Südafrika", "it": "Sudafrica"},
	"SS":             {"en": "South Sudan", "fr": "Soudan du Sud", "de": "Südsudan", "it": "Sudan del Sud"},
	"ES":             {"en": "Spain", "fr": "Espagne", "de": "Spanien", "it": "Spagna"},
	"LK":             {"en": "Sri Lanka", "fr": "Sri Lanka", "de": "Sri Lanka", "it": "Sri Lanka"},
	"LC":             {"en": "St. Kitts and Nevis", "fr": "Saint-Kitts-et-Nevis", "de": "St. Kitts und Nevis", "it": "Saint Kitts e Nevis"},
	"KN":             {"en": "St. Lucia", "fr": "Sainte-Lucie", "de": "St. Lucia", "it": "Santa Lucia"},
	"MF":             {"en": "St. Martin (French part)", "fr": "Saint-Martin (fr)", "de": "Saint-Martin", "it": "Saint Martin"},
	"VC":             {"en": "St. Vincent and the Grenadines", "fr": "Saint-Vincent-et-les Grenadines", "de": "St. Vincent und die Grenadinen", "it": "Saint Vincent e Grenadine"},
	"SD":             {"en": "Sudan", "fr": "Soudan", "de": "Sudan", "it": "Sudan"},
	"SR":             {"en": "Suriname", "fr": "Suriname", "de": "Suriname", "it": "Suriname"},
	"SE":             {"en": "Sweden", "fr": "Suède", "de": "Schweden", "it": "Svezia"},
	"CH":             {"en": "Switzerland", "fr": "Suisse", "de": "Schweiz", "it": "Svizzera"},
	"SY":             {"en": "Syrian Arab Republic", "fr": "République arabe syrienne", "de": "Syrien", "it": "Siria"},
	"TJ":             {"en": "Tajikistan", "fr": "Tadjikistan", "de": "Tadschikistan", "it": "Tagikistan"},
	"TZ":             {"en": "Tanzania", "fr": "Tanzanie", "de": "Tansania", "it": "Tanzania"},
	"TH":             {"en": "Thailand", "fr": "Thaïlande", "de": "Thailand", "it": "Thailandia"},
	"TL":             {"en": "Timor-Leste", "fr": "Timor-Leste", "de": "Timor-Leste", "it": "Timor Est"},
	"TG":             {"en": "Togo", "fr": "Togo", "de": "Togo", "it": "Togo"},
	"TO":             {"en": "Tonga", "fr": "Tonga", "de": "Tonga", "it": "Tonga"},
	"TT":             {"en": "Trinidad and Tobago", "fr": "Trinité-et-Tobago", "de": "Trinidad und Tobago", "it": "Trinidad e Tobago"},
	"TN":             {"en": "Tunisia", "fr": "Tunisie", "de": "Tunesien", "it": "Tunisia"},
	"TR":             {"en": "Turkey", "fr": "Turquie", "de": "Türkei", "it": "Turchia"},
	"TM":             {"en": "Turkmenistan", "fr": "Turkménistan", "de": "Turkmenistan", "it": "Turkmenistan"},
	"TC":             {"en": "Turks and Caicos Islands", "fr": "Îles Turques-et-Caïques", "de": "Turks- und Caicosinseln", "it": "Isole Turks e Caicos"},
	"TV":             {"en": "Tuvalu", "fr": "Tuvalu", "de": "Tuvalu", "it": "Tuvalu"},
	"TW":             {"en": "Taiwan", "fr": "Taïwan", "de": "Taiwan", "it": "Taiwan"},
	"UG":             {"en": "Uganda", "fr": "Ouganda", "de": "Uganda", "it": "Uganda"},
	"UA":             {"en": "Ukraine", "fr": "Ukraine", "de": "Ukraine", "it": "Ucraina"},
	"AE":             {"en": "United Arab Emirates", "fr": "Émirats arabes unis", "de": "Vereinigte Arabische Emirate", "it": "Emirati Arabi Uniti"},
	"GB":             {"en": "United Kingdom", "fr": "Royaume-Uni", "de": "Großbritannien", "it": "Regno Unito"},
	"US":             {"en": "United States", "fr": "États-Unis", "de": "Vereinigte Staaten", "it": "Stati Uniti"},
	"UY":             {"en": "Uruguay", "fr": "Uruguay", "de": "Uruguay", "it": "Uruguay"},
	"UZ":             {"en": "Uzbekistan", "fr": "Ouzbékistan", "de": "Usbekistan", "it": "Uzbekistan"},
	"VU":             {"en": "Vanuatu", "fr": "Vanuatu", "de": "Vanuatu", "it": "Vanuatu"},
	"VE":             {"en": "Venezuela, RB", "fr": "Venezuela", "de": "Venezuela", "it": "Venezuela"},
	"VN":             {"en": "Vietnam", "fr": "Viet Nam", "de": "Vietnam", "it": "Vietnam"},
	"VI":             {"en": "Virgin Islands (U.S.)", "fr": "Îles Vierges (EU)", "de": "Amerikanische Jungferninseln", "it": "Isole Vergini"},
	"west_bank_gaza": {"en": "West Bank and Gaza", "fr": "Cisjordanie et Gaza", "de": "Westjordanland", "it": "Cisgiordania e Gaza"},
	"YE":             {"en": "Yemen, Rep.", "fr": "Yémen, Rép. du", "de": "Jemen", "it": "Yemen"},
	"ZM":             {"en": "Zambia", "fr": "Zambie", "de": "Sambia", "it": "Zambia"},
	"ZW":             {"en": "Zimbabwe", "fr": "Zimbabwe", "de": "Simbabwe", "it": "Zimbabwe"},
}

// CountryISOs provides some stable default sorting
var CountryISOs = []string{}

func init() {
	for k := range Countries {
		CountryISOs = append(CountryISOs, k)
	}
	sort.Strings(CountryISOs)
}

// FederalStatesGermany for multiple questionnaires
var FederalStatesGermany = map[string]S{
	"BW": {"de": "Baden-Württemberg"},
	"BY": {"de": "Bayern"},
	"BE": {"de": "Berlin"},
	"BB": {"de": "Brandenburg"},
	"HB": {"de": "Bremen"},
	"HH": {"de": "Hamburg"},
	"HE": {"de": "Hessen"},
	"NI": {"de": "Niedersachsen"},
	"MV": {"de": "Mecklenburg-Vorpommern"},
	"NW": {"de": "Nordrhein-Westfalen"},
	"RP": {"de": "Rheinland-Pfalz"},
	"SL": {"de": "Saarland"},
	"SN": {"de": "Sachsen"},
	"ST": {"de": "Sachsen-Anhalt"},
	"SH": {"de": "Schleswig-Holstein"},
	"TH": {"de": "Thüringen"},
}

// FederalStatesGermanyISOs - sorted by ISO code
var FederalStatesGermanyISOs = []string{}

func init() {
	for k := range FederalStatesGermany {
		FederalStatesGermanyISOs = append(FederalStatesGermanyISOs, k)
	}
	sort.Strings(FederalStatesGermanyISOs)
}

type sorterDe struct {
	Key string
	S   S
}

type sorterDeSl []sorterDe

func (s sorterDeSl) Len() int {
	return len(s)
}

func (s sorterDeSl) Less(i, j int) bool {
	return s[i].S["de"] < s[j].S["de"]
}

func (s sorterDeSl) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// FederalStatesGermanyISOs2 - sorted by German label
var FederalStatesGermanyISOs2 = sorterDeSl{}

func init() {
	for k, v := range FederalStatesGermany {
		FederalStatesGermanyISOs2 = append(FederalStatesGermanyISOs2, sorterDe{k, v})
	}
	sort.Sort(FederalStatesGermanyISOs2)
}
