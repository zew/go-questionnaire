# GoQuestionaire

A http(s) webserver serving a questionaire.

## Design principles


* All content and all results are driven  
by a __single JSON file__ .


* No database, but JSON and CSV result files.


* Transmission of the results is accomplished by _another_ component.  
The Transmitter. 
