### 多个条件使用switch
```
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 4:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
        }
```
