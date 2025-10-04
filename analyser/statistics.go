package analyser

import "fmt"

type LogEntry struct {
	Timestamp string
	Level     string
	Thread    string
	Class     string
	Message   string
}

type LogStats struct {
	TotalEntries int
	ByLevel      map[string]int
	ByClass      map[string]int
	ByThread     map[string]int
}

func GenerateStatistics(entries []LogEntry) LogStats {
	stats := LogStats{
		TotalEntries: len(entries),
		ByLevel:      make(map[string]int),
		ByClass:      make(map[string]int),
		ByThread:     make(map[string]int),
	}

	for _, entry := range entries {
		stats.ByLevel[entry.Level]++
		stats.ByClass[entry.Class]++
		stats.ByThread[entry.Thread]++
	}

	return stats
}

func (s LogStats) PrintSummary() {
	fmt.Println("\n===== Catalina Log Summary =====")
	fmt.Printf("Total Entries: %d\n", s.TotalEntries)

	fmt.Println("\nBy Level:")
	for lvl, count := range s.ByLevel {
		fmt.Printf("  %-8s %d\n", lvl, count)
	}

	fmt.Println("\nTop Classes:")
	for cls, count := range s.ByClass {
		if count > 5 { // display only classes with more than 5 entries
			fmt.Printf("  %-50s %d\n", cls, count)
		}
	}

	fmt.Println("\nTop Threads:")
	for thread, count := range s.ByThread {
		if count > 5 {
			fmt.Printf("  %-30s %d\n", thread, count)
		}
	}
	fmt.Println("=================================")
}
