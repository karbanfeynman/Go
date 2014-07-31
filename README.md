## Go-CAM
A remote controller of DSLR

This project combines with several components:
* Raspberry Pi which runs RASPIAN connects to DSLR by USB cable. 
* Use [gPhoto2] to control DSLR. 
* A light-weight web api written by [go-martini] runs on Raspbery Pi 
* combines [gPhoto2] and web-api. Therefore, DSLR can be controlled by the browser of any platform.
* Exif-Analyser can analysis images saved in the SD card of the camera connted to Raspberry Pi.



## Changelog
* 2014/07/27: The project starts.


==
## License
[The BSD 3-Clause License][bsd]

[gPhoto2]: http://gphoto.org/proj/
[go-martini]: https://github.com/go-martini/martini
[bsd]: http://opensource.org/licenses/BSD-3-Clause
