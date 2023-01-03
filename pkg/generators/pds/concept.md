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


## RC4

*  Excel file:  
   "Ist es möglich automatisch jeweils das relevante Quartal für "QQ YYYY" einzusetzen?"  
   Wie ist das gemeint? Letztes Quartal? Oder aktuelles Quartal?

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