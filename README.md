# calculation
Reusable Golang library to provide functions for mathematical calculation

## Status
Under development; not ready for use

## Areas of emphasis
The functions are a diverse set related to:
* Algebra
* Probability
* Statistics

## Comparison to OpenOffice
Here is a list of functions in this stat library compared to OpenOffice functions.

| stat | OpenOffice |
| ---- | ---------- |
| BinProb, BinDist | B, BINOMDIST |
| FracLim | PERCENTILE, QUARTILE |
| FracRank | PERCENTRANK |
| GeomMean | GEOMEAN |
| HarMean | HARMEAN |
| Kurt | KURT |
| Large | LARGE |
| Max | MAX |
| Mean | AVERAGE |
| MeanAbsDev | AVEDEV |
| Med | MEDIAN |
| Midrange | n/a |
| Min | MIN |
| Mode | MODE |
| Prod | n/a |
| Range | n/a |
| Rank | RANK |
| RootMeanSqr | n/a |
| Skew | SKEW |
| Small | SMALL |
| StdDev | STDEV |
| StdErr | n/a |
| StdDevPop | STDEVP |
| StdErrPop | n/a |
| SumSqrDev | DEVSQ |
| TrimmedMean | TRIMMEAN |
| Var | VAR |
| VarPop | VARP |
| WeightedMean | n/a |

## Math notes
* The Near function calculates near-floating point equality within +/-1e-7, so test values are given with 8 decimal
  points of accuracy to pass this test.

## Programming notes
See [abbreviation](https://www.github.com/BluntSporks/abbreviation) for a list of abbreviations used.

