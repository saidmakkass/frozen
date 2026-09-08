# frozen

An IRC (Internet Relay Chat) server implementation in go.

## Goals
- Learn go.
- Get practical experience in implementing protocols.
- Have fun.
## Constraints
- Only standard library.
- Must use concurrency and goroutines.
## Implemented Features
- [ ] user sign-up (unique username thats immutable and password)
- [ ] user sign-in
- [ ] user nicknames (unique and mutable)
- [ ] user info is kept across reconnects of the user
- [ ] commands
    - [ ] PASS NICK USER - Initial authentication for a user.
    - [ ] NICK - Change nickname
    - [ ] JOIN - Makes the user join a channel. If the channel doesn’t exist, it will be created.
    - [ ] PART - Makes the user leave a channel.
    - [ ] NAMES - Lists all users connected to the server (bonus: make it RFC compliant with channel modes).
    - [ ] LIST - Lists all channels in the server (bonus: make it RFC compliant with channel modes).
    - [ ] PRIVMSG - Send a message to another user or a channel.

## Project Structure
```
└── frozen
    └── bin
    └── cmd
        └── frozen
            ├── main.go
    └── internal
    └── tests
    ├── go.mod
    ├── Makefile
    └── README.md
```