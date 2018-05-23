 [![GoDoc](http://godoc.org/github.com/zew/go-questionaire?status.svg)](http://godoc.org/github.com/zew/go-questionaire) 

# Go-Questionaire 

A http(s) webserver serving a questionaire.

## Status 

Under development. Unready.

## Design principles


* All content and all results are driven  
by a __single JSON file__ .

* No database, but JSON result files.

* Transfer of the results is accomplished by _another_ component.  
The Transferrer. 



* Validation happens on server side

* Client side JS validation is deliberately omitted

* [A simple client lib](http://www.javascript-coder.com/html-form/form-validation.phtml)



## Todo

* Tests

* JSON schema validator



## About Go-App-Tpl

* Go-Questionaire is based on Go-App-Tpl

* Go-App-Tpl is a template for a go web app.  

It features

  * Http router with safe settings and optional https encryption

  * Session package by Alex Edwards

  * Configurable url prefix allows running multiple instances on same server:port

  * Middleware for logging, access restrictions etc.

  * Middleware blocking request handler panics from taking down the server

  * Static file handlers
  
  * Markdown file handler rewriting image links and serving global README

  * Layout template with jQuery from CDN cache; fallback to localhost 

  * Configurable compilation of templates

  * Dynamic subtemplate calls 

  * Templates having access to session and request

  * JSON config file with reloadable app settings 

  * JSON file with reloadable logins

  * Shell script to control your go server under Linux


## Design guidelines

* Subpackaging is done by concern, neither too amorphous nor too atomic. 

* Go-App-Tpl has no "hooks" or interfaces for perfect isolation of "framework" code
and "custom handlers".  
Just copy it and add your handlers. Future updates can be merged.

