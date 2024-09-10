package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"cube/stats"
	"cube/utils"
)

type Node struct {
	Name            string
	IP              string
	Api             string
	Memory          int64
	MemoryAllocated int64
	Disk            int64
	DiskAllocated   int64
	Stats           stats.Stats
	Role            string
	TaskCount       int
}

func NewNode(name, api, role string) *Node {
	return &Node{
		Name: name,
		Api:  api,
		Role: role,
	}
}

func (n *Node) GetStats() (*stats.Stats, error) {
	var resp *http.Response
	var err error

	url := fmt.Sprintf("%s/stats", n.Api)
	resp, err = utils.HTTPWithRetry(http.Get, url)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to %v. Permanent failure.\n", n.Api)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Error retrieving stats from %v: %v", n.Api, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	defer resp.Body.Close()

	var stats stats.Stats
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		msg := fmt.Sprintf("error decoding message while getting stats for node %s", n.Name)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	if stats.MemStats == nil || stats.DiskStats == nil {
		return nil, fmt.Errorf("error getting stats from node %s", n.Name)
	}

	n.Memory = int64(stats.MemTotalKb())
	n.Disk = int64(stats.DiskTotal())
	n.Stats = stats

	log.Printf("collect stats from worker %s success, CPU detail: linux.CPUStat{Id:cpu, User:%d, Nice:%d, System:%d, Idle:%d, IOWait:%d, IRQ:%d, SoftIRQ:%d, Steal:%d, Guest:%d, GuestNice:%d}", n.Api, n.Stats.CPUStats.User, n.Stats.CPUStats.Nice, n.Stats.CPUStats.System, n.Stats.CPUStats.Idle, n.Stats.CPUStats.IOWait, n.Stats.CPUStats.IRQ, n.Stats.CPUStats.SoftIRQ, n.Stats.CPUStats.Steal, n.Stats.CPUStats.Guest, n.Stats.CPUStats.GuestNice)
	log.Printf("collect stats from worker %s success, Disk detail: linux.Disk{All:%d, Used:%d, Free:%d, FreeInodes:%d}", n.Api, n.Stats.DiskStats.All, n.Stats.DiskStats.Used, n.Stats.DiskStats.Free, n.Stats.DiskStats.FreeInodes)
	log.Printf("collect stats from worker %s success, Memory detail: linux.MemInfo{MemTotal:%d, MemFree:%d, MemAvailable:%d, Buffers:%d, Cached:%d, SwapCached:%d, Active:%d, Inactive:%d, ActiveAnon:%d, InactiveAnon:%d, ActiveFile:%d, InactiveFile:%d, Unevictable:%d, Mlocked:%d, SwapTotal:%d, SwapFree:%d, Dirty:%d, Writeback:%d, AnonPages:%d, Mapped:%d, Shmem:%d, Slab:%d, SReclaimable:%d, SUnreclaim:%d, KernelStack:%d, PageTables:%d, NFS_Unstable:%d, Bounce:%d, WritebackTmp:%d, CommitLimit:%d, Committed_AS:%d, VmallocTotal:%d, VmallocUsed:%d, VmallocChunk:%d, HardwareCorrupted:%d, AnonHugePages:%d, HugePages_Total:%d, HugePages_Free:%d, HugePages_Rsvd:%d, HugePages_Surp:%d, Hugepagesize:%d, DirectMap4k:%d, DirectMap2M:%d, DirectMap1G:%d}", n.Api, n.Stats.MemStats.MemTotal, n.Stats.MemStats.MemFree, n.Stats.MemStats.MemAvailable, n.Stats.MemStats.Buffers, n.Stats.MemStats.Cached, n.Stats.MemStats.SwapCached, n.Stats.MemStats.Active, n.Stats.MemStats.Inactive, n.Stats.MemStats.ActiveAnon, n.Stats.MemStats.InactiveAnon, n.Stats.MemStats.ActiveFile, n.Stats.MemStats.InactiveFile, n.Stats.MemStats.Unevictable, n.Stats.MemStats.Mlocked, n.Stats.MemStats.SwapTotal, n.Stats.MemStats.SwapFree, n.Stats.MemStats.Dirty, n.Stats.MemStats.Writeback, n.Stats.MemStats.AnonPages, n.Stats.MemStats.Mapped, n.Stats.MemStats.Shmem, n.Stats.MemStats.Slab, n.Stats.MemStats.SReclaimable, n.Stats.MemStats.SUnreclaim, n.Stats.MemStats.KernelStack, n.Stats.MemStats.PageTables, n.Stats.MemStats.NFS_Unstable, n.Stats.MemStats.Bounce, n.Stats.MemStats.WritebackTmp, n.Stats.MemStats.CommitLimit, n.Stats.MemStats.Committed_AS, n.Stats.MemStats.VmallocTotal, n.Stats.MemStats.VmallocUsed, n.Stats.MemStats.VmallocChunk, n.Stats.MemStats.HardwareCorrupted, n.Stats.MemStats.AnonHugePages, n.Stats.MemStats.HugePages_Total, n.Stats.MemStats.HugePages_Free, n.Stats.MemStats.HugePages_Rsvd, n.Stats.MemStats.HugePages_Surp, n.Stats.MemStats.Hugepagesize, n.Stats.MemStats.DirectMap4k, n.Stats.MemStats.DirectMap2M, n.Stats.MemStats.DirectMap1G)

	return &n.Stats, nil
}
