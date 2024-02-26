# Modern Applied Statistics with Go
For this assignment, students were asked to evaluate the ease of using Go instead of R for one of the modern computer intensive statistical methods, as listed in Gelman and Vehtari's "What Are the Most Important Statistical Ideas of the Past 50 Years?" (2021) article. For this assignment, I chose to implement Go code for bootstraping. Bootstrapping is a resampling technique used to estimate the sampling distribution of a statistic by repeatedly resampling from the observed data with replacement. The analysis looks at Celtics data that contains the player, position, percentage drafted in Fantasy league teams, and the fantasy points.
### Working in R
R contains a specific library [boot](https://cran.r-project.org/web/packages/bootstrap/index.html) to conduct the analysis. Aside from programming time, the bootstraping code could be written in 33 lines (includes imports and time/memory benchmarking). The code and dataset for the R analysis came from [A Practical Guide to Bootstrap in R](https://medium.com/p/bd975ec6dcea) by Dr. Leihua Ye in Towards Data Science. 
I also used the 'pryr' package to measure memory usage in R.
### Go Implementation
Unlike in R, instead of using a specific package or library specific to bootstraping, I implemented the bootstaping from scratch using the [gonum/stat]("gonum.org/v1/gonum/stat") general statistics package.

Th R code output contained three key attributes: mean, standard deviation, and range.  To compute these statistics, I first needed to calculate the bootstrap correlation values between the %Drafted and FTPS columns.
#### Resampling with 'sampleWithReplacement'
Resampling with sampleWithReplacement
The sampleWithReplacement function randomly selects elements from the input slices XDrafted and FPTS with replacement, which means the same element can be selected multiple times. This process creates new datasets with the same length as the original ones but with random permutations.
#### Correlation Calculation with 'bootstrapCorrelation'
The bootstrapCorrelation function takes the resampled datasets created by sampleWithReplacement, calculates the correlation coefficient for each resampled dataset, and stores the results in a slice. By repeating this process R times (where R is the number of bootstrap iterations), we obtain R correlation coefficients that represent the variability of the correlation estimate.
#### Testing
To make sure the bootstraping is occuring correctly, I created a unit test to assess 'bootstrapCorrelation's ability to calculate accurate correlation values. The code passed the test.
### Results
The outputs of the R and GO code are comparable. Given the random resampling process, the results won't be exactly the same. 
The R outputs:
- Range: 0.6839681 0.9929641
- Mean correlation: 0.8955649
- St.Dev: 0.04318599
The Go outputs:
- 