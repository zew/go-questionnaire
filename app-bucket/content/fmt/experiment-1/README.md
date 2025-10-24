# Protokoll Meeting 2025-10-24

* Change request procedure
    * git clone https://github.com/zew/go-questionnaire
    * Changes to   \app-bucket\content\fmt\prototype\

    * Oder Screenshots


* Test - über Teilnehmer-Login-URLs  - nur bis Umfragestart
    *    \app-bucket\content\fmt\prototype\tmp-treatment.xlsx


* Participants without previous data?
    * exclude; beschlossen


* Participants English (~4)
    * exclude - otherwise lots of work to make everything English
    * erstmal deutsche Fassung perfektionieren
    * ab kommenden Mittwoch

* History funktion - Fields, Pages
    * Demo der History-Funktion
    * Todo pbu: Einbau in Seiten pg1, pg3, pg4, pg5
    * Davud: History auch für Seite 6 - Pro­gno­se Quartal
        Q4 2025, Q1 2026, Q2 2026, Q3 2026
 
 
* Who gets treatment?
    * ForecastData(participantId)["group"] == "T"

    * Todo pbu: Treatment - sieht auch den Text nicht
    * Todo pbu: Chart2, 3 - Quarter ist natürlich das geschätzte Quarter - fix 2025-Q4


* Rounding:  Coarser numbers - if percentage > 10

* Changes to the Echarts b and c:  
    * width maximized,  - label inset
    * we need the unit somewhere (%) 


* Dynamic numbering the charts 
    * Die Charts sind ja keine "Fragen" - 
      zur Vermeidung von Irritationen: Eigener Nummerkreis "Abbildung 1-3"
    * Dann auch keine Irritation, wenn nach 4a-c. dann folgt 5a, 5b, !6, !7  
    * Todo pbu: 5a, 6, ... einen Zähler runter, besser Slider Frage: mit 4 nummerieren


* Number of quarters ahead:  
    Question 3a - four quarter, 5a/b - three quarters




