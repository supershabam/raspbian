# raspbian
scripts to customize a raspian image

After running `go main.go` in the root of this project, the file `image.img` will be a burnable raspbian image that can run on a raspberry pi 3.

## features

* automatically connects to WPA networks listed in vars.yml
* starts ssh server on boot
* any client ssh key signed with the CA private key can ssh as root to the pi
* keyboard set to US
* WPA set to US
* HDMI set to safe mode for my small 3.5-inch display

## todo

[ ] download image and verify sha256

[ ] vendor golang dependencies
