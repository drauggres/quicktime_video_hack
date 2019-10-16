package main

import (
    "context"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
    "os/signal"

    "github.com/danielpaulus/quicktime_video_hack/screencapture"
)

func main() {
    stopSignal := make(chan interface{})
    waitForSigInt(stopSignal)
    activate()
    hub := newHub()
    go hub.run(stopSignal)

    m := http.NewServeMux()
    s := http.Server{Addr: ":8080", Handler: m}

    m.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("ws-scrcpy/dist/public"))))

    m.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
       serveWs(hub, w, r)
    })
    go func() {
        err := s.ListenAndServe()
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }
    }()

    <-stopSignal
    s.Shutdown(context.Background())
}


func waitForSigInt(stopSignalChannel chan interface{}) {
    log.Info("c := make(chan os.Signal, 1)")
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        for sig := range c {
            log.Infof("Signal received: %s", sig)
            var stopSignal interface{}
            stopSignalChannel <- stopSignal
        }
    }()
}

// Just dump a list of what was discovered to the console
func devices(devices []screencapture.IosDevice) {
    log.Infof("(%d) iOS Devices with UsbMux Endpoint:", len(devices))

    output := screencapture.PrintDeviceDetails(devices)
    log.Info(output)
}

// This command is for testing if we can enable the hidden Quicktime device config
func activate() {
    cleanup := screencapture.Init()
    deviceList, err := screencapture.FindIosDevices()
    defer cleanup()
    if err != nil {
        log.Fatal("Error finding iOS Devices", err)
    }

    log.Info("iOS Devices with UsbMux Endpoint:")

    output := screencapture.PrintDeviceDetails(deviceList)
    log.Info(output)

    err = screencapture.EnableQTConfig(deviceList)
    if err != nil {
        log.Fatal("Error enabling QT config", err)
    }

    qtDevices, err := screencapture.FindIosDevicesWithQTEnabled()
    if err != nil {
        log.Fatal("Error finding QT Devices", err)
    }
    qtOutput := screencapture.PrintDeviceDetails(qtDevices)
    if len(qtDevices) != len(deviceList) {
        log.Warnf("Less qt devices (%d) than plain usbmux devices (%d)", len(qtDevices), len(deviceList))
    }
    log.Info("iOS Devices with QT Endpoint:")
    log.Info(qtOutput)
}
