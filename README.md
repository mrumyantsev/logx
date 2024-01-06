# Multilog Go Library

This logging library can write logs mainly to a console and secondly to the specific places: databases, files, etc. But the logger does not have the implementations for any of this kind log objects. Therefore, it's up to programmer to create any needed for him *log writers* endpoints. Every log writer should implement **Writer** interface from `multilog/log` package:

``` Go
type Writer interface {
	WriteLog(datetime time.Time, levelId uint8, message string) error
}
```

# Demonstration

Use the command below to see a console demo.

```
make run/demo
```
