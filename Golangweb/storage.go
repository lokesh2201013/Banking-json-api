package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Storage is an interface that defines methods for interacting with account data.
type Storage interface {
	// CreateAccount creates a new account.
	CreateAccount(*Account) error

	// DeleteAccount deletes an account by its ID.
	DeleteAccount(int) error

	// UpdateAccount updates an existing account.
	UpdateAccount(*Account) error

	// GetAccountById retrieves an account by its ID and returns either an ApiError or the account.
	GetAccountById(int) (*ApiError, error)

	GetAccounts()([]*Account,error)
}

type PostgresStore struct{
	db *sql.DB
}
 func NewPostgresStore ()(*PostgresStore,error){
	//connect to database
	connstr:=""
    db,err:=sql.Open("postgres",connstr)
	if err!=nil{
		return nil,err
	}
	if err:=db.Ping();err!=nil{
		return nil,err
	}
	return &PostgresStore{
		db: db,
	},nil
}  

func(s *PostgresStore)Init ()error{
	return s.createAccountTable()
}

func (s *PostgresStore)createAccountTable()error{
	query:=`create table  if not exists account(
	 id serial primary key,
	 first_name varchar(50),
	  last_name varchar(50),
	  number serial,
	  balance int,
	  created_at timestamp
	)`

	_,err:=s.db.Exec(query)
	return err
}
func (s *PostgresStore) CreateAccount(acc *Account)error{
	
	query:=`insert into account(first_name,last_name,number,balance,created_at) values($1,$2,$3,$4,$5)`

	resp,err:=s.db.Query(query,acc.FirstName,acc.LastName,acc.Number,acc.Balance,acc.CreatedAt)
	if err!=nil{
		return err
	}
	fmt.Printf("%+v\n",resp)
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account)error{
	return nil
}

func (s *PostgresStore) DeleteAccount(id int)error{
	return nil
}

func (s *PostgresStore) GetAccountById(id int)(*ApiError,error){
	return nil,nil
}


func (s *PostgresStore) GetAccounts()([]*Account,error){
	rows,err:=s.db.Query("select * from account ")
	if err!=nil{
		return nil,err
	}
	accounts:=[]*Account{}

	for rows.Next(){
		account:=new(Account)
		err:=rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err!=nil{
			return nil,err
		}
	accounts=append(accounts,account)
	}

	return accounts,nil
}
