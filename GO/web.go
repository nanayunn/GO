package main

import (
	"database/sql"

	"log"

	"github.com/mattn/go-redmine"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	projectUrl := "projectUrl"
	apiKey := "apiKey"

	check := redmine.NewClient(projectUrl, apiKey) //새로운 클라이언트 인터페이스를 생성

	/* redmine 내의 IssueFilter 구조체에 접근하여 특정 값을 할당한다.
	모든 상태의 일감을 가져오기 위해 StatusId:"*"로 값을 대입.*/
	is := redmine.IssueFilter{StatusId: "*"}

	/* IssuesByFilter로 redmine.Resoureceset을 가져오는 함수는 Filter 구조체의 포인터를 인자로 받는다. */
	iss, err := check.IssuesByFilter(&is)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "/home/nykim/project/go_project.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() /* 프로그램 종료되면 db도 닫힌다. */

	/* table이 존재하면 drop */
	dropStmt := `
	drop table teveloper
	`
	_, err = db.Exec(dropStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, dropStmt)
		return
	}

	/* table 생성 */
	sqlStmt := `
	create table teveloper (project_name text not null, issue_id integer not null primary key, tracker_name text not null, status_name text not null, subject text not null, author text not null, created_date text not null);
	delete from teveloper
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	/* index는 사용하지 않으므로 언더바 처리, 일감 묶음을 range 범위로 주었다. */
	for _, value := range iss {

		/* 정보 추출 */
		subject := value.Subject
		project_name := value.Project.Name
		issue_id := value.Id
		tracker_name := value.Tracker.Name
		status_name := value.Status.Name
		author := value.Author.Name
		created_date := value.CreatedOn

		/* db 시작,  */
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		tx.Exec("insert into teveloper values(?,?,?,?,?,?,?)", project_name, issue_id, tracker_name, status_name, subject, author, created_date)

		tx.Commit()
	}

}
