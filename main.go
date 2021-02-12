package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {

	subnets := []string{}

	type myType struct {
		a, b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%# v", pretty.Formatter(x))

	_ = subnets

	// arr := strings.Split("10.145.68.48/27", "/")
	// host := arr[0]
	// mask, _ := strconv.Atoi(arr[1])
	// subnetToValid := ipsubnet.SubnetCalculator(host, mask)
	// _ = subnetToValid
	// ip, valid, _ := net.ParseCIDR("10.145.68.48/27")

	// fmt.Println(valid)
	// fmt.Println(ip)
	// fmt.Println(valid.IP)
	// fmt.Println(valid.IP.Equal(ip))

	// for _, subnetStr := range subnets {
	// 	_, subnetMap, _ := net.ParseCIDR(subnetStr)

	// 	if subnetMap.Contains(net.ParseIP(subnetToValid.GetIPAddress())) || subnetMap.Contains(net.ParseIP(subnetToValid.GetBroadcastAddress())) {
	// 		return errors.New("This subnet " + subnet + " is Overlaped by " + subnetStr)
	// 	}
	// }

}
