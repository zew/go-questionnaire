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


## MCI world standard deviation - 2024-01

https://de.wikipedia.org/wiki/MSCI_World#Berechnung

https://de.wikipedia.org/wiki/Aktienindex#Performanceindex

https://en.wikipedia.org/wiki/Stock_market_index
https://de.wikipedia.org/wiki/Total-Return-Index
https://en.wikipedia.org/wiki/Standard_deviation

MSCI world total return index  sigma 
MSCI world total return index  Variance
MSCI world total return index  Standard deviation


<https://investingintheweb.com/blog/msci-world-index-historical-data/>

MSCI World Total Net Return - Standard deviation (yearly)

 annualised standard deviation of 14.99%
 
 
## Annualised standard deviation of an retaining stock fund

* <https://financetrain.com/calculate-annualized-standard-deviation>
* <https://quant.stackexchange.com/questions/42445/total-returns-from-adjusted-close-prices>

The annualized standard deviation of daily returns is calculated as follows:

Annualized Standard Deviation = Standard Deviation of Daily Returns * Square Root (250)

## Standard deviation of 20 draws - variance of the _product_

* https://math.stackexchange.com/questions/2935743/

* expection of the product of two random variables  
  E[XY]=E[X]⋅E[Y]

* variance  of the product of two random variables  
  Var[X]⋅Var[Y]+Var[Y](E[X])2+Var[X](E[Y])2 

* if Var[X]=Var[Y]=vr and E[X]=E[Y]=mn  
    vr*vr +   (vr*mn)^2 + (vr*mn)^2 
    vr^2  + 2*(vr*mn)^2 
    vr^2  + 2*vr^2 *mn^2
    vr^2(1 + 2*mn^2) 

    vr^3(1 + 2*mn^2 + 4*mn^4) 

## Standard deviation of 20 draws - simple formula

* <https://quant.stackexchange.com/questions/48914/calculating-annualized-standard-deviation-from-monthly-returns-and-the-differe>

* <https://financetrain.com/calculate-annualized-standard-deviation>

_Annualised_ standard deviation of 14.62%  equals  0.1462

=> 20 draws => 0,1462 * sqrt(20)
=> 20 draws => 0,1462 * 4,472136
=> 20 draws => 0,6538


