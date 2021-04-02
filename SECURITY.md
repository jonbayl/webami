# Security Policy
Security is important to us. webami is small, and there isn't much which can go wrong. However, that doesn't mean it can't and when it does, we take it seriously and address is as soon as is possible.

## Supported Versions

As webami is new, the only major version is supported. To ensure you are using the version with the latest security updates, make sure you have the latest minor release. As the project progresses, we plan to support webami on an n-1 basis, so that the two most recent releases receive security updates.

Currently supported versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.x     | :white_check_mark: |


## Reporting a Vulnerability

We hope you never have to report one, but if you do - please don't use a public GitHub issue. Drop us an e-mail here: webami.opensource@gmail.com. We'll do our best to get back to you as quickly as possible to confirm the issue, and will keep you updated until the issue is closed.

# A note on IP retrieval services
webami is essentially a client, which requests your IP address from a selected IP retrieval service over the internet (or your local network, providing you specify a local provider). We'll do what we can to make sure the webami client is secure and any vulnerabilities are addressed. However, you still need to take care when selecting an IP retrieval service. For example, there is no guarantee that a retrieval service doesn't log your IP address for analytics. Use webami with the same level of caution you would visiting these services through your web browser.

If you do not select your own IP retrieval service, by default, webami utilises [Ipify](https://www.ipify.org/). If you choose to use the defaults, you should read their webpage and ensure you are happy with the way the service operates.
