package analyser

import (
	"fmt"
	"os"
	"text/tabwriter"
)

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

func PrintFilteredResults(entries []LogEntry) {

	if len(entries) == 0 {
		fmt.Println("No matching log entries found")
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, '\t', 0)
	fmt.Fprintln(writer, "LEVEL/\tCLASS\tMESSAGE")
	for _, entry := range entries {
		message := entry.Message
		if len(message) > 80 {
			message = message[:77] + "..."

		}

		fmt.Fprintf(writer, "%s\t\t%s\t\t%s\n", entry.Level, entry.Class, message)
	}

	writer.Flush()
	fmt.Printf("Total Entries : %d\n", len(entries))
}
