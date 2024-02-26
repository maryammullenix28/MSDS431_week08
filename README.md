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
### Results & Recommendations
The outputs of the R and GO code are comparable after 10000 iterations. Given the random resampling process, the results won't be exactly the same.

The R outputs:
- Range: 0.6839681 0.9929641
- Mean correlation: 0.8955649
- St.Dev: 0.04318599

The Go outputs:
- Range: 0.7035759526865184 0.9897836519728416 
- Mean: 0.8946565705620442
- St.Dev: 0.04300598711824439

As for timing and performance, at 10000 iterations, the Go program took 0.228 s whereas the R program took 0.384 s. It's worth noting that the Go program did perform faster, despite more code complexity and steps.
At 100000 iterations, the R program took 4.153 s, a significant increase from 10000 iterations. Go, on the other hand, only increased by 0.02 seconds at 0.249 total run time.

In regards to computing costs, based on Google Cloud's calculator, at 100000 iterations and 50 calls per month, the R program would cost about $2,146.41. Alternatively, running the Go program would cost $146.74. This saves companies $2000 per month. 

Based on this analysis, I would recommend using Go when it requires a large number of trials or when conducting analysis with large-scale datasets due to its efficiency and memory management capabilities. This would save companies money in cloud computing costs, despite the greater time investment in the beginning to research and/or create custom functions/programs to conduct their analyses.