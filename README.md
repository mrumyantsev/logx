# Multilog Go Library

This logging library can write logs mainly to a console and secondly to the specific places: databases, files, etc. But the logger does not have the implementations for any of this kind log objects. Therefore, it's up to programmer to create any needed for him *log writers* endpoints. Every log writer should implement **LogWriter** interface:

``` Go
type LogWriter interface {
	WriteLog(datetime string, messageType string, message string) error
}
```

# Demonstraction

Use the command below to see a console demo.

```
make run/demo
```
