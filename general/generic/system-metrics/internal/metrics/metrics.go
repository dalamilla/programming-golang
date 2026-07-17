package metrics

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"log"
	"time"
)

type SystemMetrics struct {
	Hostname string             `json:"hostname"`
	OS       string             `json:"os"`
	Uptime   uint64             `json:"uptime"`
	CPU      CPU                `json:"cpu"`
	Disk     Disk               `json:"disk"`
	Memory   Memory             `json:"memory"`
	Network  []NetworkInterface `json:"network"`
}

type SystemMetricsError struct {
	Error string `json:"error"`
}

type CPU struct {
	Cores       int     `json:"core"`
	UsageActive float64 `json:"used-active"`
	UsageIDLE   float64 `json:"used-idle"`
	UsageSystem float64 `json:"usage-system"`
	UsageUser   float64 `json:"usage-user"`
}

type Disk struct {
	Path        string  `json:"path"`
	Free        float64 `json:"free"`
	Total       float64 `json:"total"`
	UsedPercent float64 `json:"used-percent"`
}

type Memory struct {
	Free        uint64  `json:"free"`
	Total       uint64  `json:"total"`
	UsedPercent float64 `json:"used-percent"`
	Buffered    uint64  `json:"buffered"`
	Cached      uint64  `json:"cached"`
}

type NetworkInterface struct {
	Interface string   `json:"interface"`
	IPs       []string `json:"ips"`
}

func GetCPU() (CPU, error) {
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		log.Printf("Error retrieving cpu info: %v", err)
		return CPU{}, fmt.Errorf("reading cpu information: %w", err)
	}

	before, _ := cpu.Times(false)
	time.Sleep(1 * time.Second)
	after, _ := cpu.Times(false)

	t1 := before[0]
	t2 := after[0]

	userDelta := t2.User - t1.User
	systemDelta := t2.System - t1.System
	idleDelta := t2.Idle - t1.Idle
	niceDelta := t2.Nice - t1.Nice
	iowaitDelta := t2.Iowait - t1.Iowait
	irqDelta := t2.Irq - t1.Irq
	softirqDelta := t2.Softirq - t1.Softirq
	stealDelta := t2.Steal - t1.Steal

	totalDelta := userDelta + systemDelta + idleDelta + niceDelta + iowaitDelta + irqDelta + softirqDelta + stealDelta

	userPercent := (userDelta / totalDelta) * 100
	systemPercent := (systemDelta / totalDelta) * 100
	idlePercent := (idleDelta / totalDelta) * 100
	activePercent := 100.0 - idlePercent

	return CPU{
		Cores:       logicalCores,
		UsageActive: activePercent,
		UsageIDLE:   idlePercent,
		UsageSystem: systemPercent,
		UsageUser:   userPercent,
	}, nil
}

func GetDisk() (Disk, error) {
	path := "/"
	const GiB = float64(1024 * 1024 * 1024)

	usage, err := disk.Usage(path)
	if err != nil {
		log.Printf("Error retrieving disk info: %v", err)
		return Disk{}, fmt.Errorf("reading disk information: %w", err)
	}
	return Disk{
		Path:        path,
		Free:        float64(usage.Free) / GiB,
		Total:       float64(usage.Total) / GiB,
		UsedPercent: usage.UsedPercent,
	}, nil
}

func GetMemory() (Memory, error) {
	const MiB = 1024 * 1024

	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error retrieving memory info: %v", err)
		return Memory{}, fmt.Errorf("reading memory information: %w", err)
	}

	return Memory{
		Free:        v.Free / MiB,
		Total:       v.Total / MiB,
		UsedPercent: v.UsedPercent,
		Buffered:    v.Buffers / MiB,
		Cached:      v.Cached / MiB,
	}, nil
}

func GetNetwork() ([]NetworkInterface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error retrieving interfaces info: %v", err)
		return []NetworkInterface{}, fmt.Errorf("reading network information: %w", err)

	}

	var networkInterfaces []NetworkInterface

	for _, iface := range interfaces {
		var ips []string
		for _, addr := range iface.Addrs {
			ips = append(ips, addr.Addr)
		}

		if len(ips) > 0 {
			networkInterfaces = append(networkInterfaces, NetworkInterface{
				Interface: iface.Name,
				IPs:       ips,
			})
		}
	}
	return networkInterfaces, nil

}

func GetSystemMetrics() (SystemMetrics, error) {
	log.Println("Get system metrics")

	info, err := host.Info()
	if err != nil {
		log.Printf("Error retrieving host info: %v", err)
		return SystemMetrics{}, fmt.Errorf("reading platform information: %w", err)
	}
	hostname := info.Hostname
	os := info.OS
	uptime := info.Uptime

	cpu, err := GetCPU()
	if err != nil {
		return SystemMetrics{}, err
	}

	disk, err := GetDisk()
	if err != nil {
		return SystemMetrics{}, err
	}

	memory, err := GetMemory()
	if err != nil {
		return SystemMetrics{}, err
	}

	network, err := GetNetwork()
	if err != nil {
		return SystemMetrics{}, err
	}

	return SystemMetrics{
		Hostname: hostname,
		OS:       os,
		Uptime:   uptime,
		CPU:      cpu,
		Disk:     disk,
		Memory:   memory,
		Network:  network,
	}, nil
}
