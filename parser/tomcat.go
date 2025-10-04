package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// Unified structure for both Tomcat & OpenMRS logs
type TomcatLogEntry struct {
	Timestamp string
	Level     string
	Thread    string
	Class     string
	Message   string
}

// --- REGEX PATTERNS ---

// 1️⃣ OpenMRS (Log4j-style Catalina)
var log4jPattern = regexp.MustCompile(
	`^(?P<Level>[A-Z]+)\s*-\s*(?P<Class>[A-Za-z0-9_.<>$]+(?:\([0-9]+\))?)\s*\|(?P<Timestamp>[0-9T:,\-]+)\|\s*(?P<Message>.*)$`)

// 2️⃣ Native Tomcat (Catalina.out style)
var tomcatPattern = regexp.MustCompile(
	`^(?P<Timestamp>\d{2}-[A-Za-z]{3}-\d{4} \d{2}:\d{2}:\d{2}\.\d{3})\s+(?P<Level>[A-Z]+)\s+\[(?P<Thread>[^\]]+)\]\s+(?P<Class>\S+)\s+(?P<Message>.*)$`)

// --- DETECTION + PARSING ---

func ParseTomcatLogLine(line string) (*TomcatLogEntry, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	// Try OpenMRS Log4J-style pattern first
	if log4jPattern.MatchString(line) {
		matches := log4jPattern.FindStringSubmatch(line)
		return &TomcatLogEntry{
			Level:     matches[1],
			Class:     matches[2],
			Timestamp: matches[3],
			Message:   matches[4],
		}, nil
	}

	// Fallback: Native Tomcat format
	if tomcatPattern.MatchString(line) {
		matches := tomcatPattern.FindStringSubmatch(line)
		return &TomcatLogEntry{
			Timestamp: matches[1],
			Level:     matches[2],
			Thread:    matches[3],
			Class:     matches[4],
			Message:   matches[5],
		}, nil
	}

	// No match
	return nil, fmt.Errorf("log line does not match expected format: %s", line)
}

func ParseTomcatLogFile(path string) ([]*TomcatLogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []*TomcatLogEntry

	for scanner.Scan() {
		line := scanner.Text()
		entry, err := ParseTomcatLogLine(line)
		if err != nil {
			// Skip JVM/system messages and noise
			log.Debug().Msgf("Skipping non-log line: %s", line)
			continue
		}
		if entry != nil {
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %w", err)
	}

	if len(entries) == 0 {
		log.Warn().Msg("No valid Catalina or Log4j log entries found.")
	}

	return entries, nil
}
