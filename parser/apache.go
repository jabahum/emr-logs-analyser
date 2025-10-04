package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ApacheLogEntry struct {
	IP        string
	Ident     string
	User      string
	Timestamp string
	Method    string
	Endpoint  string
	Protocol  string
	Status    string
	Size      string
	Referrer  string
	UserAgent string
}

var apacheLogPattern = regexp.MustCompile(`^(?P<IP>\S+) (?P<Ident>\S+) (?P<User>\S+) \[(?P<Timestamp>[^\]]+)\] "(?P<Method>\S+) (?P<Endpoint>\S+) (?P<Protocol>[^"]+)" (?P<Status>\d{3}) (?P<Size>\S+) "(?P<Referrer>[^"]*)" "(?P<UserAgent>[^"]*)"`)

func ParseApacheLogLine(line string) (*ApacheLogEntry, error) {
	matches := apacheLogPattern.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("log line does not match expected format: %s", line)
	}

	entry := &ApacheLogEntry{
		IP:        matches[1],
		Ident:     matches[2],
		User:      matches[3],
		Timestamp: matches[4],
		Method:    matches[5],
		Endpoint:  matches[6],
		Protocol:  matches[7],
		Status:    matches[8],
		Size:      matches[9],
		Referrer:  matches[10],
		UserAgent: matches[11],
	}

	return entry, nil
}

func ParseApacheLogFile(filePath string) ([]*ApacheLogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	var entries []*ApacheLogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		entry, err := ParseApacheLogLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log line: %v", err)
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %v", err)
	}

	return entries, nil
}
