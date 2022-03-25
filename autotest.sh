go test -v
fswatch -o ./ | xargs -n1 -I{} go test -v
