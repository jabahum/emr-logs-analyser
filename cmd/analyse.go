package cmd

import (
	"github/jabahum/emr-log-analyser/analyser"
	"github/jabahum/emr-log-analyser/parser"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(analyseCmd)

	analyseCmd.Flags().StringP("logfile", "f", "", "Path to the EMR log file")
	analyseCmd.Flags().StringP("output", "o", "", "Path to the output file")
	analyseCmd.Flags().StringP("type", "t", "catalina", "Type of analysis to perform (apache|catalina)")

	// Apache filters
	analyseCmd.Flags().String("ip", "", "Filter by IP address (apache only)")
	analyseCmd.Flags().String("path", "", "Filter by request path (apache only)")
	analyseCmd.Flags().String("status", "", "Filter by HTTP status code (apache only)")

	// Catalina filters
	analyseCmd.Flags().String("level", "", "Filter by log level (catalina only)")
	analyseCmd.Flags().String("thread", "", "Filter by thread name (catalina only)")
	analyseCmd.Flags().String("class", "", "Filter by class/package (catalina only)")

	// Stats flag
	analyseCmd.Flags().Bool("stats", false, "Generate and display log statistics")

	_ = analyseCmd.MarkFlagRequired("logfile")
}

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyse EMR logs",
	Long: `Analyse EMR logs (Apache or Catalina) to extract useful information, 
apply filters, and optionally generate statistics.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting log analysis...")

		// --- Read flags ---
		logFile, _ := cmd.Flags().GetString("logfile")
		outputFile, _ := cmd.Flags().GetString("output")
		analysisType, _ := cmd.Flags().GetString("type")
		statsFlag, _ := cmd.Flags().GetBool("stats")

		// Apache filters
		filterIP, _ := cmd.Flags().GetString("ip")
		filterPath, _ := cmd.Flags().GetString("path")
		filterStatus, _ := cmd.Flags().GetString("status")

		// Catalina filters
		filterLevel, _ := cmd.Flags().GetString("level")
		filterThread, _ := cmd.Flags().GetString("thread")
		filterClass, _ := cmd.Flags().GetString("class")

		// --- Validation ---
		if logFile == "" {
			log.Error().Msg("Log file path is required.")
			cmd.Println("Error: missing required flag '--logfile'")
			return
		}

		if analysisType != "apache" && analysisType != "catalina" {
			log.Error().Str("analysisType", analysisType).Msg("Invalid analysis type selected.")
			cmd.Println("Invalid analysis type. Please use 'apache' or 'catalina'.")
			return
		}

		log.Info().
			Str("logFile", logFile).
			Str("outputFile", outputFile).
			Str("analysisType", analysisType).
			Msg("Analysis configuration")

		cmd.Printf("Performing %s log analysis...\n", analysisType)

		// --- Dispatch by type ---
		switch analysisType {
		case "catalina":
			analyzeCatalinaLogs(cmd, logFile, statsFlag, filterLevel, filterThread, filterClass)
		case "apache":
			analyzeApacheLogs(cmd, logFile, statsFlag, filterIP, filterPath, filterStatus)
		}

		log.Info().Msg("Analysis complete.")
	},
}

func analyzeCatalinaLogs(cmd *cobra.Command, logFile string, statsFlag bool, level, thread, class string) {
	log.Info().Msg("Using Tomcat log analyser for Catalina logs...")

	entries, err := parser.ParseTomcatLogFile(logFile)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse Catalina log file")
		return
	}

	// --- Apply filters ---
	if level != "" {
		entries = filterEntriesByLevel(entries, level)
		var converted []analyser.LogEntry
		for _, e := range entries {
			converted = append(converted, analyser.LogEntry{
				Level:   e.Level,
				Thread:  e.Thread,
				Class:   e.Class,
				Message: e.Message,
			})
		}
		analyser.PrintFilteredResults(converted)

	}
	if thread != "" {
		entries = filterEntriesByThread(entries, thread)
		var converted []analyser.LogEntry
		for _, e := range entries {
			converted = append(converted, analyser.LogEntry{
				Level:   e.Level,
				Thread:  e.Thread,
				Class:   e.Class,
				Message: e.Message,
			})
		}
		analyser.PrintFilteredResults(converted)

	}
	if class != "" {
		entries = filterEntriesByClass(entries, class)
		var converted []analyser.LogEntry
		for _, e := range entries {
			converted = append(converted, analyser.LogEntry{
				Level:   e.Level,
				Thread:  e.Thread,
				Class:   e.Class,
				Message: e.Message,
			})
		}
		analyser.PrintFilteredResults(converted)

	}

	cmd.Printf("Processed %d Catalina log entries.\n", len(entries))
	log.Info().Int("entries", len(entries)).Msg("Catalina log entries processed")

	// --- Stats ---
	if !statsFlag {
		return
	}

	log.Info().Msg("Generating Catalina statistics...")
	var converted []analyser.LogEntry
	for _, e := range entries {
		converted = append(converted, analyser.LogEntry{
			Level:   e.Level,
			Thread:  e.Thread,
			Class:   e.Class,
			Message: e.Message,
		})
	}

	stats := analyser.GenerateStatistics(converted)
	stats.PrintSummary()
}
func analyzeApacheLogs(cmd *cobra.Command, logFile string, statsFlag bool, ip, path, status string) {
	log.Info().Msg("Using Apache log analyser for access logs...")

	entries, err := parser.ParseApacheLogFile(logFile)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse Apache log file")
		cmd.Printf("Error parsing Apache log file: %v\n", err)
		return
	}

	// --- Apply filters ---
	if ip != "" {
		entries = filterApacheEntriesByIP(entries, ip)
	}
	if path != "" {
		entries = filterApacheEntriesByPath(entries, path)
	}
	if status != "" {
		entries = filterApacheEntriesByStatus(entries, status)

	}

	cmd.Printf("Processed %d Apache log entries.\n", len(entries))

	// --- Stats ---
	if statsFlag {
		log.Info().Msg("Generating Apache log statistics...")
		// (Optional) implement Apache stats generator next
	}
}

// filterEntriesByLevel filters Catalina log entries by log level.
func filterEntriesByLevel(entries []*parser.TomcatLogEntry, level string) []*parser.TomcatLogEntry {
	var filtered []*parser.TomcatLogEntry
	for _, entry := range entries {
		if entry.Level == level {
			filtered = append(filtered, entry)
		}
	}
	return filtered

}

// filterEntriesByThread filters Catalina log entries by thread name.
func filterEntriesByThread(entries []*parser.TomcatLogEntry, thread string) []*parser.TomcatLogEntry {
	var filtered []*parser.TomcatLogEntry
	for _, entry := range entries {
		if entry.Thread == thread {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// filterEntriesByClass filters Catalina log entries by class/package name.
func filterEntriesByClass(entries []*parser.TomcatLogEntry, class string) []*parser.TomcatLogEntry {
	var filtered []*parser.TomcatLogEntry
	for _, entry := range entries {
		if entry.Class == class {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// filterApacheEntriesByIP filters Apache log entries by IP address.
func filterApacheEntriesByIP(entries []*parser.ApacheLogEntry, ip string) []*parser.ApacheLogEntry {
	var filtered []*parser.ApacheLogEntry
	for _, entry := range entries {
		if entry.IP == ip {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// filterApacheEntriesByPath filters Apache log entries by request path.
func filterApacheEntriesByPath(entries []*parser.ApacheLogEntry, path string) []*parser.ApacheLogEntry {
	var filtered []*parser.ApacheLogEntry
	for _, entry := range entries {
		if entry.Endpoint == path {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// filterApacheEntriesByStatus filters Apache log entries by HTTP status code.
func filterApacheEntriesByStatus(entries []*parser.ApacheLogEntry, status string) []*parser.ApacheLogEntry {
	var filtered []*parser.ApacheLogEntry
	for _, entry := range entries {
		if entry.Status == status {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}
