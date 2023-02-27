package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
On branch main
Your branch is up to date with 'origin/main'.

<<<<<<< HEAD
nothing to commit, working tree clean
=======
nothing to commit, working tree clean
    
>>>>>>> 91a84beccf1e179bffa2727285b30b5417f78e63
