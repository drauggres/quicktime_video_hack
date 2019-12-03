
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/danielpaulus/quicktime_video_hack/screencapture"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	mp *screencapture.MessageProcessor
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run(stopSignal chan interface{}) {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
		if len(h.clients) == 0 && h.mp != nil {
			log.Info("last client left")
		}
		if len(h.clients) > 0 && h.mp == nil  {
			log.Info("New client")
			go func() {
				log.Info("Start")
				deviceList, err := screencapture.FindIosDevices()
				if err != nil {
					log.Fatal("Error finding iOS Devices", err)
				}
				dev := deviceList[0]

				writer := NewNaluHubWriter(h)
				adapter := screencapture.UsbAdapter{}
				mp := screencapture.NewMessageProcessor(&adapter, stopSignal, writer)
				h.mp = &mp

				adapter.StartReading(dev, &mp, stopSignal)
			}()
		}
	}
}


