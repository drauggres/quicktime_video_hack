[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CircleCI](https://circleci.com/gh/danielpaulus/quicktime_video_hack.svg?style=svg)](https://circleci.com/gh/danielpaulus/quicktime_video_hack)
[![codecov](https://codecov.io/gh/danielpaulus/quicktime_video_hack/branch/master/graph/badge.svg)](https://codecov.io/gh/danielpaulus/quicktime_video_hack)
[![Go Report](https://goreportcard.com/badge/github.com/danielpaulus/quicktime_video_hack)](https://goreportcard.com/report/github.com/danielpaulus/quicktime_video_hack)

## 1. What is this?
This is an Operating System indepedent implementation for Quicktime Screensharing for iOS devices :-)

[Check out my presentation](https://danielpaulus.github.io/quicktime_video_hack_presentation)

[See a demo on YouTube](https://youtu.be/8v5f_ybSjHk)

This repository contains all the code you will need to grab and record video and audio from your iPhone or iPad 
without needing one of these expensive MacOS X computers :-D
It probably does something similar to what `QuickTime` and `com.apple.cmio.iOSScreenCaptureAssistant` are doing on MacOS.
Currently you can use it to dump a h264 file and a wave file or mirror your device in a desktop window. Transcoding it to anything else is very easy since I used Gstreamer to render the AV data. 

## 2. Installation

1. Install Gstreamer and LibUSB with brew or apt depending on your platform.
3. Download the latest release and run it

### MAC OS X LIBUSB -- IMPORTANT
1. Make sure to use either this fork `https://github.com/GroundControl-Solutions/libusb`
   or a LibUsb version BELOW 1.0.20 or iOS devices won't be found on Mac OS X.
   [See Github Issue](https://github.com/libusb/libusb/issues/290)


## 3. Usage
```
Q.uickTime V.ideo H.ack (qvh) v0.2-beta

Usage:
  qvh devices [-v]
  qvh activate [--udid=<udid>] [-v]
  qvh record <h264file> <wavfile> [-v] [--udid=<udid>]
  qvh gstreamer [-v]
  qvh --version | version


Options:
  -h --help       Show this screen.
  -v              Enable verbose mode (debug logging).
  --version       Show version.
  --udid=<udid>   UDID of the device. If not specified, the first found device will be used automatically.

The commands work as following:
	devices		lists iOS devices attached to this host and tells you if video streaming was activated for them
	activate	enables the video streaming config for the device specified by --udid
	record		will start video&audio recording. Video will be saved in a raw h264 file playable by VLC.
				Audio will be saved in a uncompressed wav file.
				Run like: "qvh record /home/yourname/out.h264 /home/yourname/out.wav"
	gstreamer   qvh will open a new window and push AV data to gstreamer.
```

## 3.1 Usage
1. `npm run dist:web:client`
2. connect device
3. `npm run start:web:server`
4. open http://127.0.0.1:8080/ in browser

## 3. Technical Docs
I have written some documentation here [doc/technical_documentation.md](https://github.com/danielpaulus/quicktime_video_hack/blob/master/doc/technical_documentation.md)
So if you are just interested in the protocol or if you want to implement this in a different programming language than golang, read the docs.

I have given up on windows support  :-)
~~[Port to Windows](https://github.com/danielpaulus/quicktime_video_hack/tree/windows/windows) (I don't know why, but still people use Windows nowadays)~~ Did not find a way to do it



