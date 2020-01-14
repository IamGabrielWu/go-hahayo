### 常量
与其他语言的差异
快速设置连续值
const (
  Monday = iota+1
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
  Sunday
)

const (
  Open = 1 << iota
  Close
  Pending
)
### Iota basic example
* The iota keyword represents successive integer constants 0, 1, 2,…
* It resets to 0 whenever the word const appears in the source code,
* and increments after each const specification.
