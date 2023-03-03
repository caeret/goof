# Usage

```go
t := goof.New[string]()
t.Go(func(ctx context.Context) (string, error) {
	time.Sleep(time.Millisecond * 10)
	return "1", nil
})
t.Go(func(ctx context.Context) (string, error) {
	time.Sleep(time.Millisecond * 5)
	return "2", nil
})
fmt.Println(t.First())
// Output: 2 true
```