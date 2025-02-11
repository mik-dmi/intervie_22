package utils

import "github.com/docker/docker/api/types"

func CalculateCPUPercentUnix(stat types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		cpuDelta   = float64(stat.CPUStats.CPUUsage.TotalUsage) - float64(stat.PreCPUStats.CPUUsage.TotalUsage)

		systemDelta = float64(stat.CPUStats.SystemUsage) - float64(stat.PreCPUStats.SystemUsage)
	)
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(stat.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}
