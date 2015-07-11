vertigo
=======
[![Codeship Status for 9uuso/vertigo](https://codeship.com/projects/b2de9690-b16b-0132-08f1-3edef27c5b65/status?branch=master)](https://codeship.com/projects/69843) [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
[![Deploy vertigo via gitdeploy.io](https://img.shields.io/badge/gitdeploy.io-deploy%20vertigo/master-green.svg)](https://www.gitdeploy.io/deploy?repository=https%3A%2F%2Fgithub.com%2F9uuso%2Fvertigo.git)
![Vertigo](http://i.imgur.com/ZnAQR6I.gif)


Vertigo is blogging platform similar to [Ghost](https://ghost.org), [Medium](https://medium.com) or [Tumblr](https://www.tumblr.com). What makes Vertigo different is that it has JSON API for reading and writing data and it is written in Go. Therefore, Vertigo is not only fast, but can be run using single binary on all major operating systems like Windows, Linux and MacOSX without the language development tools.

The frontend code is powered by Go's `template/html` package, which syntax is similar to Mustache.js. The template files are in plain HTML and JavaScript (vanilla) only appears on few pages. JavaScript in general is aimed to be stripped down as much as possible to provide a better user experience on different devices. Vertigo also ships without any CSS frameworks, so it is easy to start customizing the frontend with the framework of your choice.

Thanks to the JSON API, it is easy to add your preferred JavaScript MVC on top of Vertigo. This means that you can create users, submit posts and read data without writing a single line of Go code. So basically, one could write a SPA application on top of Vertigo just by using JavaScript. Whether you want to take that path or just edit the HTML template files found in `/templates/` is up to you.

##Features

- Installation wizard
- JSON API
- SQLite and PostgreSQL support
- Search
- Multiple account support
- Auto-saving of posts to LocalStorage
- RSS and Atom feeds
- Password recovery with Mailgun
- Markdown support

##Demo

See [my personal website](http://www.juusohaavisto.com/)

##Installation

Note: By default the HTTP server starts on port 3000. This can changed by declaring `PORT` environment variable or by passing one with the binary execution command.

###Gitdeploy

Deploy and try out vertigo using gitdeploy:

[![Deploy vertigo via gitdeploy.io](https://img.shields.io/badge/gitdeploy.io-deploy%20vertigo/master-green.svg)](https://www.gitdeploy.io/deploy?repository=https%3A%2F%2Fgithub.com%2F9uuso%2Fvertigo.git)

###Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

For advanced usage, see [Advanced Heroku deployment](https://github.com/9uuso/vertigo/wiki/Advanced-Heroku-deployment)

Note: To deploy to Heroku you need to have a credit card linked to them. If you wish not to link one, you may follow instructions on [here](https://github.com/9uuso/vertigo/issues/8) to remove Mailgun from Heroku add-ons list. If you remove Mailgun you cannot use password reminder, but everything else should work.

###Source

1. Install Go (I recommend using [gvm](https://github.com/moovweb/gvm))
2. `go get github.com/tools/godep`
3. `git clone https://github.com/9uuso/vertigo`
4. `cd vertigo && godep get ./ && godep go build`
5. `PORT="80" MARTINI_ENV="production" ./vertigo`

###Docker
1. [Install docker](https://docs.docker.com/installation/)
2. `cd vertigo`
3. `docker build -t "vertigo" .`
4. `docker run -d -p 80:80 vertigo`

###Environment variables
* `PORT` - the HTTP server port
* `MARTINI_ENV` - used by Martini to enable production optimizations such as template caching
* `MAILGUN_API_KEY` - Mailgun API key (declared by default with Heroku Mailgun Addon)
* `MAILGUN_SMTP_LOGIN` - Another Mailgun API key (declared by default with Heroku Mailgun Addon)
* `DATABASE_URL` - Database connection URL for PostgreSQL - if empty, SQLite will be used

##Contribute

Contributions are welcome, but before creating a pull request, please run your code trough `go fmt` and [`golint`](https://github.com/golang/lint). If the changes introduce new features, please also add tests for them. Try to also squash your commits into one big one instead many small, to avoid unnecessary CI runs.

##Support

If you have any questions in mind, please file an issue.

##License

MIT
