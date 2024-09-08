package service

import (
	"backend/database"
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func GitPull() {
	cmd := exec.Command("git", "pull")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running git pull:", err)
		return
	}
	fmt.Println(string(out))
}

func GetChangeFiles() {
	// 假设你已知的上次 commit 号
	lastCommit, err := database.GetLastCommit()
	if err != nil {
		fmt.Println("Error getting last commit:", err)
		return
	}
	fmt.Println("Last Commit:", lastCommit)

	getChangeFilesFromTwoCommits(lastCommit)
}

func getChangeFilesFromTwoCommits(lastCommit string) {
	// 获取当前的 commit 号
	currentCommit, err := getCurrentCommit()
	if err != nil {
		fmt.Println("Error getting current commit:", err)
		return
	}
	fmt.Println("Current Commit:", currentCommit)

	// 获取两次 commit 之间的文件变化
	added, deleted, err := getCommitDifferences(lastCommit, currentCommit)
	if err != nil {
		fmt.Println("Error getting commit differences:", err)
		return
	}

	fmt.Println("Added files:", added)
	fmt.Println("Deleted files:", deleted)
}

func getCurrentCommit() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getCommitDifferences(oldCommit, newCommit string) ([]string, []string, error) {
	cmd := exec.Command("git", "diff", "--name-status", oldCommit, newCommit)
	out, err := cmd.Output()
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	var added []string
	var deleted []string
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			status := parts[0]
			file := parts[1]
			switch status {
			case "A": // Added
				added = append(added, file)
			case "D": // Deleted
				deleted = append(deleted, file)
			}
		}
	}
	return added, deleted, nil
}
