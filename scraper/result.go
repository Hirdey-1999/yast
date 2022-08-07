/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package scraper

// Interface that accepts any type of result coming from Scraper
type Result interface {
}

type Results []Result
