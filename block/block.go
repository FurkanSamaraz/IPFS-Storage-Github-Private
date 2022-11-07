package block

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

func Blockencrypt(paths string, key string, zipname string) {
	fzip, err := os.Create(`./` + zipname + `.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	file, _ := os.Open(paths)
	defer file.Close()
	liste, _ := file.Readdirnames(0)
	for _, isim := range liste {

		dosya, _ := ioutil.ReadFile(paths + isim)
		w, _ := zipw.Encrypt(paths+isim, key)

		contents := []byte(dosya)
		_, err = io.Copy(w, bytes.NewReader(contents))
		if err != nil {
			log.Fatal(err)
		}
		zipw.Flush()

	}
}

/*

func Blockdecryption(paths string, password string) {
	r, err := zip.OpenReader(paths)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}

		r, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
	}
}

*/
