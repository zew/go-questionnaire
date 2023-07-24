# Caroline Knebel Experiment

* Radio-Werte sind immer 1...5 oder 1...10 - `Weiß nicht` ist dann bspw. 9 und `keine Antwort` ist 10  
  Sonst Fehlergefahr

* Treatment is not yet coded

* Mobile phone:  
	Layout des Experiments für Smartphones würde nicht-triviiale Umordnung erforden.  
	a.) aufwändig
	b.) Verfälschungen zwischen Smartphone-Teilnehmern und PC-Teilnehmern durch Anordnung
	=> Idee: Keine Anpassung des Layouts für Smartphones  
	=> Der Panelprovider sollte erzwingen, dass die Teilnehmer über einen Desktop-Computer (PC,Apple) antworten  

Testing under <https://survey2.zew.de/survey/?u=9990&sid=kneb1&wid=2022-08&p=1&h=9DBZ8ko6n9vv4elnU5cg4duPQFQFD36lx7VVnNjQCno#>

Test URLs

pbu
* <https://survey2.zew.de?u=9990&sid=kneb1&wid=2023-06&p=1&h=DNWNoIHHR8RIUiOq9quWU1V1TgHt2tLtv1xaP8nVjyk>
* <https://survey2.zew.de?u=9991&sid=kneb1&wid=2023-06&p=1&h=MY1DwxojlLBrci7k8XubmI9aW4IPrTFj6AOiAsO5xMs>

Knebel
* <https://survey2.zew.de?u=9992&sid=kneb1&wid=2023-06&p=1&h=S1LMUHYI8fpY2fHsaUsjiCHdSyckr7vsDnzcUnZyc0c>
* <https://survey2.zew.de?u=9993&sid=kneb1&wid=2023-06&p=1&h=06OApCzf6-CuvXgyXTDRfDZfJNOZ7TnSmpVSyctmtbM>

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


## MSCI world return and sigma

<https://investingintheweb.com/blog/msci-world-index-historical-data/>

In Dollar, in the last 10 years (to April 30th, 2023), the MSCI World Index has had an annualised return of 9.29% 
  (MSCI World Total Net Return).
	$100 invested on April 30th, 2013, would be worth $243 on April 30th, 2023.
	1,0929^10=2,43

In Dollar, in the last 37 years to 1987, the MSCI World Index has had an annualised return of 8.06% 
  (MSCI World Total Net Return). 
	1,0806^37=17,6

Annualised standard deviation of 14.62%
Inflation rate for the US has been roughly 2.60% during the same period
Sharpe ratio of 0.61 - Überrendite bei vergleichbarem Risiko

In Euro, for last 10 years to 2013, the MSCI World Index in EUR has had an annualised return of 9.09%.
	1,0909^10=2,39
In Euro, for last 25 years to 1998, the MSCI World Index in EUR has had an annualised return of 5.90%.
	1,059^25=4,19

90 percent confidene interval is 1.645 standard deviations from the mean; 
<https://i1.wp.com/makemeanalyst.com/wp-content/uploads/2017/05/Confidence-Coefficients-for-90-Confidence.png?resize=760%2C492>



<https://www.companisto.com/en/academy/recht-steuern-und-hilfsthemen/bruttorendite-vs-nettorendite-wie-wird-der-ertrag-richtig-gemessen>
gross return
 % taxation, inflation, costs (management fees)
net return or real return or Total Net Return 

