# Password Checker

## Background

Password checker allows you to check your password against a dictionary file and set up a web app that allows others to check theirs. If a password is found then it wouldn't take long for someone to access your account using dictionary attack.

To use the app, you will need to download a password list / word dictionary. Then run the `splitter` program to split them into multiple files based on the password's first character. This would allow `finder` program to efficiently query the text file that has the password.

## Install

```bash
$ go get github.com/SystemR/password-checker/splitter
$ go get github.com/SystemR/password-checker/finder
```

## Usage

1.  Use the splitter to split the large password file. This would create pass folder:

```bash
# splitter <password file> <output folder>
$ $GOBIN/splitter ~/giant-password-file.txt ~/pass
```

2.  Check password with finder app (it's basically doing cat grep but exits as soon as 2nd char is different)

```bash
# Password file should match the 1st character of the password
# finder <split password file> <query>
$ $GOBIN/finder ~/pass/l.txt 'letmein'
```

3.  Or use with express webserver

```bash
# Run the express app to let users check their passwords
# node server.js <password folder>
$ node $GOPATH/src/github.com/systemr/password-checker/server/server.js ~/pass
```