package logger

import (
	"log"
	"os"
)

type LogType int

const (
	GeneralLog LogType = 1
	ErrorLog
)

type Manager struct {
	logEntries map[LogType][]*LogEntry
	tick       int
}

type LogEntry struct {
	message    string
	tick       int
	timeString string
	entity     *LocatableEntity
}

var Logger = &Manager{}

func InitLogger() {
	Logger.logEntries = make(map[LogType][]*LogEntry)
	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func (l *LogEntry) Message() string {
	return l.message
}

func General(message string, entity *LocatableEntity) {
	Log(GeneralLog, message, entity)
}

func Log(key LogType, message string, entity *LocatableEntity) {
	aLog := &LogEntry{}
	aLog.message = message
	aLog.tick = Logger.tick
	aLog.entity = entity
	Logger.addLogEntry(key, aLog)
	if key != GeneralLog {
		Logger.addLogEntry(GeneralLog, aLog)
	}

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err2 := file.WriteString(message + "\n"); err2 != nil {
		log.Println(err)
	}
}

func (l *Manager) addLogEntry(key LogType, entry *LogEntry) {
	l.InitKey(key)
	l.logEntries[key] = append(l.logEntries[key], entry)
}

func Get(key LogType, amount int) []*LogEntry {
	Logger.InitKey(key)
	if len(Logger.logEntries[key]) <= amount {
		return Logger.logEntries[key]
	}
	return Logger.logEntries[key][len(Logger.logEntries[key])-amount:]
}

func (l *Manager) InitKey(key LogType) {
	if _, ok := l.logEntries[key]; !ok {
		l.logEntries[key] = make([]*LogEntry, 0)
	}
}
