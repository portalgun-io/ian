Ian
===

.. image:: https://travis-ci.org/thylong/ian.svg?branch=master
    :target: https://travis-ci.org/thylong/ian

.. image:: https://goreportcard.com/badge/github.com/thylong/ian
    :target: https://goreportcard.com/report/github.com/thylong/ian

.. image:: https://godoc.org/github.com/thylong/ian?status.svg
    :target: https://godoc.org/github.com/thylong/ian

.. image:: https://beerpay.io/thylong/ian/badge.svg?style=flat
    :target: https://beerpay.io/thylong/ian

Ian is a simple CLI tool to make your development environment portable.


Installing
==========

Ian comes as a binary. Once in your $PATH, you're good to go !

Linux
--------

.. code-block:: console

    $ go get github.com/thylong/ian


Mac OS X
--------

Ian requires Homebrew_.

.. code-block:: console

    $ go get -u github.com/thylong/ian


Usage
=====

.. code-block:: console

    $ ian
    Ian is a simple tool to manage your development environment and repositories.

    Learn more about Ian at http://goian.io

    Usage:
      ian [command]

    Default Commands:
      env         Manage development environment
      repo        Manage local repositories
      setup       Set up ian configuration

    Use "ian [command] --help" for more information about a command.

Features
========

- Manage development environment (update OS package manager packages, etc).
- Self-Setup (import dotfiles and install related packages)
- Support pip_, npm_, Homebrew_, RubyGems_, Cask_, apt_, yum_


Documentation
=============

Documentation can be seen here_. It was built thanks to the awesome Hugo project.
If you want to check the docs locally or to contribute to it, you can install hugo
and serve the static website locally using these commands:

.. code-block:: console

    $ brew update && brew install hugo
    $ cd docs/
    $ hugo server -t hugo-theme-learn --buildDrafts



Contributing
============

- Fork it
- Create your feature branch (git checkout -b my-new-feature)
- Commit your changes (git commit -am 'Add some feature')
- Push to the branch (git push origin my-new-feature)
- Create new Pull Request

Special thanks
==============
- `Devin Wilmot`_
- `Trisha Batchoo`_

.. _`template`: https://github.com/thylong/ian/blob/master/config/config_example.yml
.. _Homebrew: http://brew.sh
.. _here: https://goian.io
.. _Cask: https://caskroom.github.io
.. _RubyGems: https://rubygems.org/
.. _pip: https://packaging.python.org/
.. _npm: https://www.npmjs.com/
.. _apt: https://wiki.debian.org/Apt
.. _yum: https://fedoraproject.org/wiki/Yum
.. _`Devin Wilmot`: mailto:devwilmot@gmail.com
.. _`Trisha Batchoo`: https://github.com/tbat
.. _`Carl Chinatomby`: https://github.com/Carl-Chinatomby

## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/thylong/ian/badge.svg?style=beer-square)](https://beerpay.io/thylong/ian)  [![Beerpay](https://beerpay.io/thylong/ian/make-wish.svg?style=flat-square)](https://beerpay.io/thylong/ian?focus=wish)
