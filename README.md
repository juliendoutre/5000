# 5000

## Summary
5000 is a simple dice game.

This repository implements an engine as a Go library for this game and runs some simulations with various strategies as an attempt to find an optimal one (WIP).

## Rules
> There are a lot of variations for this game's rules. I sticked to the following ones:

A player rolls five dices and counts its points:
- any `1` marks 100 points
- any `5` marks 50 points
- three-of-a-kinds mark 100 times their rank (for instance `3,3,3` marks 300 and `5,5,5` marks 500) except `1,1,1` which marks 1000.

If a round does not mark any points, the turns stops right there and the player ends with 0 points.

Else, it can choose to stop right there and confirm its score for this turn or to roll some dices again, starting a new round to try to increase its score.

It has to exclude at least one counter before starting the next round though. This means if it starts with `n` dices, it will roll at maximum `n - 1` dices the next one.

There's one exception to this last rule: if the player achieves to have all dices to mark in one round, it can choose to get all 5 dices back in its hand and start again with all of them.

## Testing
I implemeted some property based tests to try this pattern out.
