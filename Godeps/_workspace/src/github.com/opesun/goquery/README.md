goquery
=======

Jquery style selector engine for HTML documents, in Go.

Future
======
If the package sees some usage then it will get a more comprehensive API.

Example
=======
See "remote.go" in the examples folder.

```
package main

import(
	"github.com/opesun/goquery"
)

func main() {
	x, err := goquery.ParseUrl("http://www.youtube.com/watch?v=ob_nh1WMMzU")
	if err != nil {
		panic(err)
	}
	x.Find("#eow-title").Print()
}
```
This will output (if it can load the url):

```
 span

    a
       Bounty Killer
     - Look
```