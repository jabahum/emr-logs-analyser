package analyser

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
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
		fmt.Println("⚠️  No matching log entries found.")
		return
	}
	fmt.Println(colorCyan + "🔍 Filtered Log Results:" + colorReset)
	fmt.Println(strings.Repeat("=", 90))

	// Create a tabwriter for clean column alignment
	writer := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintf(writer, "%s\t%s\t%s\n", "LEVEL", "CLASS", "MESSAGE")
	fmt.Fprintf(writer, "%s\t%s\t%s\n", strings.Repeat("-", 10), strings.Repeat("-", 50), strings.Repeat("-", 100))

	for _, entry := range entries {
		levelColor := colorForLevel(entry.Level)
		level := fmt.Sprintf("%s%-8s%s", levelColor, entry.Level, colorReset)
		class := truncate(entry.Class, 50)
		message := truncate(entry.Message, 100)
		fmt.Fprintf(writer, "%s\t%s\t%s\n", level, class, message)
	}

	writer.Flush()
	fmt.Println(strings.Repeat("=", 90))
	fmt.Printf(colorGreen+"✅ Total Entries: %d"+colorReset+"\n\n", len(entries))
	fmt.Println()
}

// truncate shortens a string and appends "..." if it exceeds max length
func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}

// colorForLevel returns color based on log level
func colorForLevel(level string) string {
	switch strings.ToUpper(level) {
	case "ERROR", "SEVERE", "FATAL":
		return colorRed
	case "WARN", "WARNING":
		return colorYellow
	case "INFO":
		return colorGreen
	case "DEBUG", "TRACE":
		return colorBlue
	default:
		return colorWhite
	}
}
