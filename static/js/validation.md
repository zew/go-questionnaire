# HTML5 Validation

* Error messages are displayed in the language, configured in the browser software.  
Chrome: Setting - Language - [List of installed languages] - Display Google Chrome in this language - [restart] will affect the messages

* The error messages are pretty mathematical anyway
 "Value must be greate or equal to ..."

* The method setCustomValidity('my message'), puts the input into invalid state, regardless of its content.

* We can disable the build-in bubbles in various ways.  
By far the best way is to set  
`<form novalidate=true>`.  
However we can still use validity and checkValidity().  
We could also call reportValidity(), to re-establish the previous behavior.

* We have the :valid and :invalid CSS pseudo classes;  
they trigger an instant valdiation error on every keystroke;  
they *can* be used for focus-dependent error feedback;
they cannot be used for form.submit() validation.

