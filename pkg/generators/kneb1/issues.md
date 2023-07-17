# Caroline Knebel Experiment

* Treatment is not yet coded

Testing under <https://survey2.zew.de/survey/?u=9990&sid=kneb1&wid=2022-08&p=1&h=9DBZ8ko6n9vv4elnU5cg4duPQFQFD36lx7VVnNjQCno#>

STRG+F5 - zum Neuladen der Seite

## Anmerkungen

* Die Einstellung des "Slider" ist nocht nicht mit der Berechnung des Portfolio-Mix verbunden. Man muss noch testen, ob er unter IPhone funktioniert.

* Nach Laden der Seite werden automatisch sechs Ziehungen ausgeführt.

* Danach kann man entweder mit `Next Step` eine einzelne weitere Ziehung durchführen,  
  oder per `Forever` auf Dauerfeuer schalten

* Der linke untere Chart ist relativ wertlos; er zeigt nur die Ziehungen der Höhe nach an. Der letzte/aktuellste Wert hat einen größeren Kreis.

* Man kann den linken unteren Chart weglassen, oder durch eine Animation ersetzen; Münzwurf oder Slotmaschine oder Würfel sind keine guten Metaphern, weil sie Gleichverteilung implizieren würden.  
Man könnte ein [Galton Board](https://www.youtube.com/watch?v=3m4bxse2JEQ) andeuten.

* Der rechte untere Chart zeigt das Histogramm. Bisher noch mit absoluten Häufigkeiten, damit ich eventuelle Fehler leichter sehen kann. Die grünen Balken sind das risky asset. Die Höhe der letzen Ziehung links korrespondiert mit dem gründen Balken, der wächst. Der ockerne/dunkelgelbe Balken ist das safe asset.

Es wäre schön, wenn die beiden _Flächen_ (gelb einerseits, Summe der grünen Balken andererseits) die Wahrnehmung der Auszahlungen nicht irritieren würden. Vielleicht muss man alle Balken extrem dünn machen?

Die Achsenbeschriftungen sind ebenfalls noch unschön.

Schließlich ist das Histogramm noch nicht exakt beim Mittelwert von 153.0 zentriert, so dass die grünen Balken auch bei großen Zahlen immer etwas skewed sind. Auch könnte man die Balken noch feiner machen (resolution 1, statt wie jetzt 5)...

