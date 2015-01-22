package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)


func GetGitInfo() string {
	
	
	// Read in the gitconfig file
    file, err := os.Open("./.git/config")
	if err != nil {
		log.Println("Can't find the project.  Use the -p option to specifiy the project you want to build.")
	    log.Fatal(err)		
	}
    defer file.Close()
	
	// find the first reference to an atlas repository
	remote := ""
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		for _,s := range strings.Split(scanner.Text()," ") {
			if strings.Contains(s, "git@git.atlas.oreilly.com") {
	        	remote = s	
				break
			}			
		}
    }

	// Error our if we can't find anything
	if remote == "" {
        log.Fatal("Cannot find an atlas remote in the git config file.  Use the -p option to suppy one directly.")
    }
	
	// Assuming we found the remote, parse out just the username and project, which will look like this:
	//    git@git.atlas.oreilly.com:odewahn/ost-python.git
	
	repo := strings.Split(remote,":")
	if len(repo) !=2 {
        log.Fatal("Cannot find an atlas remote in the git config file.  Use the -p option to suppy one directly.")		
	}
	
	ret_val := repo[1]
	if strings.Contains(ret_val, ".git") {
		ret_val = ret_val[:len(ret_val)-4]
	}
		
	return ret_val

}