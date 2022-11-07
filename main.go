package main

import (
	"IPFS-Github-Storage-main/block"
	"IPFS-Github-Storage-main/pulls"

	"fmt"
	"log"

	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

var (
	path = "./pullipfs/"
)

func main() {
	var (
		i        int
		username string
		token    string
		url      string
		key      string
		zipname  string
	)
	os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	fmt.Print("1- public repo \n2- private repo\n")

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	switch {
	case i == 1:
		fmt.Print("Url: ")
		fmt.Scan(&url)
		fmt.Print("Key: ")
		fmt.Scan(&key)
		fmt.Print("Zip Name: ")
		fmt.Scan(&zipname)
		pulls.Pullrepo(url)
		encrypt(path, key, zipname)
		time.Sleep(time.Millisecond * 15)
		up := `./` + zipname + `.zip`
		ipfspull(up)
	case i == 2:
		fmt.Print("Note: Do not forget to create personal access tokens!\n")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Token: ")
		fmt.Scan(&token)
		fmt.Print("Url: ")
		fmt.Scan(&url)
		fmt.Print("Password Key: ")
		fmt.Scan(&key)
		fmt.Print("Zip Name: ")
		fmt.Scan(&zipname)
		pulls.Pulssrepo(url, username, token)
		encrypt(path, key, zipname)
		time.Sleep(time.Millisecond * 15)
		up := `./` + zipname + `.zip`
		ipfspull(up)
	/*case i == 2:
	fmt.Print("Hash: ")
	fmt.Scan(&hash)

	ipfsdown(hash, "./")
	time.Sleep(time.Millisecond * 15)
	*/

	default:
		fmt.Println("-----")
	}

}

/*
	func ipfsdown(hash string, paths string) {
		sh := shell.NewShell("localhost:5001")
		sh.Get(hash, paths)
	}
*/
func ipfspull(paths string) {
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.AddDir(paths)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("https://ipfs.io/ipfs/%s", cid)
}
func encrypt(paths string, key string, name string) {
	block.Blockencrypt(paths, key, name)
}

/*func decryption(name string, password string) {
	block.Blockdecryption(name, password)
}
*/
