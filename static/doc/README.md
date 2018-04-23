# GoQuestionaire

A http(s) webserver serving a questionaire.

## Status 

Under development. Unready.

## Design principles


* All content and all results are driven  
by a __single JSON file__ .


* No database, but JSON and CSV result files.


* Transmission of the results is accomplished by _another_ component.  
The Transmitter. 


## Todo

The config should be reloadable.
This means that every access to it runs over a global lock.