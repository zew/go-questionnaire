package trl

// CoreTranslations returns a core of transations,
// to be extended or overwritten by application config data.
// To reference app translations in questionnaires, use cfg.Get().Mp[key]
func CoreTranslations() Map {
	return coreTranslations
}

var coreTranslations = Map{
	"app_org": {
		"en": "My Org",
		"de": "Meine Organisation",
		"es": "Mi organización",
		"fr": "Mon organisation",
		"it": "La mia organizzazion",
		"pl": "Moja organizacja",
	},
	"app_label": {
		"en": "My Example App", // yes, repeat of AppName
		"de": "Meine Beispiel Anwendung",
		"es": "Mi aplicación de ejemplo",
		"fr": "Mon exemple d'application",
		"it": "La mia App esempio",
		"pl": "Moja Przykładowa aplikacja",
	},
	"yes": {
		"de": "Ja",
		"en": "Yes",
		"es": "Sí",
		"fr": "Oui",
		"it": "Sì",
		"pl": "Tak",
	},
	"no": {
		"de": "Nein",
		"en": "No",
		"es": "No",
		"fr": "Non",
		"it": "No",
		"pl": "Nie",
	},
	"questionnaire": {
		"de": "Fragebogen",
		"en": "Survey",
		"es": "Encuesta",
		"fr": "Sondage", // "Terminé cette enquête" ahv,
		"it": "Sondaggio",
		"pl": "Ankietę",
	},
	"user": {
		"de": "Benutzer",
		"en": "User",
		"es": "Usuario",
		"fr": "Usager",
		"it": "Utente",
		"pl": "Użytkownik",
	},
	"logout": {
		"de": "Abmelden",
		"en": "Logout",
		"es": "Desconectar",
		"fr": "Déconnexion",
		"it": "Disconnessione",
		"pl": "Wyloguj się",
	},
	"about": {
		"de": "Über",
		"en": "About",
		"es": "Acerca",
		"fr": "À propos",
		"it": "Sul",
		"pl": "O",
	},
	"imprint": {
		"de": "Impressum",
		"en": "Imprint",
		"es": "Empreinte",
		"fr": "Mentions légales", //"Impresión" ahv,
		"it": "Impronta",
		"pl": "Nadrukiem",
	},

	"language": {
		"de": "Sprache/<br>Language",
		"en": "Language",
		"es": "Idioma/<br>Language",
		"fr": "Langue/<br>Language",
		"it": "Linguaggio/<br>Language",
		"pl": "Język/<br>Language",
	},
	// corresponding to iso code suffix mp["lang_"+lc]:
	"lang_de": {
		"de": "Deutsch",
		"en": "German",
		"es": "Alemán",
		"fr": "Allemand",
		"it": "Tedesco",
		"pl": "Niemiecki",
	},
	"lang_en": {
		"de": "English",
		"en": "English",
		"es": "Inglés",
		"fr": "Anglais",
		"it": "Inglese",
		"pl": "Angielsku",
	},
	"lang_es": {
		"de": "Spanisch",
		"en": "Spanish",
		"es": "Español",
		"fr": "Espagnol",
		"it": "Spagnolo",
		"pl": "Hiszpańsku",
	},
	"lang_fr": {
		"de": "Französisch",
		"en": "French",
		"es": "Francés",
		"fr": "Français",
		"it": "Francese",
		"pl": "Francusku",
	},
	"lang_it": {
		"de": "Italienisch",
		"en": "Italian",
		"es": "Italiano",
		"fr": "Italien",
		"it": "Italiano",
		"pl": "Włosku",
	},
	"lang_pl": {
		"de": "Polnisch",
		"en": "Polish",
		"es": "Polaco",
		"fr": "Polonais",
		"it": "Polacco",
		"pl": "Polsku",
	},
	// navigation stuff
	"page": {
		"en": "Page",
		"de": "Seite",
		"es": "Página",
		"fr": "Page",
		"it": "Pagina",
		"pl": "Strona",
	},
	"pages": {
		"en": "Pages",
		"de": "Seiten",
		"es": "Páginas",
		"fr": "Pages",
		"it": "Pagine",
		"pl": "Stron",
	},
	"start": {
		"en": "Start",
		"de": "Start",
		"es": "Inicia",
		"fr": "Commencer",
		"it": "Inizia",
		"pl": "Uruchomić",
	},
	"previous": {
		"en": "Previous",
		"de": "Zurück",
		"es": "Previa",
		"fr": "Précédente",
		"it": "Precedente",
		"pl": "Poprzedni",
	},
	"next": {
		"en": "Next",
		"de": "Weiter",
		"es": "Continuar",
		"fr": "Continuer",
		"it": "Continuare",
		"pl": "Kontynuować",
	},
	"end": {
		"en": "End",
		"de": "Ende",
		"es": "Fin",
		"fr": "Fin",
		"it": "Fine",
		"pl": "Końcu",
	},
	// error stuff
	"valid_entry": {
		"de": "Bitte geben Sie einen gültigen Wert ein.",
		"en": "Please enter a valid value",
		"es": "Por favor ingrese una cifra válida",
		"fr": "Veuillez entrer une valeur valide",
		"it": "Per favore inserisci un valore valido",
		"pl": "Wprowadź prawidłową wartość",
	},
	"entry_range": {
		"de": "Bitte geben Sie einen Wert zwischen %v und %v ein",
		"en": "Please enter a value between %v and %v",
		"es": "Introduzca un valor entre %v y %v",
		"fr": "Veuillez saisir une valeur comprise entre %v et %v",
		"it": "Immettere un valore compreso tra %v e %v",
		"pl": "Wprowadź wartość od %v do %v",
	},
	"entry_stepping": {
		"de": "in %ver Schritten",
		"en": "in %ver steps",
		"es": "en %ver pasos",
		"fr": "en %v étapes",
		"it": "in passaggi %ver",
		"pl": "w %ver krokach",
	},
	"correct_errors": {
		"de": "Bitte korrigieren Sie die unten angezeigten Fehler.",
		"en": "Please correct the errors displayed below.",
		"es": "Por favor corrija los errores que aparecen a continuación",
		"fr": "Veuillez corriger les erreurs affichées ci-dessous",
		"it": "Per piacere correga gli errori sottostanti.",
		"pl": "Popraw błędy wyświetlane poniżej",
	},
	"not_a_number": {
		"de": "'%v' keine Zahl",
		"en": "'%v' not a number",
		"es": "'%v' no es un número",
		"fr": "'%v' pas un certain nombre",
		"it": "'%v' non è un numero",
		"pl": "'%v' nie liczba",
	},
	"too_big": {
		"de": "Max %.0f",
		"en": "max %.0f",
		"es": "Máximo %.0f",
		"fr": "Max %.0f",
		"it": "Massimo %.0f",
		"pl": "Maksymalna %.0f",
	},
	"too_small": {
		"de": "Min %.0f",
		"en": "min %.0f",
		"es": "Mínimo %.0f",
		"fr": "Min %.0f",
		"it": "Minimo %.0f",
		"pl": "Minimalne %.0f",
	},
	"must_one_option": {
		"de": "Bitte eine Option wählen",
		"en": "Please choose one option",
		"es": "Por favor elija una opción",
		"fr": "Veuillez choisir une option",
		"it": "Si prega di selezionare una opzione",
		"pl": "Proszę wybrać jedną z opcji",
	},
	"must_not_empty": {
		"de": "Eingabe erforderlich",
		"en": "Entry necessary",
		"es": "Entrada necesaria",
		"fr": "Entrée nécessaire",
		"it": "Ingresso necessario",
		"pl": "Wejście jest konieczne",
	},
	// system messages
	"login_by_hash_failed": {
		"de": "Anmeldung via Link gescheitert oder Sitzung verfallen.\nBitte nutzen Sie den übermittelten Link um sich anzumelden.\nWenn der Link in zwei Zeilen geteilt wurde, verbinden Sie die Zeilen wieder.",
		"en": "Login via link failed or session timed out.\nPlease use the provided link to login.\nIf the link was split into two lines, reconnect them.",
		"es": "Error al iniciar sesión por hash.\nPor favor, utilice el enlace proporcionado para iniciar sesión.\nSi el enlace se dividió en dos líneas, vuelva a conectarlas.",
		"fr": "Login par hachage a échoué.\nVeuillez utiliser le lien fourni pour vous connecter.\nSi le lien a été divisé en deux lignes, reconnectez-les.",
		"it": "Il login non è andato a buon fine.\nPer piacere si utilizzi il link fornitovi per effettuare il login.\nSe il link è spezzato in due, le due parti devono essere riconnesse.",
		"pl": "Logowanie przez hash nie powiodło się. \nProszę użyć przesłanego linku, aby się zarejestrować. \nJeśli łącze zostało podzielone na dwa wiersze, Połącz ponownie wiersze.",
	},
	"finish_questionnaire": {
		"de": "Fragebogen abschließen",
		"en": "Finish this survey",
		"es": "Terminé esta encuesta",
		"fr": "Terminer ce sondage", // "Terminé cette enquête" ahv,
		"it": "Finire questo sondaggio",
		"pl": "Zakończyłem tę ankietę",
	},
	"finish_save_questionnaire": {
		"de": "Fragebogen abschließen um die Daten final zu speichern.",
		"en": "Finish this survey and finalize your answers.",
	},
	"entries_saved": {
		"de": "Ihre Eingaben wurden gespeichert.",
		"en": "Your entries have been saved.",
		"es": "Sus entradas se han guardado.",
		"fr": "Vos réponses ont été sauvegardées.",
		"it": "Le Sue risposte sono state salvate.",
		"pl": "Twoje wpisy zostały zapisane.",
	},
	"thanks_for_participation": {
		"de": "Danke für Ihre Teilnahme an unserer Umfrage.",
		"en": "Thank you for your participation in our survey.",
		"es": "Gracias por haber contestado a nuestro cuestionario.",
		"fr": "Nous vous remercions d'avoir répondu à nos questions.",
		"it": "Grazie per aver risposto al nostro questionario.",
		"pl": "Dziękujemy za uczestnictwo w ankiecie.",
	},
	"finished_by_participant": {
		"de": "Sie haben den Fragebogen bereits abgeschlossen (%v).",
		"en": "You already finished this survey wave at %v",
		"es": "Usted ya terminó esta ola de encuestas en %v",
		"fr": "Vous avez déjà terminé cette vague de sondage à %v",
		"it": "Lei ha già completato questo questionario (%v)",
		"pl": "Już skończyłeś tę falę pomiarową na %v",
	},
	"deadline_exceeded": {
		"de": "Diese Umfrage wurde am %v beendet.",
		"en": "Current survey was closed at %v.",
		"es": "La encuesta actual se cerró en %v",
		"fr": "L'enquête en cours a été clôturée à %v",
		"it": "Questo questionario è stato chiuso il %v.",
		"pl": "Aktualna Ankieta została zamknięta w %v",
	},
	"percentage_answered": {
		"de": "Sie haben %v von %v Fragen beantwortet: %2.1f&nbsp;Prozent.  <br>\n",
		"en": "You answered %v out of %v questions: %2.1f&nbsp;percent.  <br>\n",
		"es": "Usted contestó %v de %v preguntas: %2.1f por ciento. <br>\n",
		"fr": "Vous avez répondu %v sur %v questions: %2.1f pour cent. <br>\n",
		"it": "Lei ha risposto a %v domande su %v: %2.1f per cento.  <br>\n",
		"pl": "Odpowiedziałeś %v na %v pytania: %2.1f procent. <br>\n",
	},
	"survey_ending": {
		"de": "Umfrage endet am %v. <br>\nVeröffentlichung am %v.  <br>\n",
		"en": "Survey will finish at %v. <br>\nPublication will be at %v.<br>\n",
		"es": "La encuesta terminará en %v.\nPublicación será en %v. <br>\n",
		"fr": "L'enquête se terminera à %v.\nPublication sera à %v. <br>\n",
		"it": "Il sondaggio verrà concluso il %v. <br>\nLa pubblicazione avverrà il %v.<br>\n",
		"pl": "Ankieta zakończy się w %v.\nPublikacja będzie %v. <br>\n",
	},
	"review_by_personal_link": {
		"de": "Sie können Ihre Daten jederzeit über Ihren persönlichen Link prüfen/ändern. <br>\n",
		"en": "You may review or change your data using your personal link. <br>\n",
		"es": "Usted puede revisar o cambiar sus datos usando su enlace personal.<br>\n",
		"fr": "Vous pouvez consulter ou modifier vos données à l'aide de votre lien personnel.<br>\n",
		"it": "Può rivedere o modificare i suoi dati usando il Suo link personale. <br>\n",
		"pl": "Dane można przejrzeć lub zmienić przy użyciu osobistego łącza.<br>\n",
	},
	"link_to_previous_page": {
		"de": "<a href='/?submitBtn=prev'>Zurück</a><br>\n",
		"en": "<a href='/?submitBtn=prev'>Back</a><br>\n",
		"es": "<a href='/?submitBtn=prev'>Atrás</a><br>\n",
		"fr": "<a href='/?submitBtn=prev'>Précédent</a><br>\n",
		"it": "<a href='/?submitBtn=prev'>Indietro</a><br>\n",
		"pl": "<a href='/?submitBtn=prev'>Wstecz</a><br>\n",
	},
	"review_by_permalink": {
		"de": `
		<ul  style='margin-top: -1.2rem'>
			<li>
			<!-- Bis zum Umfrage-Ende-->
			Sie können den Fragebogen
			über folgenden Link erneut aufrufen: 
			<a href='%v'>%v</a>.   
			<br>
			(Vielleicht wollen Sie sich diesen Link kopieren. Sie müssten sonst wieder von vorne beginnen.)
			</li>
		</ul>			
		`,
		"en": `
		<ul  style='margin-top: -1.2rem'>
			<li>
			Until the end of the survey,
			you can change your entries
			using following link: 
			<a href='%v'>%v</a>.   
			<br>
			(Maybe you want to copy this link)
			</li>
		</ul>
		`,
	},
}
