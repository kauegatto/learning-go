http.HandlerFunc
Earlier we explored that the Handler interface is what we need to implement in order to make a server. Typically we do that by creating a struct and make it implement the interface by implementing its own ServeHTTP method. However the use-case for structs is for holding data but currently we have no state, so it doesn't feel right to be creating one.

HandlerFunc lets us avoid this.

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

```
type HandlerFunc func(ResponseWriter, *Request)
```

From the documentation, we see that type HandlerFunc has already implemented the ServeHTTP method. By type casting our PlayerServer function with it, we have now implemented the required Handler.

# ListenAndServe
http.ListenAndServe(":5000"...)
ListenAndServe takes a port to listen on a Handler. If there is a problem the web server will return an error, an example of that might be the port already being listened to. For that reason we wrap the call in log.Fatal to log the error to the user.

What we're going to do now is write another test to force us into making a positive change to try and move away from the hard-coded value.
