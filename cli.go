package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/mdb/proto"
)

func ipv4ToString(ipv4 uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, ipv4)
	return ip.String()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewMDBServiceClient(conn)

	resp, err := client.ListMachines(ctx, &pb.ListMachinesRequest{})
	if err != nil {
		log.Fatalf("Unable to drain queue: %v", err)
	}

	fmt.Printf("[host]\ntoru\n\n")

	fmt.Printf("[dev]\n")
	for _, machine := range resp.GetMachines() {
		if machine.GetUse() == pb.MachineUse_MACHINE_USE_DEV_DESKTOP || machine.GetUse() == pb.MachineUse_MACHINE_USE_DEV_SERVER {
			fmt.Printf("%v # %v\n", ipv4ToString(machine.GetIpv4()), machine.GetHostname())
		}
	}

	fmt.Printf("\n[raspberrypi]\n")
	for _, machine := range resp.GetMachines() {
		if machine.GetType() == pb.MachineType_MACHINE_TYPE_RASPBERRY_PI {
			if !strings.Contains(machine.GetHostname(), "homeassistant") {
				fmt.Printf("%v\n", ipv4ToString(machine.GetIpv4()))
			}
		}
	}
}
