package visitor

type Database struct { 
	power int 
}

func NewDatabase(power int) *Database { 
	return &Database{power} 
}

func (db *Database) Accept(visitor Visitor) { 
	visitor.GetDatabase(db) 
}

func (db *Database) AssignDatabase() int { 
	return db.power 
}