symbolhash
==========

[![Build Status](https://travis-ci.org/xyproto/symbolhash.svg?branch=master)](https://travis-ci.org/xyproto/symbolhash)
[![GoDoc](https://godoc.org/github.com/xyproto/symbolhash?status.svg)](http://godoc.org/github.com/xyproto/symbolhash)


Online API Documentation
------------------------

[godoc.org](http://godoc.org/github.com/xyproto/symbolhash)

Features and limitations
------------------------

Given a string, returns a unicode hash of the desired length.

The entropy is reflected by how much repetition there is in the resulting hash.

One possible use, if you administer a user database or is creating a web application admin dashboard, is to visually examine how similar password hashes are without viewing the actual password hashes.

Example
-------

Example input and output:

```"Several different words lalalalala"```, ```16``` -> ```♥☻☘⇧❤❄תּ☂☢♞✂♻⚡▷☮♥```

```"hi"```, ```16``` -> ```☭☢☭☢☭☢☭☢☭☢☭☢☭☢☭☢```

General info
------------

* Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: MIT
* Version: 1.0.0
