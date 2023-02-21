package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
} ali 7715

On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    Ali.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        Ali_7715.go

no changes added to commit (use "git add" and/or "git commit -a")