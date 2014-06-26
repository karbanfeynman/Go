## Web Related APP in Go
The repository includes two kinds of web APP written in Go. 

Web crawler - Use [fetchbot] and [goquery]. The content of source code is based on the example of fetchbot. According to the example, the app only extracts one level urls.

Web Service - Use [go-martini] and the extension "render" to build up a light-weight web service. Basically, I follow the lessons on [gophercasts]. So far, the web service can read data from database, and mix the template to produce a table webpage.

Others - Use SQLite as I/O for two apps.


## Changelog
* 2014/06/26: Web crawler can read seed urls from database and save the urls fetched from webpages into database.
* 2014/07/01
* 


==
## License
[The BSD 3-Clause License][bsd]

[fetchbot]: https://github.com/PuerkitoBio/fetchbot
[goquery]: https://github.com/PuerkitoBio/goquery
[go-martini]: https://github.com/go-martini/martini
[gophercasts]: https://gophercasts.io/lessons
[bsd]: http://opensource.org/licenses/BSD-3-Clause
