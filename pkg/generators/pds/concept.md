# Prime Capital AG

2006 gegründet

Alternative Anlageklassen
Private debt

Sebastian Stehling - Dept Leader
Jan Mellert  - Investment Advisory and Solutions
Daniel Huss  - Leiter Manager Selektion

Corp lending debt - 80 Teilnehmer
Infrastruct - 30
Real estate - 80

Unternehmen
    Drei Assetklassen
    Fond - Manager - Kontaktperson

URL - private-debt-survey.zew.de


## Questions 2022-11-14

<https://private-debt-survey.zew.de/a>

* Alles in Englisch?  
  Oder auch 'ne deutsche Fassung? 

* strategy as checkbox  
  asset class as checkbox  
  repeat portfolio questions for each strategy/assestclass dynamically?

### Slider

* Browser+HTML - very free  
  Builtin restrictions to styling  

* Slider for interval

* Slider - mobile

Infrastructure

### Questions 2022-11-23

* New iteration of prototype
  * Single survey - also for real estate and infrastruct debt;  
      good for participants who do multiple asset classes;  
      good for simplicity of emails tasks and results processing;  
      `Real estate debt` and `Inftrastructure debt` can be hiddden in Q1 23
  * Strategies/tranches as dependent sub-choice of asset class
  * Strategies/tranches in columns structure
  * Not yet optimized for mobile
  * No feedback for slider questions yet;  
     I go ahead with a suggestion, because time is running out
  
  * Complete prototype for `Corp lending` by middle of next week;  
     submit to customer for feedback;  
     no mobile optimization yet;  
     sliders without feedback/final decision


* Communication with customer
  * When will the two other surveys start - for `Real estate debt` and `Inftrastructure debt`? Q1 2023 or later?
  * DSGVO text  
     plus Thomas Wirth approval
  * Introduction text
  * Invitation text of emails
  * Invitation mode: personal link
  * No registration, login yet in Q1 23

  * Last round of change requests before internal preview on 10. Dec
  * Official announcement of internal preview release on 10. Dec


## RC3

Done

* Page2: Employee-Number for each asset class
* All pages: Tranche names aligned
* Page2: Question "Weeks to closing": As dropdown
* Sliders: Display "from - to" width changes as needed
* Sliders: Ticks complete
* Sliders: Init state and no-answer state complete
* Tested on Firefox, Chrome, Edge, Opera, Mac-Safari;  
  desktop versions only
* Internal names of all questions sequentially numbered

Todo

* Frage an Prime Capital zum Konzept-Update 2022-11-28_ZEW_Survey.xlsx:
   Die Reiter 'Real Estate Debt' und 'Infrastructure Debt'  
   sind _nur_ bei den fett markierten Fragen unterschiedlich zu 'Direct Lending'?  
   Oder gibt's noch andere Unterschiede?


## RC4 - Punkte

* Fragengruppe 3. Markt Sentiment        - für alle asset classes gemeinsam? Nein
* Fragengruppe 4. Qualitative Questions  - für alle asset classes gemeinsam? Nein

* Fully dynamic - asset classes + strategies

* Fully dynamic - zero transactions

* Fully dynamic - new transaction page skipped

* Mobile Layout

* globale Checkbox: Neue Transaktionen unter 1.1

* 1.1a.) Drei ESG Fragen als drei Eingabefelder?  
        1.1a       Total number

           -   I floating rate
           -  II ESG + Doc 
           - III Ratchets  

           -   I Low mid cap
           -  II Core+Upper 
           - III Other  

    Summation-Prüfung dann über jeweils römisch I-II-III muss total number of deals ergeben?

Prio b - Number Input - replace Komma mit Punkt; wenn möglich

b.) total summe
   d.) e.) f.) ein Feld mit live Summe der eingetragenen Werte

   Bei Seitenwechsel: "Summe von e.) ist ungleich b.) total

Prio
   11a.) a-I-III for each must lte 11a.)
   bei Seitenwechsel

Horizonatels Scrollen für Mobile

## Test-Links
* [Leerer Fragebogen   ](https://private-debt-survey.zew.de/a)
* [Vorausgefüllter Werte](https://private-debt-survey.zew.de/d/PDS--P3XDDGD4V)

### RC4 Release Notes

#### Inhatliche Anmerkungen

* Teil A3/B3/C3 - Sentiment bezieht sich auf den Gesamtmarkt und nicht nur die eigenen Transaktionen. Man sollte das in einem einleitenden Satz den Teilnehmern erklären.

* Teil A3/B3/C3 - Bitte prüfen Sie die bei den Fragen zu den kommenden 3 Monaten die Formulierung:  
"Bei den Fragen zu den kommenden 3 Monaten finde ich die Formulierung seltsam: Next 3 months How do you expect the pricing (taking into account margins over the relevant reference rate and other return components like fees) for new deals ??you observed?? in the market change in the coming quarter?"

* Teil A4/B4/C4 - Frage 4.3: Passt hier die Erklärung von "not relevant" zu "potential dealbreaker"? 

#### Technische Anmerkungen

* Request war: "Dezimal-Trennzeichen soll immer Komma sein. Eingabe von Punkt soll möglich sein, wird aber in Komma verwandelt."  
  Firefox wurde auf Komma umgestellt.  
  Chrome und Safari lassen sich nicht zwingen, ein Komma anzuzeigen, wenn das Betriebssystem bspw. Englisch ist.  
  Das automatische Ersetzen von Punkt durch Komma würde diese Situation noch verschlechtern.
  Hinweis: In der Datenbank und im Datenexport erscheint einheitlich immer ein Punkt als Dezimaltrenner.

* Pages A1,B1,C1,D1 - nachfolgende Page - "Unlevered returns..." wird nicht angezeigt,  
    wenn "Total number of transactions" für _alle_ Tranchentypen "0" ist.  
    Wir hatten hier eine extra Checkbox auf Pages A1,B1,C1,D1 (über "Total number of transactions") besprochen;
    aber die ist jetzt vielleicht nicht nötig?

* Pages A1,B1,C1,D1 - Number of deals == 0  
     => ganze Spalte disabled  

* Pages A1,B1,C1,D1 - Number of deals  
     => Summanden _kleiner_-gleich  Summe  

* Pages A1,B1,C1,D1 - "Total transaction volume "  
     => Wert wird kopiert nach d.), e.) und f.)  

* "QQ YYYY" ist immer das Vorquartal.  
     Für Q2-2023 wäre es Q1-2023...

* Die meisten Slider haben im letzten Excel-Dokument neue Sonderwerte bekommen; bspw. <2% und  >20%  
   * Es wurde entsprechende Sonder-Programmierung hinzugefügt.
   * Im Ergebnis-Export werden nur die "Rohwerte" gespeichert sein;  
      eine doppelte Speicherung Rohwert _und_ angezeigter Wert (-1 => "<2%" oder 2 => "2-2.5 mn €") erzeugt zu viele riskante Sonderfälle. Man muss sich einmalig ein Mapping der Rohwerte zu den Displaywerten bspw. als Excel-Makro anlegen.

* Anzeige-Version für mobile phones wurde fertiggestellt


### PCAG Todo

* Excel spec - Real estate debt - 1.3d bis 1.3j - ranges are missing

* Excel spec - Infrastruct debt - 1.3d bis 1.3f - ranges are missing

<!-- * Excel - Real estate debt - question 2d.) - Multiple on Invested Capital - should really be omitted?  -->


<!-- 
https://localhost:8083/survey/d/PDS--P3XDDGD4V
-->


### RC5 Release Notes

* Todo: Imprint - DSGVO, Kontakt
* Ergebnisdownload - Frank Brückbauer
* Spalte deaktivieren - auf den fortfolgenden Seiten

* Frage 4.3: Wirklich `core principal`? Oder ist `core principle` gemeint?



Telefonisch/Zoom besprechen:

#### Komma-Eingabe

* Wie soll das im englischen/angelsächsichen Chrome aussehen?

* Slider display: Immer Komma?

#### Slider Rohwerte - Mapping

Im Ergebnis-Export werden nur die "Rohwerte" gespeichert sein;
eine doppelte Speicherung Rohwert und angezeigter Wert (-1 => "<2%" oder 2 => "2-2.5 mn €") erzeugt zu viele riskante Sonderfälle. Man muss sich einmalig ein Mapping der Rohwerte zu den Displaywerten bspw. als Excel-Makro anlegen.

Zum Verständnis: bedeutet das, wenn ich beispielsweise bei Frage 1.2 a) den Schiebe auf die erste Position setze, dann wird eine 1 in die Datenbank geschrieben? Da wir ja nicht notwendiger weise alle Mikrodaten erhalten, ist es möglich das benötigte Mapping dann im Auswertungs/Aggregationsschritt durchzuführen, bevor sie uns die Daten zuschicken?

#### IPhone Bugfix

Ich hatte das Problem mit IPhone und Safari als letztes in RC4 korrigiert;  
die entsprechenden Hilfsdateien werden vom Smartphone/Browser teilweise gecached, obwohl ich dem Browser eine neue Fassung anzeige...  
Bitte versuchen Sie es nochmal.

#### Progressbar/Navigation

Der PCAG Vorschlag mit einer zweigliedrigen Navigation (Zeile1: Assetklasse, Zeile2: Frageblöcke X1,X2 ... X4) erscheint mir elegant. Die Logik-Programmierung (mit zwei verschachtelten Zeilen/Listen) würde ich kurzfristig riskieren. Aber die Ausarbeitung des Browserlayouts (in HTML-Sprache) ist so kurz vor Liveschaltung zu unsicher. Ich habe einen hoffentlich passablen Umsetzungsvorschlag innerhalb der bestehenden Technik eingebaut. Bitte prüfen. Wenn das unzulänglich ist, dann können wir in _Q2_-2023 die vorgeschlagene Navigationsleiste einbauen.
