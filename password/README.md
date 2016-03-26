# Passphrase creation

The main idea for this repo is to create an algorithm to generate safe passphrases.

The user must choose a personal phrase such as: **MyNameIsJohnAndIWasBornOn1968** and the desired length for the passphrase.

The personal information, being:

- **John**
- **1968**

is then used to generate the passphrase. The algorithm uses `math.rand` to have something different each time the program is run.

The algorithm is based on the [Blowfish algorithm](http://www.usenix.org/event/usenix99/provos/provos.pdf) and is used at this other [repo](github.com/LeonardoCastro/myapp).