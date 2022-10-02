package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chifflier/nfqueue-go/nfqueue"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func real_callback(payload *nfqueue.Payload) int {
	fmt.Println("Real callback")
	fmt.Printf("  id: %d\n", payload.Id)
	// fmt.Println(hex.Dump(payload.Data))
	// Decode a packet
	packet := gopacket.NewPacket(payload.Data, layers.LayerTypeIPv4, gopacket.Default)

	// Get the TCP layer from this packet
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		fmt.Println("This is a TCP packet!")
		// Get actual TCP data from this layer
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
	}

	if packet.ApplicationLayer() != nil {
		body := packet.ApplicationLayer().Payload()
		if len(body) > 0 {
			bodyStr := string(body)

			fmt.Println(bodyStr)

			// modify payload of application layer
			*packet.ApplicationLayer().(*gopacket.Payload) = []byte("Hello World!")

			// if its tcp we need to tell it which network layer is being used
			// to be able to handle multiple protocols we can add a if clause around this
			packet.TransportLayer().(*layers.TCP).SetNetworkLayerForChecksum(packet.NetworkLayer())

			buffer := gopacket.NewSerializeBuffer()
			options := gopacket.SerializeOptions{
				ComputeChecksums: true,
				FixLengths:       true,
			}

			// Serialize Packet to get raw bytes
			if err := gopacket.SerializePacket(buffer, options, packet); err != nil {
				log.Fatalln(err)
			}

			packetBytes := buffer.Bytes()

			fmt.Println(packetBytes)

			fmt.Println("-- ")
			payload.SetVerdictModified(nfqueue.NF_ACCEPT, payload.Data)
			return 0
		}
	}

	// Iterate over all layers, printing out each layer type
	// for _, layer := range packet.Layers() {
	// 	fmt.Println("PACKET LAYER:", layer.LayerType())

	// 	body := string(layer.LayerPayload())

	// 	fmt.Println(body)

	// 	if strings.HasPrefix(body, "GET") {
	// 		fmt.Print("Modify Packet!!!!!")
	// 		fmt.Println(hex.Dump(payload.Data))
	// 	}
	// }

	fmt.Println("-- ")
	payload.SetVerdictModified(nfqueue.NF_ACCEPT, payload.Data)
	// payload.SetVerdict(nfqueue.NF_ACCEPT)
	return 0
}

func main() {
	q := new(nfqueue.Queue)

	q.SetCallback(real_callback)

	q.Init()

	q.Unbind(syscall.AF_INET)
	q.Bind(syscall.AF_INET)

	q.CreateQueue(0)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			// sig is a ^C, handle it
			_ = sig
			q.StopLoop()
		}
	}()

	// XXX Drop privileges here

	q.Loop()
	q.DestroyQueue()
	q.Close()
	os.Exit(0)
}
