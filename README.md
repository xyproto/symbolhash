symbolhash
==========

Description
-----------

Given a string, returns a unicode hash of the desired length.

The entropy is reflected by how much repetition there is in the resulting hash.

One possible use, if you administer a user database for instance, is to visually examine how similar password hashes are without viewing the actual password hashes.

Example
-------

Example input and output:

```"Several different words lalalalala"```, ```16``` -> ```♥☻☘⇧❤❄תּ☂☢♞✂♻⚡▷☮♥```

```"hi"```, ```16``` -> ```☭☢☭☢☭☢☭☢☭☢☭☢☭☢☭☢```

License and author
------------------

MIT license

Alexander Rødseth, 2013
