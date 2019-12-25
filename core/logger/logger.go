package logger

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
