gh-ssh-pubkey
=============


Add SSH public keys into authorized_keys from GitHub.


Installation
--------------

```sh
go get github.com/pocke/gh-ssh-pubkey
```

Or download a binary from Latest release.

https://github.com/pocke/gh-ssh-pubkey/releases/latest


Usage
-------


Get pocke's public keys.

```sh
$ gh-ssh-pubkey pocke  # display keys
```


Get for more users.

```sh
$ gh-ssh-pubkey pocke users2 user3 ...
```

Get and Write into ~/.ssh/authorized_keys

```sh
$ gh-ssh-pubkey pocke -w
# or
$ gh-ssh-pubkey pocke --write
```
