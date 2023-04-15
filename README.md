Task for lesson

Usage:
```
import "store/store"

// ...

cache := store.New() // instantiate
    
cache.Set(<string>, <any>) // set any value by string key

cache.Get(<string>) // get any value by string key

cache.Delete(<string>) // delete any value by string key
```
