# webami
A whoami for the internet. Easily retrieve your public IP address from the command-line. webami utilises [Ipify](https://www.ipify.org/) to retrieve your public IP address and return it within the command-line.

webami is designed to be a simple way to avoid context-switching (we :heart: terminal!) and provide a simple way to get your public IP should you need to pipe it into things.

[![Webami Build](https://github.com/jonbayl/webami/actions/workflows/webami.yml/badge.svg)][![CodeQL](https://github.com/jonbayl/webami/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/jonbayl/webami/actions/workflows/codeql-analysis.yml)

# Installation
webami is very easy to install:
``` 
cd webami
make
```

Move the resultant "webami" executable to your `PATH` for easy usage.

Alternatively, you can grab the latest release from the current [releases](https://github.com/jonbayl/webami/releases). Release executables are made available for both Linux and Windows.

# Usage
There really isn't much to using webami! Once it's compiled, simply use the executable:
`./webami`

You can specify a different IP retrieval service if you wish. You can do this using the use command: 

```
./webami use <ip retrieval service URL>
```
e.g:
```
./webami use https://api.ipify.org
```

You are free to use whatever IP retrieval service you like, for example your own self-hosted version of Ipify or an alternative service. However, webami expects the output of the service used to follow some simple rules:
- The service must be accessible using either the HTTP or HTTPS protocol.
- The service must present the Content-Type: plain/text.
- The IP address must be the only thing returned in the response body.

To find out the current version of your webami executable, you can use:
```
./webami version
```

More features will be added soon and will be documented. However, for a quick assist, the help functionality is an easy way to get it:
```
./webami help
```

# The TODO List...
The current feature set is minimal and whilst it's enough to get the job done (get your public IP), it can definitely be improved. Future features on the TO DO list include:

- ~The ability to specify additional sources to retrieve your Public IP address. This could include a self-hosted version of Ipify.~
- ~Addition of simple build and tests.~

# Credit where credit is due
webami uses [Ipify](https://www.ipify.org/) to retrieve your public IP address. It is essentially a command-line wrap around the Ipify API. Ipify is a seperate, unrelated project (maintained by Randall Degges) which has it's source code in a [Github Repository](https://github.com/rdegges/ipify-api). 

# License
![license](https://img.shields.io/github/license/jonbayl/webami)
