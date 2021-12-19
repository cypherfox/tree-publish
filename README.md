# tree-publish
publish a tree-like data structure from a backend to a front-end.

It needs to be a tree in order to publish the data as JSON document. If you need more complex graph structures you need to compute the needed references before handing the data over to the library.

The backend is written in golang, the frontend in javascript.

# Aims
This library was born out of the need to publish a data structure from a golang backend to the frontend running in a browser. I could not find a suitable solution in a few moments of googling that would cover the following requirements:

* only publishing to the frontend needed. The backend is the source of truth and requests for changes are handled through other means.
* minimalistic, i.e. not part of a large framework. Ideally as little code as possible.

# Usage

On the backend side:
```
func Publish(document interface{})
```

once the front end calls `update()` the next time, the local copy of the document will be updated.

