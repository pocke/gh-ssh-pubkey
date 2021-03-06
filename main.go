package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/ogier/pflag"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	var write bool
	pflag.BoolVarP(&write, "write", "w", false, "write to ~/.ssh/authorized_keys")
	pflag.Parse()

	if len(pflag.Args()) == 0 {
		return fmt.Errorf("Usage: gh-ssh-pubkey [-w] username [, username2...]")
	}

	out, err := Out(write)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, name := range pflag.Args() {
		keys, err := GetSSHKey(name)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "\n# from https://github.com/%s\n", name)
		fmt.Fprint(out, keys)
	}
	return nil
}

func GetSSHKey(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://github.com/%s.keys", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

func Out(write bool) (io.WriteCloser, error) {
	if !write {
		return os.Stdout, nil
	}

	path, err := homedir.Expand("~/.ssh/authorized_keys")
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
}
