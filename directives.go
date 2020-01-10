package main

type Directives map[string]bool

func (d Directives) Exist(arg string) bool {
    _, ok := d[arg]
    return ok
}

func (d Directives) Find(args []string, directive string) bool {
    for _, arg := range args {
        if directive == arg {
            return true
        }
    }
    return false
}

func (d Directives) Collect(args []string, directive string) (string, bool) {
    for i, arg := range args {
        if directive == arg {
            hasArgument := d[arg]
            if !hasArgument {
                return "", true
            } else if i+1 < len(args) {
                return args[i+1], true
            }
            break
        }
    }
    return "", false
}