package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gh-ssh-pubkey [username, ...]")
		os.Exit(1)
	}

	for _, name := range os.Args[1:] {
		keys, err := GetSSHKey(name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("# from https://github.com/%s\n", name)
		fmt.Println(keys)
	}
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
