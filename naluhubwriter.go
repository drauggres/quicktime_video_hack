package main

import (
	"encoding/binary"
	"github.com/danielpaulus/quicktime_video_hack/screencapture/coremedia"
	log "github.com/sirupsen/logrus"
)

type NaluHubWriter struct {
	hub *Hub
}

func NewNaluHubWriter(hub *Hub) NaluHubWriter {
	return NaluHubWriter{hub: hub}
}

func (nhw NaluHubWriter) Consume(buf coremedia.CMSampleBuffer) error {
	if buf.HasFormatDescription {
		log.Info("PPS " + buf.FormatDescription.String())
			err := nhw.writeNalu(buf.FormatDescription.PPS)
			if err != nil {
				return err
			}
			err = nhw.writeNalu(buf.FormatDescription.SPS)
			if err != nil {
				return err
			}
	}
	return nhw.writeNalus(buf.SampleData)
}

func (nhw NaluHubWriter) writeNalus(bytes []byte) error {
	slice := bytes
	for len(slice) > 0 {
		length := binary.BigEndian.Uint32(slice)
		err := nhw.writeNalu(slice[4 : length+4])
		if err != nil {
			return err
		}
		slice = slice[length+4:]
	}
	return nil
}

func (nhw NaluHubWriter) writeNalu(bytes []byte) error {
	if len(bytes) > 0 {
		bytes = append([]byte{00, 00, 00, 01}, bytes...);
		nhw.hub.broadcast <- bytes
	}
	return nil
}
