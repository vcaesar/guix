GUIX - A Go cross platform UI library.
=======

[![CircleCI Status](https://circleci.com/gh/vcaesar/guix.svg?style=shield)](https://circleci.com/gh/vcaesar/guix)
![Appveyor](https://ci.appveyor.com/api/projects/status/github/vcaesar/guix?branch=master&svg=true)
<!-- [![codecov](https://codecov.io/gh/vcaesar/guix/branch/master/graph/badge.svg)](https://codecov.io/gh/vcaesar/guix) -->
[![Build Status](https://travis-ci.org/vcaesar/guix.svg)](https://travis-ci.org/vcaesar/guix)
[![Go Report Card](https://goreportcard.com/badge/github.com/vcaesar/guix)](https://goreportcard.com/report/github.com/vcaesar/guix)
[![GoDoc](https://godoc.org/github.com/vcaesar/guix?status.svg)](https://godoc.org/github.com/vcaesar/guix)
[![Release](https://github-release-version.herokuapp.com/github/vcaesar/guix/release.svg?style=flat)](https://github.com/vcaesar/guix/releases/latest)
[![Join the chat at https://gitter.im/go-ego/ego](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-ego/ego?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Dependencies
---

### Linux:

In order to build GUIX on linux, you will need the following packages installed:

    sudo apt-get install libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev mesa-common-dev libgl1-mesa-dev libxxf86vm-dev

### Common:

After setting up ```GOPATH``` (see [Go documentation](https://golang.org/doc/code.html)), you can then fetch the GUIX library and its dependencies:

    go get -u github.com/vcaesar/guix/...

Samples
---
Samples can be found in [`guix/samples`](https://github.com/vcaesar/guix/tree/master/samples).

To build all samples run:

    go install github.com/vcaesar/guix/samples/...

And they will be built into ```GOPATH/bin```.

If you add ```GOPATH/bin``` to your PATH, you can simply type the name of a sample to run it. For example: ```image_viewer```.

Web
---

guix code is cross platform and can be compiled using GopherJS to JavaScript, allowing it to run in browsers with WebGL support. To do so, you'll need the [GopherJS compiler](https://github.com/gopherjs/gopherjs) and some additional dependencies:

    go get -u github.com/gopherjs/gopherjs
    go get -u -d -tags=js github.com/vcaesar/guix/...
    
Afterwards, you can try the samples by running `gopherjs serve` command and opening <http://localhost:8080/github.com/vcaesar/guix/samples/> in a browser.

Fonts
---
Many of the samples require a font to render text. The dark theme (and currently the only theme) uses `Roboto`.
This is built into the gxfont package.

Make sure to mention this font in any notices file distributed with your application.

Contributing
---
GUIX was written by a couple of Googlers as an experiment and is now unmaintained.

Contributions, however small, will require the author to have signed the [Google Individual Contributor License Agreement](https://developers.google.com/open-source/cla/individual?csw=1).

The CLA is necessary mainly because you own the copyright to your changes, even after your contribution becomes part of our codebase, so we need your permission to use and distribute your code. We also need to be sure of various other things—for instance that you'll tell us if you know that your code infringes on other people's patents. You don't have to sign the CLA until after you've submitted your code for review and a member has approved it, but you must do it before we can put your code into our codebase. Before you start working on a larger contribution, you should get in touch with us first through the issue tracker with your idea so that we can help out and possibly guide you. Coordinating up front makes it much easier to avoid frustration later on.

Notice:
---

Unfortunately due to a shortage of hours in a day, [GXUI](https://github.com/google/gxui) is no longer maintained, this is an alternative attempt.

Disclaimer
---

The code is mostly undocumented, and is certainly **not idiomatic Go**.

This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.
