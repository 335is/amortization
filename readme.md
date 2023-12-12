# amortization

Command line program that calculates and displays a mortgage loan payment schedule.

Takes command line input for the required variables.

## How to Use

Example:  
.\amortization.exe -loan 500000 -rate 6.25 -term 360

## Command Line Parameters
| option | parameter type | description                            |
| ------ | -------------- | -----------                            |
| -h     |                | display command line parameter usage   |
| -extra | float          | extra monthly principal in dollars     |
| -loan  | float          | loan amount in dollars                 |
| -rate  | float          | interest rate in percentage            |
| -term  | int            | loan duration in months                |

