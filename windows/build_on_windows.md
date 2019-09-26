install pkg-config with choco

get libusb win binaries from libusb site
copy dll from MS64 folder into  C:\Windows\System32 or project folder or ou will get exitcode 0xc0000135

configure env var PKG_CONFIG_PATH to be where libusb-1.0.pc file is
edit  libusb-1.0.pc to point to libusb location
run `choco install mingw` because CGO needs gcc