# sshdb — SSH Key Lifecycle Management

Store, organize, and export SSH keys in a local SQLite database.

```
sshdb import mykey -i ~/.ssh/
sshdb list
sshdb show mykey
sshdb export mykey -o ~/.ssh/
sshdb rename old new
sshdb delete mykey
```

## Commands

### List

```
$ sshdb list
Name           Type            Comment         Created              Fingerprint
───────────────────────────────────────────────────────────────────────────────
alice@work     ssh-ed25519     alice corp key  2026-07-05 02:53:31  SHA256:OMVrasDoW7ZPpiHPEr0LQeI9bjXMuGKVWVPQxMvAPCw
```

### Show

Print the private and public key contents for a given name:

```bash
sshdb show mykey
```

### Import

Looks for `<name>` and `<name>.pub` in the current directory by default. Use `-i` to set the input directory, `--priv` / `--pub` for custom paths:

```bash
sshdb import mykey -i ~/.ssh/
sshdb import mykey --priv ~/.ssh/id_ed25519 --pub ~/.ssh/id_ed25519.pub
```

### Export

Writes `<name>` and `<name>.pub` in the current directory by default. Use `-o` to set the output directory, `--priv` / `--pub` for custom paths:

```bash
sshdb export mykey -o ~/.ssh/
sshdb export mykey --priv ~/.ssh/mykey --pub ~/.ssh/mykey.pub
```

## Storage

All keys live in `~/.ssh/sshgen.db`. No file I/O unless you explicitly export or import.

## Build from source

```bash
git clone https://github.com/gralliry/sshdb.git && cd sshdb
go build -o sshdb .
```

## License

MIT
