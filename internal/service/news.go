package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/v7ktory/test/internal/storage/postgres/ref"
)

func (s *Service) GetList() ([]ref.News, error) {
	var news ref.News

	res, err := s.db.SelectAllFrom(news.View(), "")
	if err != nil {
		return nil, fmt.Errorf("error selecting from news table: %w", err)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no news found")
	}

	var newsList []ref.News

	for _, value := range res {
		news := *value.(*ref.News)
		categories, err := s.db.SelectAllFrom(ref.NewscategoriesTable, "WHERE news_id = $1", news.ID)

		if err != nil {
			return nil, fmt.Errorf("error selecting from newscategories table: %w", err)
		}

		for _, c := range categories {
			news.Categories = append(news.Categories, int(c.(*ref.Newscategories).ID))
		}

		newsList = append(newsList, news)
	}
	return newsList, nil
}

func (s *Service) Edit(id int, news ref.News) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("error rolling back transaction: %w", e)
			}
		}
	}()

	if !s.ifExists(id) {
		return fmt.Errorf("news with id %d not found", id)
	}

	updateFields := make(map[string]interface{})
	if news.Title != "" {
		updateFields["title"] = news.Title
	}
	if news.Content != "" {
		updateFields["content"] = news.Content
	}

	if len(updateFields) == 0 && news.Categories == nil {
		return fmt.Errorf("nothing to update")
	} else if len(updateFields) > 0 {
		query := "UPDATE news SET"
		values := make([]interface{}, 0)
		idx := 1
		for key, value := range updateFields {
			query += " " + key + " = $" + strconv.Itoa(idx) + ","
			idx++
			values = append(values, value)
		}
		query = strings.TrimSuffix(query, ",")
		query += " WHERE id = $" + strconv.Itoa(idx)
		values = append(values, id)

		_, err = tx.Exec(query, values...)
		if err != nil {
			return fmt.Errorf("error updating news: %w", err)
		}
	}

	if news.Categories != nil {
		_, err = tx.DeleteFrom(ref.NewscategoriesTable, "WHERE news_id = $1", id)
		if err != nil {
			return errors.Wrap(err, "error deleting old categories")
		}

		_, err = tx.Exec("SELECT setval('news_categories_id_seq', (SELECT MAX(id) FROM news_categories));")
		if err != nil {
			return errors.Wrap(err, "error updating id news_categories")
		}

		for _, c := range news.Categories {
			category := &ref.Newscategories{
				ID:     int32(c),
				NewsID: int32(id),
			}
			err = tx.Save(category)
			if err != nil {
				return fmt.Errorf("error saving category: %w", err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (s *Service) ifExists(id int) bool {
	var news ref.News
	if err := s.db.FindByPrimaryKeyTo(&news, id); err != nil {
		return false
	}
	return true
}
