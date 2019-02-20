package main

import (
	"github.com/nassor/goav/avcodec"
	"github.com/nassor/goav/avdevice"
	"github.com/nassor/goav/avfilter"
	"github.com/nassor/goav/avformat"
	"github.com/nassor/goav/avutil"
	"github.com/nassor/goav/swresample"
	"github.com/nassor/goav/swscale"
	"log"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
