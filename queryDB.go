/*
Library functions used in urlVal service
*/
package main

import (
	"fmt"
	"math"
)

//QueryDB checks hostname against db and return malwareType
func (ds *AppContext) QueryDB(hname string) string {
	var mType string
	if t, ok := ds.MalMap[hname]; ok {
		mType = t
		ds.CacheCount[hname]++
	} else { // not in cache search db
		err := ds.DbHandler.QueryRow("SELECT type FROM malware WHERE hostname = ?", hname).
			Scan(&mType)
		if err != nil { // not in db either
			mType = "clean"
		}
		if len(ds.MalMap) < MAXCACHEENTRY {
			ds.MalMap[hname] = mType // adding to cache
			ds.CacheCount[hname] = 1
			fmt.Printf("Adding %s to cache.\n", ds.MalMap[hname])
		} else { // no more cache entries, replace the least used one
			var minKey string
			min := math.MaxUint32
			for k, v := range ds.CacheCount {
				if v < min {
					min = v
					minKey = k
				}
			}
			fmt.Printf("Replacing cache entry %s cache count %d with %s.\n", minKey, ds.CacheCount[minKey], hname)
			delete(ds.CacheCount, minKey) //delete the cache entry
			delete(ds.MalMap, minKey)
			ds.MalMap[hname] = mType // adding to cache
			ds.CacheCount[hname] = 1
		}

	}

	return mType
}
