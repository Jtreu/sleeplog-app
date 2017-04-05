<h1 align="center">
  <br>
  <a href="https://www.sleeplog-app.com"><img src="https://www.sleeplog-app.com/static/logo.png" alt="WebTorrent" width="200"></a>
  <br>
  Sleeplog-app
  <br>
  <br>
</h1>

> An activity logger for neurologists and patients

## Services

* Firebase Hosting - https://www.sleeplog-app.com host
* Heroku - https://api.sleeplog-app.com host
* Google Domains - Domain provider
* Mailgun - Email service
* MLab - MongoDB provider

## Features

* Manage entries
  - Add sleep record
  - Add beverage record
  - Change day type
* Register for an account
* Login to account
* Search for other users
* Email confirmation
* Password reset email

## TODO

* Update profile info (icon, banner, username, email, password, etc.)
* Add user roles
  - Neurologist role asks patients for permissions to view entries
  - Patient role receives permission requests to view entries
  - Allow other users to request permission to view other useres entries
* Add service worker for offline support
* Add a indroductory sequence for first-time users
* Make the times entries responsive & scrollable
* Change the font when the user enters something into the table
* Create a homepage with information about the application
* Make the select boxes disappear and be replaced with static information when a user is viewing a table that is not theirs.
* Change your profile picture using the imageSourceManipulator.js util
* Have top left header element link to homepage
* Remove dropdown boxes from pages after searching for users

## Technologies

* [Vue.js 2.x](https://github.com/vuejs/vue)
* [axios](https://github.com/mzabriskie/axios)
  - http request library
* [localforage](https://github.com/localForage/localForage)
  - local storage library
* [KeenUI](https://github.com/JosephusPaye/Keen-UI)
  - material design library
* [Certbot](https://certbot.eff.org/)
  - ssl certificate generation tool

* [Golang 1.7.x](https://github.com/golang/go)
* [jwt-go](github.com/dgrijalva/jwt-go)
  - go implementation of json web tokens
* [negroni](github.com/urfave/negroni)
  - go http middleware utility

[![js-standard-style](https://cdn.rawgit.com/feross/standard/master/badge.svg)](https://github.com/feross/standard)
