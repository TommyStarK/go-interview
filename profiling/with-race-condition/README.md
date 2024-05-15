# race condition

Go's race detector does runtime analysis. It has no false positives, but it does have false negatives. If it doesn't actually see a race, it can't report it.
