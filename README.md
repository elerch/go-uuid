# uuid [![Build Status](https://travis-ci.org/hashicorp/go-uuid.svg?branch=master)](https://travis-ci.org/hashicorp/go-uuid)

Generates RFC 1422 complaint UUID-format strings using high quality, purely random bytes. It can also parse UUID-format strings into their component bytes.

Note that the only difference between this library and Hashicorp's is the
RFC 1422 compliance status. Their library focuses on random strings only
for security purposes. See https://github.com/hashicorp/uuid/pull/5#issuecomment-195463301 for more detail.

Documentation
=============

The full documentation is available on [Godoc](http://godoc.org/github.com/hashicorp/go-uuid).
