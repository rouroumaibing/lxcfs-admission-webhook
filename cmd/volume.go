package main

import corev1 "k8s.io/api/core/v1"

// -v /var/lib/lxcfs/proc/cpuinfo:/proc/cpuinfo:ro
// -v /var/lib/lxcfs/proc/diskstats:/proc/diskstats:ro
// -v /var/lib/lxcfs/proc/meminfo:/proc/meminfo:ro
// -v /var/lib/lxcfs/proc/stat:/proc/stat:ro
// -v /var/lib/lxcfs/proc/swaps:/proc/swaps:ro
// -v /var/lib/lxcfs/proc/uptime:/proc/uptime:ro

const lxcfsVol = "lxcfs"

var volumeMountsTemplate = []corev1.VolumeMount{

	{
		Name:      lxcfsVol,
		MountPath: "/proc/cpuinfo",
		SubPath:   "lxcfs/proc/cpuinfo",
		ReadOnly:  true,
	},
	{
		Name:      lxcfsVol,
		MountPath: "/proc/diskstats",
		SubPath:   "lxcfs/proc/diskstats",
		ReadOnly:  true,
	},
	{
		Name:      lxcfsVol,
		MountPath: "/proc/meminfo",
		SubPath:   "lxcfs/proc/meminfo",
		ReadOnly:  true,
	},
	{
		Name:      lxcfsVol,
		MountPath: "/proc/stat",
		SubPath:   "lxcfs/proc/stat",
		ReadOnly:  true,
	},
	{
		Name:      lxcfsVol,
		MountPath: "/proc/swaps",
		SubPath:   "lxcfs/proc/swaps",
		ReadOnly:  true,
	},
	{
		Name:      lxcfsVol,
		MountPath: "/proc/uptime",
		SubPath:   "lxcfs/proc/uptime",
		ReadOnly:  true,
	},
}

var volumesTemplate = []corev1.Volume{
	{
		Name: lxcfsVol,
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/",
				Type: func() *corev1.HostPathType {
					pt := corev1.HostPathDirectoryOrCreate
					return &pt
				}(),
			},
		},
	},
}
