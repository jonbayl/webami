# webami
A whoami for the internet. Easily retrieve your public IP address from the command-line. webami utilises [Ipify](https://www.ipify.org/) to retrieve your public IP address and return it within the command-line.

# Installation
webami is very easy to install:
``` 
cd webami
make
```

Move the resultant "webami" executable to your `PATH` for easy usage.

# Usage
There really isn't much to using webami! Once it's compiled, simply use the executable:
`./webami`

If you're interested, you can get your current version of webami using: `./webami version`. 

More features will be added soon and will be documented here but to get the full webami help, use: `./webami help`.

# The TODO List...
The current feature set is minimal and whilst it's enough to get the job done (get your public IP), it can definitely be improved. Future features on the TO DO list include:

- The ability to specify additional sources to retrieve your Public IP address. This could include a self-hosted version of Ipify. 
- A set of expoted libraries to allow re-use with other Go programs.
- Addition of simple build and tests.

# Credit where credit is due
webami uses [Ipify](https://www.ipify.org/) to retrieve your public IP address. It is essentially a command-line wrap around the Ipify API. Ipify is a seperate, unrelated project (maintained by Randall Degges) which has it's source code in a [Github Repository](https://github.com/rdegges/ipify-api). 

# License
webami is licensed under a [MIT](https://choosealicense.com/licenses/mit/) license.