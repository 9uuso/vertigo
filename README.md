vertigo [![wercker status](https://app.wercker.com/status/e1f07b85320f902313d32fec503c5017/s/master "wercker status")](https://app.wercker.com/project/bykey/e1f07b85320f902313d32fec503c5017) [![Gobuild Download](https://img.shields.io/badge/gobuild-download-green.svg?style=flat)](http://gobuild.io/github.com/9uuso/vertigo) [![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/9uuso/vertigo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
=======
[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/9uuso/vertigo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Note: Vertigo is **under development** and things may still break apart.

Vertigo is blogging platform similar to Wordpress, Ghost, Medium, Svbtle or Tumblr. You can make multiple accounts and write multiple posts, which then appear on your front page. You can also make unlisted pages similar to Wordpress, which you can link to navigation with help of HTML, similar to Tumblr.

Written Go, Vertigo is not only fast, but can be run on all major operating systems like Windows, Linux and MacOSX without the language development tools.

##Is Vertigo free?

It is. Not only is the source code available online for anyone to use on MIT license, you can also deploy Vertigo to Heroku free of charge. Just click the [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy) to get started. If you are worried that Heroku might own your content, you can also deploy Vertigo to a normal server or even run it privately on your own desktop. More at [Install instructions](https://github.com/9uuso/vertigo#install-instructions).

Note: To deploy to Heroku you need to have a credit card linked to them. If you wish not to link one, you may follow instructions on [here](https://github.com/9uuso/vertigo/issues/8) to remove Mailgun from Heroku add-ons list. If you remove Mailgun you cannot use password reminder, but everything else should work.

##How is Vertigo's frontend code handled?

The frontend code is powered by Go's `template/html` package, which syntax is similar to Mustache.js. The template files are in plain HMTL and JavaScript (vanilla) only appears on a few pages, such as the post edit page, where it is used to provide backup for any writings you make. JavaScript in general is aimed to be stripped down as much as possible to provide a better user experience on different devices. That being said, Vertigo works fine even with NoScript enabled.

Vertigo's routes by default can lead to either HTML templates or JSON endpoints depending on what URL is used. This means that as features as implemented, they are both available on /api/ and the normal frontend site. This makes it easy to add your preferred JavaScript MVC's on top of Vertigo. You can also code your own plugins or generate your own widgets from the data accessible from /api/ endpoint. This means that you create users, submit posts and read user data even without writing a single line of Go code. So basically, one could write a SPA application on top of the Go code with only using JavaScript. Whether you want to take that path or just edit the template files found in /templates/ is up to you.

##Demo

See [my personal website](http://www.juusohaavisto.com/)

##Install instructions

Note: By default the HTTP server starts on port 3000. This can changed by declaring PORT environment variable or by passing one with the binary execution command.

###Binaries

[![Gobuild Download](https://img.shields.io/badge/gobuild-download-green.svg?style=flat)](http://gobuild.io/github.com/9uuso/vertigo)

###Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

###Source

1. Install Go (I recommend using [gvm](https://github.com/moovweb/gvm))
2. `go get github.com/tools/godep`
3. `git clone github.com/9uuso/vertigo`
4. `cd vertigo && godep get ./ && godep go build`
5. Start Vertigo `PORT="80" MARTINI_ENV="production" ./vertigo`

###Environment variables
* PORT - the HTTP server port
* MARTINI_ENV - used by Martini to enable production optimizations such as template caching
* MAILGUN_API_KEY - Mailgun API key (declared by default with Heroku Mailgun Addon)
* MAILGUN_SMTP_LOGIN - Another Mailgun API key (declared by default with Heroku Mailgun Addon)

##Screenshots

![](http://i.imgur.com/EGlBhjP.png)
![](http://i.imgur.com/0AfvQnW.png)
![](http://i.imgur.com/AeC9xml.png)
![](http://i.imgur.com/rDlM9IX.png)
![](http://i.imgur.com/EwFcRfq.png)

##Features

- Installation wizard
- JSON API
- SQLite, MySQL and PostgreSQL support
- Search
- Multiple account support
- Auto-saving of posts to LocalStorage
- RSS and Atom feeds
- Password recovery with Mailgun
- [Partial](https://github.com/9uuso/vertigo/issues/7) Markdown support

##License

MIT
