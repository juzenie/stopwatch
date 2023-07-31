package swer

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

type Sw interface {
	Start(taskName string)
	Stop()
	PrettyPrint() string
}

type TaskInfo struct {
	TaskName  string
	TimeNanos int64
}

type StopWatch struct {
	id              string
	KeepTaskList    bool
	taskList        []TaskInfo
	startTimeNanos  int64
	currentTaskName string
	lastTaskInfo    *TaskInfo
	taskCount       int
	totalTimeNanos  int64
}

func NewTaskInfo() *StopWatch {
	return &StopWatch{
		KeepTaskList: true,
	}
}

func NanosToMillis(duration int64) int64 {
	return duration / 1_000_000.0
}
func nanosToSeconds(duration int64) int64 {
	return duration / 1_000_000_000.0
}
func (w *StopWatch) Start(taskName string) {
	if len(w.currentTaskName) != 0 {
		log.Fatal("Can't start StopWatch: it's already running")
	}
	w.currentTaskName = taskName
	w.startTimeNanos = time.Now().UnixNano()
}

func (w *StopWatch) Stop() {
	if len(w.currentTaskName) == 0 {
		log.Fatal("Can't stop StopWatch: it's not running")
	}
	lastTime := time.Now().UnixNano() - w.startTimeNanos + 1
	w.totalTimeNanos += lastTime
	w.lastTaskInfo = &TaskInfo{w.currentTaskName, lastTime}
	if w.KeepTaskList {
		w.taskList = append(w.taskList, *w.lastTaskInfo)
	}
	w.taskCount++
	w.currentTaskName = ""
}
func (w *StopWatch) isRunning() bool {
	return len(w.currentTaskName) != 0
}
func (w *StopWatch) getLastTaskTimeNanos() int64 {
	if w.lastTaskInfo == nil {
		log.Fatal("No tasks run: can't get last task interval")
	}
	return w.lastTaskInfo.TimeNanos
}
func (w *StopWatch) getLastTaskTimeMillis() int64 {
	if w.lastTaskInfo == nil {
		log.Fatal("No tasks run: can't get last task interval")
	}
	return NanosToMillis(w.lastTaskInfo.TimeNanos)
}
func (w *StopWatch) getLastTaskName() string {
	if w.lastTaskInfo == nil {
		log.Fatal("No tasks run: can't get last task name")
	}
	return w.lastTaskInfo.TaskName
}
func (w *StopWatch) getLastTaskInfo() TaskInfo {
	if w.lastTaskInfo == nil {
		log.Fatal("No tasks run: can't get last task info")
	}
	return *w.lastTaskInfo
}
func (w *StopWatch) GetTotalTimeNanos() int64 {
	return w.totalTimeNanos
}
func (w *StopWatch) GetTotalTimeMillis() int64 {
	return NanosToMillis(w.totalTimeNanos)
}
func (w *StopWatch) getTotalTimeSeconds() int64 {
	return nanosToSeconds(w.totalTimeNanos)
}
func (w *StopWatch) GetTaskInfo() []TaskInfo {
	if !w.KeepTaskList {
		log.Fatal("Task info is not being kept!")
	}
	return w.taskList
}
func (w *StopWatch) ShortSummary() string {
	return fmt.Sprintln("StopWatch '", w.id, "': running time = ", w.totalTimeNanos, " ns")
}
func (w *StopWatch) PrettyPrint() string {
	var sb strings.Builder
	sb.WriteString(w.ShortSummary())
	sb.WriteString("\n")
	if !w.KeepTaskList {
		sb.WriteString("No task info kept")
	} else {
		sb.WriteString("---------------------------------------------\n")
		sb.WriteString("ms            %       Task name\n")
		sb.WriteString("---------------------------------------------\n")
		for _, task := range w.GetTaskInfo() {
			pf := float64(task.TimeNanos) / float64(w.GetTotalTimeNanos())
			sb.WriteString(fmt.Sprintf("%-11d  %5.1f%%   %s\n", NanosToMillis(task.TimeNanos), pf*100, task.TaskName))
		}
	}
	return sb.String()
}
func (w *StopWatch) String() string {
	var sb strings.Builder
	sb.WriteString(w.ShortSummary())
	if w.KeepTaskList {
		for _, task := range w.GetTaskInfo() {
			sb.WriteString("; [")
			sb.WriteString(task.TaskName)
			sb.WriteString("] took ")
			sb.WriteString(fmt.Sprintf("%d", task.TimeNanos))
			sb.WriteString(" ns")
			percent := int64(math.Round(100.0 * float64(task.TimeNanos) / float64(w.totalTimeNanos)))
			sb.WriteString(" = ")
			sb.WriteString(fmt.Sprintf("%d", percent))
			sb.WriteString("%")
		}
	} else {
		sb.WriteString("; no task info kept")
	}
	return sb.String()
}
