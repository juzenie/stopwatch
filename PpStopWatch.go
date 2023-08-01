package main

import (
	"fmt"
	"strings"
	"time"
)

type PpStopWatch struct {
	StopWatch

	startTimeMs int64
}

func NewPpStopWatch() *PpStopWatch {
	return &PpStopWatch{
		StopWatch:   *NewTaskInfo(),
		startTimeMs: time.Now().UnixMilli(),
	}
}
func (w *PpStopWatch) PrettyPrint() string {
	total := time.Now().UnixMilli() - w.startTimeMs
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
			pf := float64(NanosToMillis(task.TimeNanos)) / float64(total)
			sb.WriteString(fmt.Sprintf("%-11d  %5.1f%%   %s\n", NanosToMillis(task.TimeNanos), pf*100, task.TaskName))
		}
		other := total - w.GetTotalTimeMillis()
		sb.WriteString(fmt.Sprintf("%-11d  %5.1f%%   %s\n", other, float64(other)/float64(total)*100, "other"))
		sb.WriteString("---------------------------------------------\n")
		sb.WriteString(fmt.Sprintf("%-18d    %s\n", total, "total"))
	}
	return sb.String()
}
