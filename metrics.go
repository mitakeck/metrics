package main

import (
	"github.com/k0kubun/pp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// Values はメトリクス名とその数値の kv map
type Values map[string]float64

func getCPUMetrics() (Values, error) {
	cpus, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}
	pp.Println(cpus)
	ret := make(map[string]float64, 11*len(cpus))
	for _, c := range cpus {
		ret["cpu."+c.CPU+".user"] = c.User
		ret["cpu."+c.CPU+".system"] = c.System
		ret["cpu."+c.CPU+".idle"] = c.Idle
		ret["cpu."+c.CPU+".nice"] = c.Nice
		ret["cpu."+c.CPU+".iowait"] = c.Iowait
		ret["cpu."+c.CPU+".irq"] = c.Irq
		ret["cpu."+c.CPU+".softirq"] = c.Softirq
		ret["cpu."+c.CPU+".steal"] = c.Steal
		ret["cpu."+c.CPU+".guest"] = c.Guest
		ret["cpu."+c.CPU+".guestnice"] = c.GuestNice
		ret["cpu."+c.CPU+".stolen"] = c.Stolen
	}

	return ret, nil
}

func getMemoryMetics() (Values, error) {
	metric, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	ret := map[string]float64{
		"memory.total":     float64(metric.Total),
		"memory.available": float64(metric.Available),
		"memory.used":      float64(metric.Used),
		"memory.percent":   metric.UsedPercent,
		"memory.free":      float64(metric.Free),
		"memory.active":    float64(metric.Active),
		"memory.inactive":  float64(metric.Inactive),
		"memory.wired":     float64(metric.Wired),
		"memory.buffers":   float64(metric.Buffers),
		"memory.cached":    float64(metric.Cached),
	}

	return ret, nil
}

func getNetworkMetics() (Values, error) {
	networks, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]float64, 2*len(networks))

	for _, network := range networks {
		ret["network."+network.Name+".sent"] = float64(network.BytesSent)
		ret["network."+network.Name+".recv"] = float64(network.BytesRecv)
	}

	return ret, nil
}

func getDiskMetrics() (Values, error) {
	disks, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	pp.Println(disks)

	ret := make(map[string]float64, 4*len(disks))

	for _, d := range disks {
		metric, _ := disk.Usage(d.Mountpoint)
		ret["disk."+d.Device+".total"] = float64(metric.Total)
		ret["disk."+d.Device+".free"] = float64(metric.Free)
		ret["disk."+d.Device+".used"] = float64(metric.Used)
		ret["disk."+d.Device+".percent"] = metric.UsedPercent
	}

	return ret, nil
}
