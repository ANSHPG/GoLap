# Understanding Modules in GO

## Packages vs Repositories

### Package
``` 
Its Basically a way of packaging a software as simple as it gets, but as the name suggests package must combine several files inside it but if we remove one of the them the pkg is not function properly , those files combined make a package.
```

### Repository
```
Here unlike package it conains a lot of files, but even we happen to modify one of them its still gonna be that repository .
```

## Modules in GO
```
Its similar to repository but in a more general way, not attached to github
```
## Semantic Version

### Major Version V1.x.x

```
It signifies major channges to go, but also it signifies that some operations in the pervious versation wont be functioning in the laest major version, that means it may lag backward compability  
```

### Minor Version V1.1.x

```
Here Minopr changes takes place, so the backward compatibility exists, and the operations of older versions supports in the latest update
```

### Patch Version V1.1.1

```
Here some minute bugs are fixed and micro versions are released 
```