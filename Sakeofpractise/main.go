package main

import (
	"fmt"
	"syscall"
)

func checkDiskSpace() (int64, error) {
	var stat syscall.Statfs_t
	// Get filesystem statistics for the root directory
	if err := syscall.Statfs("/", &stat); err != nil {
		return 0, err
	}
	// Convert uint64 to int64 and calculate available space in bytes
	return int64(stat.Bavail) * int64(stat.Bsize), nil
}

func main() {
	// Check available disk space
	availableSpace, err := checkDiskSpace()
	if err != nil {
		fmt.Printf("Error checking disk space: %v\n", err)
		return
	}

	// Print available space in MB
	fmt.Printf("Available disk space: %.2f MB\n", float64(availableSpace)/1e6)
}
