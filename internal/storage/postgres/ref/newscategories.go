package ref

//go:generate reform

// Newscategories represents a row in newscategories table.
//
//reform:newscategories
type Newscategories struct {
	ID     int32 `reform:"id,pk"`
	NewsID int32 `reform:"news_id"`
}
