# WHAT IS THIS?

A quick script made solely for the functionality of 'scraping' gifs off of a website, saving it to a local directory.

## CONFIGURATION

Some of the configuration fields in the `.env` file include:
```
CRAWL=y
```
By setting the value to "y", you allow the script to "crawl" into the given page URL's page links, allowing to scrape gifs from multiple pages at once.
```
savePATH=/gifs/
```
savePATH is where the saved gifs are stored, the folder name is relative to the script's location.
```
saveURL=99gifshop.neocities.org
```
saveURL is the page URL that the script uses to scrape gifs off of.