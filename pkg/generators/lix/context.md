Hallo Peter, 
 
CSV Dokument mit einer Zeile pro Unternehmen haben, mit den vergebenen Punkten der Hauptkategorien und allen Nebenkategorien? Wir bekommen von Sandra die Adressen samt ein paar Merkmale (Branche, Mitarbeiteranzahl, ...) 

für jedes Unternehmen einen eigenen Link und eine ID oder eine Variable 


Wir würden die Umfrage gerne in der Woche nach Pfingsten versenden. 
 


Zweisprachigkeit

DSGVO Text?


## Änderungen am Prototypen

Der Prototpy von Karina ist ansprechend und intuitiv.  
Ich wollte ihn direkt in der Umfrage verwenden.

Hierzu habe ich ein paar Anpassungen am Umfrage-Server vorgenommen.
Und auch den Prototypen gründlich durchgearbeitet.
Insgesamt erscheint mir Aufwand und Ergebnis immer noch äußerst effizient.
Und es ist ein Workflow, bei dem die Wissenschaft direkt die Gestaltung festlegt.


Arbeiten am Prototyp:


* Google Schriften entfernt (Ladezeiten, Erreichbarkeit, Lizenz)

* Datenschutz und Impressumg Link und Textseite hinzugefügt

* Umfangreiches Prüfen der Programmierung

* Umfangreiches Testen, viele kleine Anpassungen im Programmiercode

* Code umformatiert für Lesbarkeit

* Code gegliedert für Lesbarkeit - vier Module - funcs reordered strictly by nesting

* CSV Download der Ergebnisse deaktiviert

* Tastatur: "TAB" Taste springt jetzt sinnvoll durch die Elemente

* Bei Tastatureingabe: Kein "Einrasten" +/-3 von Gleichbewertung

* Mausklick auf Gleichbewertungs-Strich wird nicht länger blockiert

* Anfasser jetzt horiz. sauber in der über der Mitte des Gleichbewertungs-Strich

* Werte werden bei jedem "Weiter" gespeichert

* Tastatur Seite  1: "Weiter" ist bereits fokussiert (just Enter)

* Tastatur Seiten 2 ff: ALT+N und ALT+P um schnell durchzublättern

* Layout - kleine vertikale Stauchungen, damit man möglichst nicht scrollen braucht

* Noch offen: Test auf IPhone und Android Geräten, Test auf Firefox Browser

* Test Links Karina

9970	https://survey2.zew.de?u=9970&sid=lix&wid=2026-06&p=1&h=_3d7goZXX8zHS1kO4I-aQtQncQ_AcnGvaL4iRw-VLiY
9971	https://survey2.zew.de?u=9971&sid=lix&wid=2026-06&p=1&h=zXwX4zWQa0kZcW_nUdJtV-AD6OdIrMwmiq0oiv1skrE
9972	https://survey2.zew.de?u=9972&sid=lix&wid=2026-06&p=1&h=Gq8jElH1C-pJRWwPOQt1epN1rEZhdEZcNdcEl56v5X0

* Test Links FHE

9973	https://survey2.zew.de?u=9973&sid=lix&wid=2026-06&p=1&h=Ih3m3yzHzSwWg8QzShguUtCCesLkyvTwO8jNdnjp9xI
9974	https://survey2.zew.de?u=9974&sid=lix&wid=2026-06&p=1&h=Hc1YHt2h690t8tKTHnTPjGiljBOHNBHJaIGT-d1F2Ac
9975	https://survey2.zew.de?u=9975&sid=lix&wid=2026-06&p=1&h=dXtH70Dpoe_AegD952m9CQRqnMoXz4triXQScZGQQPg

* Test Links N.N.

9976	https://survey2.zew.de?u=9976&sid=lix&wid=2026-06&p=1&h=YJMKCyAcdvQ4gMk9AbLZ-VSXl2EYM_r1br5noOP9hPo
9977	https://survey2.zew.de?u=9977&sid=lix&wid=2026-06&p=1&h=_RxzwMmICMjTx8KfWMWeZAMKJfh9y8pmSwa5Ib6d34Y
9978	https://survey2.zew.de?u=9978&sid=lix&wid=2026-06&p=1&h=z_rgPRnLLgqPr5IiZHYfgohkFq7Smw4ZF31IZAKuBX4
9979	https://survey2.zew.de?u=9979&sid=lix&wid=2026-06&p=1&h=ZR0vwbFYKJ4dC6scqPQjEK1lAbo-P2_auFDLJbdoZHo


* Email "unternehmensbefragung@zew.de" scheint noch nicht aktiv

* Sandra hat uns 1.4 Mio Datensätze übermittelt.
    * Email-System ist konfiguriert und kann loslegen
    * Wir können monatlich ca. 7000 E-Mails unentgeltlich versenden, 
    * Dazu würden auch Erinnerungen zählen. Sind Erinnerungen geplant?
    * Irgendeine Priorisierung ?


* Trotz umfangreicher Arbeit mit dem Prototypen:  
  * Wir verlassen uns auf eine Programmierung, die von Claude Code erzeugt wurde.
  * Die Qualität ist extrem hoch.
  * Aber einige Nuancen sind eben doch nicht perfekt
        * Die Teilüberschrift bei der Zusammenfassung lautete noch "Energiesystem und Energiepolitik"
        * Das Darstellen der Pie-Charts verließ sich darauf,  
          dass das äußere Layout der HTML-Seite nach 60 ms fertig berechnet ist.  
          Auf langsamen Geräte (oder auf Geräten mit "Sophos" Virenscannern) wäre das Layout zerschossen.
        * Probleme, die ich nicht gefunden habe

    => Nur 50 Unternehmen anschreiben - und eine Woche Karenzzeit 





