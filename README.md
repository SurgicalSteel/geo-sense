# GEO-SENSE  
Developed by Yuwono Bangun Nagoro a.k.a [SurgicalSteel](https://www.github.com/SurgicalSteel) --> the dumbest software engineer ever existed in your company!    

This is the Geo-Sense : A go program to find the nearest Warehouse from Shipping Address.  
The idea is actually really simple by using worker pool (default uses 25 workers) to speed up the calculation of distance for all Warehouse to the Shipping Address.  

The distance calculation itself is using [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula). The math implementation was quite hard.  

Some funny stats :  
I ran this code several times on my laptop (standard Lenovo Thinkpad with Intel core i5 8th gen and 16 GBs of RAM).  
By using 25 workers processing the worldcities_small.csv containing 5000 records of dummy Warehouses (which are sample of cities around the world).  
The execution time varies between 0.17844 seconds to 0.27370 seconds (see the log).  
Believe it or not!  

About the dataset :  
I got the dataset from [this site](https://simplemaps.com/data/world-cities) to get the worldcities_large.csv (which contains +/- 40K rows of cities around the world).  
And for the testing purpose, I use the first 5K rows of the dataset.  
Dataset Copyright to [https://simplemaps.com](https://simplemaps.com).  



