package main

import (
   "code.google.com/p/gcfg"
   "log"
   "os"
   "strings"
)

type Config struct {
	    Core struct {
			Repositoryformatversion int
			Filemode bool
			Bare bool
			Logallrefupdates bool
			Ignorecase bool
		}
		Branch map[string]*struct {
			Remote string
			Origin string
		}
        Remote map[string]*struct {
                Url string
				Fetch string
        }
}


// Tries to parse the ./.git/config file to find the atlas project name
func GetGitInfo() string {
	
	c := Config{}
	
	_, err := os.Stat("./.git/config")
	if err != nil {
	    log.Fatal("You're not in a git directory.")
	}
	
	err = gcfg.ReadFileInto(&c, "./.git/config")
	if err != nil {
		log.Fatal("Can't parse /.git/config")
	}
	
	// look through each remote and return the one that points to an atlas project
	project := ""
	for _,c := range c.Remote {
		if c.Url[:3] == "git" {
			repo := strings.Split(c.Url,":")
			if len(repo) == 2 {
				host := strings.Split(repo[0],"@")[1]
				if host == "git.atlas.oreilly.com" {
					project = strings.Split(repo[1],".")[0]
					break
				}
			}
			
		}
	}
	
	if len(project) == 0 {
		log.Fatal("Can't find any remotes for git.atlas.oreilly.com.")
	}
		
	return project

}