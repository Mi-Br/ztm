//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) Average() float32 {
	return ByteAverage(&b.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) Average() float32 {
	var sum Celcius = 0
	if len(c.temp) > 0 {
		for _, v := range c.temp {
			sum += v
		}
		return float32(sum) / float32(len(c.temp))
	} else {
		return 0
	}
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) Average() float32 {
	return ByteAverage(&m.amount)
}

func ByteAverage(b *[]Bytes) float32 {
	var sum Bytes = 0
	if len(*b) > 0 {
		for _, v := range *b {
			sum += v
		}
		return float32(sum) / float32(len(*b))
	} else {
		return 0
	}
}

type Averagable interface {
	Average()
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
	Averagable
}

func (d *Dashboard) Average() {
	bAvg, cAvg, mAvg := d.BandwidthUsage.Average(), d.CpuTemp.Average(), d.MemoryUsage.Average()
	fmt.Printf("Dashboard running 5s average is \n ")
	fmt.Printf("| Bandwidth | CPUTemp | Memmory | \n")
	fmt.Printf("|    %v     |   %v    |    %v   | \n", bAvg, cAvg, mAvg)
	fmt.Printf("--------------------------------- \n")
}
func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	var D Dashboard = Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	D.Average()

}
