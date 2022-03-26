package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/spf13/viper"
)

type DAO interface {
	NewUserQuery() UserQuery
	NewCourseQuery() CourseQuery
	NewAnswerQuery() AnswerQuery
	NewIndicatorQuery() IndicatorQuery
	NewQuestionQuery() QuestionQuery
	NewReviewQuery() ReviewQuery
	NewScoreQuery() ScoreQuery
	NewSectionQuery() SectionQuery
	NewTransactionQuery() TransactionQuery
}

type dao struct{}

var DB *sql.DB

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(int)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (d *dao) NewTransactionQuery() TransactionQuery {
	return &transactionQuery{}
}

func (d *dao) NewSectionQuery() SectionQuery {
	return &sectionQuery{}
}

func (d *dao) NewScoreQuery() ScoreQuery {
	return &scoreQuery{}
}

func (d *dao) NewReviewQuery() ReviewQuery {
	return &reviewQuery{}
}

func (d *dao) NewQuestionQuery() QuestionQuery {
	return &questionQuery{}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewCourseQuery() CourseQuery {
	return &courseQuery{}
}

func (d *dao) NewAnswerQuery() AnswerQuery {
	return &answerQuery{}
}

func (d *dao) NewIndicatorQuery() IndicatorQuery {
	return &indicatorQuery{}
}
