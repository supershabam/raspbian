package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/yaml.v2"
)

type config struct {
	ImageURL    string  `yaml:"image_url"`
	ImageSHA256 string  `yaml:"image_sha256"`
	Mounts      []mount `yaml:"mounts"`
}

type mount struct {
	Offset   uint64    `yaml:"offset"`
	Sh       []string  `yaml:"sh"`
	Networks []network `yaml:"networks"`
}

type network struct {
	SSID       string `yaml:"ssid"`
	Passphrase string `yaml:"passphrase"`
}

func main() {
	var c config
	b, err := ioutil.ReadFile("./vars.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	spew.Dump(c)
	b, err = exec.Command("docker", "build", "-t", "raspbian", ".").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for _, mount := range c.Mounts {
		for _, cmd := range mount.Sh {
			b, err = exec.Command("docker", "run", "--privileged", "-v", fmt.Sprintf("%s:/run", pwd), "raspbian", "/bin/bash", "-c", fmt.Sprintf(`mount -o loop,offset=%d /run/image.img /mnt && %s`, mount.Offset, cmd)).CombinedOutput()
			if err != nil {
				fmt.Printf("%s\n", b)
				panic(err)
			}
			fmt.Printf("%s\n", b)
		}
		if mount.Networks == nil {
			continue
		}
		for _, network := range mount.Networks {
			b, err = exec.Command("docker", "run", "--privileged", "-v", fmt.Sprintf("%s:/run", pwd), "raspbian", "/bin/bash", "-c", fmt.Sprintf(`mount -o loop,offset=%d /run/image.img /mnt && wpa_passphrase "%s" "%s" >> /mnt/etc/wpa_supplicant/wpa_supplicant.conf`, mount.Offset, network.SSID, network.Passphrase)).CombinedOutput()
			if err != nil {
				fmt.Printf("%s\n", b)
				panic(err)
			}
			fmt.Printf("%s\n", b)
		}
	}
}
