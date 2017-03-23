# go_table_test
playing with golang table driven testing

#### Credits
[fib testing source](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)

### Sample Test Output:  
```
$ go test ./... -v
=== RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok      github.com/m17ch/go_table_test/fib      0.075s
=== RUN   TestHalves
--- PASS: TestHalves (0.00s)
=== RUN   TestIsEven
--- PASS: TestIsEven (0.00s)
=== RUN   TestAddition
--- PASS: TestAddition (0.00s)
PASS
ok      github.com/m17ch/go_table_test/mymath   0.069s
=== RUN   TestSlow
--- PASS: TestSlow (5.00s)
PASS
ok      github.com/m17ch/go_table_test/slowtest 5.053s
```