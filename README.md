# 🩺 EMR Logs Analyser CLI

A powerful command-line tool written in **Go** for parsing and analyzing **EMR system logs** — including **Tomcat (Catalina)** and **OpenMRS (Log4j)** logs.  
It helps you extract insights, detect errors, and summarize patterns from large log files.

---

## 📋 Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Installation](#installation)  
4. [Usage](#usage)  
5. [Command Reference](#command-reference)  
6. [Supported Log Formats](#supported-log-formats)  
7. [Filters](#filters)  
8. [Examples](#examples)  
9. [Using the Makefile](#using-the-makefile)  
10. [Development](#development)  
11. [Planned Enhancements](#planned-enhancements)  

---

## 🧩 Overview

The **EMR Logs Analyser** helps health system administrators and developers quickly make sense of massive EMR server logs.  
It supports both **Tomcat (Catalina)** and **OpenMRS (Log4j)** formats, with filtering and summary statistics built in.

---

## ✨ Features

✅ Parse and analyze **Tomcat / Catalina** logs  
✅ Parse and analyze **OpenMRS (Log4j / SLF4J)** logs  
✅ Apply **filters** for level, thread, class, IP, or status  
✅ Generate **summary statistics**  
✅ Handle **large log files** efficiently (stream-based reading)  
✅ Export results to file  
✅ Extensible architecture — easily add new log formats  

---

## ⚙️ Installation

### **Clone the repository**
```bash
git clone https://github.com/<your-username>/emr-logs-analyser.git
cd emr-logs-analyser
```

### **Build the binary**
```bash
make build
```

The binary `emr-logs-analyser` will be created in the current directory.

---

## 🚀 Usage

You can run the tool directly with Go or via the compiled binary.

### **Run with Go**
```bash
go run main.go analyse --logfile=logs/catalina.out --type=catalina --level=ERROR
```

### **Run compiled binary**
```bash
./emr-logs-analyser analyse --logfile=logs/openmrs.log --type=openmrs --level=ERROR
```

---

## 🧭 Command Reference

### **analyse**
> Analyse EMR logs to extract useful information and insights.

#### **Flags**
| Flag | Shorthand | Description | Example |
|------|------------|--------------|----------|
| `--logfile` | `-f` | Path to the log file to analyze | `--logfile=catalina.out` |
| `--type` | `-t` | Type of logs (`catalina`, `apache`, `openmrs`) | `--type=catalina` |
| `--output` | `-o` | Output file for analysis results | `--output=stats.json` |
| `--stats` |  | Display summary statistics | `--stats` |

---

## 🧱 Supported Log Formats

### 🐱 Catalina / Tomcat Logs
```
03-Oct-2025 09:12:06.003 INFO [main] org.apache.catalina.startup.Catalina.start Server startup in [4234] milliseconds
```

### 🩺 OpenMRS (Log4j / SLF4J) Logs
```
WARN - SessionImpl.createCriteria(1837) |2025-10-04T08:45:59,974| HHH90000022: Hibernate's legacy org.hibernate.Criteria API is deprecated
ERROR - FlagServiceImpl.generateFlagsForPatient(113) |2025-10-04T08:45:59,985| Unable to test flag Eligible for HIV program on patient #7032
```

---

## 🎯 Filters

| Filter | Type | Description | Example |
|--------|------|--------------|----------|
| `--level` | Catalina / OpenMRS | Filter by log level (INFO, WARN, ERROR) | `--level=ERROR` |
| `--thread` | Catalina | Filter by thread name | `--thread=main` |
| `--class` | Catalina / OpenMRS | Filter by class or package | `--class=org.openmrs.module` |
| `--ip` | Apache | Filter by IP address | `--ip=192.168.1.1` |
| `--path` | Apache | Filter by request path | `--path=/api/patients` |
| `--status` | Apache | Filter by HTTP status code | `--status=500` |

---

## 💡 Examples

### 1. Analyze Catalina logs
```bash
./emr-logs-analyser analyse --logfile=logs/catalina.out --type=catalina --level=ERROR --stats
```

### 2. Analyze OpenMRS application logs
```bash
./emr-logs-analyser analyse --logfile=logs/openmrs.log --type=openmrs --level=WARN
```

### 3. Export statistics to JSON
```bash
./emr-logs-analyser analyse -f logs/catalina.out -t catalina -o stats.json --stats
```

---

## 🧰 Using the Makefile

The included `Makefile` provides shortcuts for common tasks:

```makefile
build:      ## Build the binary
	go build -o emr-logs-analyser main.go

run:         ## Run the CLI with arguments
	./emr-logs-analyser $(ARGS)

help:        ## Show available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
```

### Example usage:
```bash
make run ARGS="analyse --logfile=logs/openmrs.log --type=openmrs --level=ERROR --stats"
```

---

## 🧪 Development

### Directory Structure
```
.
├── cmd/
│   └── analyse.go         # Cobra command definition
├── parser/
│   ├── apache_parser.go
│   ├── catalina_parser.go
│   └── openmrs_parser.go
├── analyser/
│   └── stats.go
├── main.go
└── Makefile
```

### Run locally
```bash
go run main.go analyse --logfile=logs/catalina.out --type=catalina
```

---

## 🚀 Planned Enhancements

- [ ] Add **progress tracking** for large log files  
- [ ] Add **JSON / CSV export** options  
- [ ] Add **REST API** layer for frontend integration (React dashboard)  
- [ ] Implement **log correlation** (linking errors to user actions)  
- [ ] Add **visual summary report generation (PDF)**  

---

## 👨‍💻 Author
**Jabar Jeremy**  
Full-Stack Developer • Health Tech | Go | React | Kotlin | Cloud-Native Architect  
