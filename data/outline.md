1   get the data from the xml.files and creating tables Users, Task Perform, Status(3 entries)
    names   egn     zdravnaKasa_id      user_id
    last3status     dates   typeService
    
    pointer to external Referenses Xray pictures ASD manner

2   the xml file should be processed by date created asd

3   search optimization 1 - by names initials
                        2 - egn using # to switch the search fiels to 
                        3 - zdrKasa_id % to
                        4 - by entry != (last)Year

4   determine what and how to be presented  1

5   actions     1 check health status NZOK
                2 last dental status
                3 work done new dental status
                    send the status
                    postpone/pending
                    cancel


program interface - web browser
                    its own inteface                    
                
dbeaver password - masterkey

	"database/sql"                      -> interface
	_ "github.com/mattn/go-sqlite3"     -> driver for sqlite3

ctrl+r string  to cycle trough the command contaning that string
