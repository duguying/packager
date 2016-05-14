# packager [![Build Status](https://travis-ci.org/duguying/packager.svg?branch=master)](https://travis-ci.org/duguying/packager) [![Coverage Status](https://coveralls.io/repos/github/duguying/packager/badge.svg?branch=master)](https://coveralls.io/github/duguying/packager?branch=master)

packager is a tool for you to upload software archive from a ci.

## build

```shell
go get github.com/duguying/packager
```

## usage

```shell
packager -u=http://127.0.0.1/uploader.php -a=archive
```

then the `packager` will zip archive file or directory as archive.zip, and then it will upload the file archive.zip and send md5 sum code of archive.zip file form field via post request.

## license

MIT License
