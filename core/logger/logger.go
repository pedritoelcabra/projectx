package logger

import (
	"github.com/pedritoelcabra/projectx/world"
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
	entity     *world.Entity
}

var Logger = &Manager{}

func InitLogger() {
	Logger.logEntries = make(map[LogType][]*LogEntry)
}

func (l *LogEntry) Message() string {
	return l.message
}

func General(message string, entity *world.Entity) {
	Log(GeneralLog, message, entity)
}

func Log(key LogType, message string, entity *world.Entity) {
	aLog := &LogEntry{}
	aLog.message = message
	aLog.tick = Logger.tick
	aLog.entity = entity
	Logger.addLogEntry(key, aLog)
	if key != GeneralLog {
		Logger.addLogEntry(GeneralLog, aLog)
	}
}

func (l *Manager) addLogEntry(key LogType, entry *LogEntry) {
	l.InitKey(key)
	l.logEntries[key] = append(l.logEntries[key], entry)
}

func Get(key LogType) []*LogEntry {
	Logger.InitKey(key)
	return Logger.logEntries[key]
}

func (l *Manager) InitKey(key LogType) {
	if _, ok := l.logEntries[key]; !ok {
		l.logEntries[key] = make([]*LogEntry, 0)
	}
}
