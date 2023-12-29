* Radio-Werte sind immer 1...5 oder 1...10 -   
  `Weiß nicht` ist dann bspw. 9 und `keine Antwort` ist 10  
  Sonst Fehlergefahr

* Set `ah=30` and `ahV=25` in the debugging console, to change the chart

* Die Berechnung erfolgt in `echart-config.mjs`  
  `pComputeData()...`

* Rendite und Standardabweichung der Aktien-Anlage:  
  Detaillierte Diskussion unten: [MSCI world returns and sigma]

* Which deflated returns should we use for the _safe_ asset (`mnbd1`)?  
  Currently its set to 1 percent.

* Bei Anlagehorizont von 10 Jahren dominieren die Einzahlungen.  
  Die relativ höhere Aktien-Rendite schlägt nach 30 oder 40 Jahren viel stärker durch.  
  Spaßeshalber kann `az = 50` gesetzt werden.

* Y-Max wird ab 40 TSD dynamisch erhöht - in Schritten von 40.000.  
  Vertikale Abstände einerseits nicht zu klein.  
  Andererseits sollen sie nicht nach oben ausbrechen.  
  `Anlagehorizont` nicht in die Mitte, sondern mehr nach rechts? Bei 80%? 


# MSCI world returns and sigma

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

