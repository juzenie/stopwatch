# 程序耗时统计
Program Time Counter

| ms   | %      | Task name |
| ---- | ------ | --------- |
| 1010 | 20.1%  | step1     |
| 2001 | 39.9%  | step2     |
| 2008 | 40.0%  | other     |
| ---- | ------ | --------- |
| 5020 |        | total     |

```go
func main() {
	var w swer.Sw = swer.NewPpStopWatch()
	w.Start("step1")
	time.Sleep(time.Second)
	w.Stop()
	time.Sleep(time.Second * 2)
	w.Start("step1")
	time.Sleep(time.Second * 2)
	w.Stop()
	fmt.Println(w.PrettyPrint())
}
```
## stopwatch

这是一个 golang 程序的耗时统计报告，报告中显示了不同模块的任务耗时。下面是每个模块的详细说明:

- **step1**：该模块耗时 1010 毫秒，占总耗时的 20.1%。
- **step2**：该模块耗时 2001 毫秒，占总耗时的 39.9%。
- **other**：该模块耗时 2008 毫秒，表示该模块除了其他已统计的模块之外的耗时，占总耗时的 40.0%。

总耗时为 5020 毫秒。

这个报告可以帮助您了解程序中各个模块的执行时间，以及每个模块在总执行时间中所占的比例。根据这些数据，您可以更好地优化程序性能，特别是耗时较长的模块。

请根据实际情况分析耗时报告并采取相应的优化措施，以提高程序的执行效率和性能。
